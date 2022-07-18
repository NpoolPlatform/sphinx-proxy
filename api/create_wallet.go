package api

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	coininfopb "github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	cconst "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	ct "github.com/NpoolPlatform/sphinx-plugin/pkg/types"
	putils "github.com/NpoolPlatform/sphinx-plugin/pkg/utils"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	ocodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type walletDoneInfo struct {
	success bool
	message string
	address string
}

var walletDoneChannel = sync.Map{}

func (s *Server) CreateWallet(ctx context.Context, in *sphinxproxy.CreateWalletRequest) (out *sphinxproxy.CreateWalletResponse, err error) {
	_, span := otel.Tracer(sconst.ServiceName).Start(ctx, "CreateWallet")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(ocodes.Error, "call CreateWallet")
			span.RecordError(err)
		}
	}()

	span.SetAttributes(
		attribute.String("Name", in.GetName()),
	)

	if in.GetName() == "" {
		logger.Sugar().Errorf("CreateWallet Name: %v empty", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name empty")
	}

	coinType, err := putils.ToCoinType(in.GetName())
	if err != nil {
		logger.Sugar().Errorf("CreateWallet Name: %v invalid", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name Invalid")
	}

	span.AddEvent("call getProxySign")
	signProxy, err := getProxySign()
	if err != nil {
		logger.Sugar().Errorf("Get ProxySign client not found")
		return out, status.Error(codes.Internal, "internal server error")
	}

	// query current coin net
	span.AddEvent("call GetGRPCConn")
	conn, err := grpc2.GetGRPCConn(cconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		logger.Sugar().Errorf("GetGRPCConn not get valid conn: %v", err)
		return out, status.Error(codes.Internal, "internal server error")
	}
	defer conn.Close()

	cli := coininfopb.NewSphinxCoinInfoClient(conn)

	ctx, cancel := context.WithTimeout(ctx, sconst.GrpcTimeout)
	defer cancel()

	span.AddEvent("call grpc GetCoinInfo")
	coinInfo, err := cli.GetCoinInfo(ctx, &coininfopb.GetCoinInfoRequest{
		Name: in.GetName(),
	})
	if err != nil {
		logger.Sugar().Errorf("GetCoinInfo Name: %v error: %v", in.GetName(), err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	payload, err := json.Marshal(ct.NewAccountRequest{
		ENV: coinInfo.GetInfo().GetENV(),
	})
	if err != nil {
		logger.Sugar().Errorf("Marshal balance request Addr: %v error %v", "", err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	var (
		uid  = uuid.NewString()
		done = make(chan walletDoneInfo)
	)

	span.AddEvent("send chan message to sign",
		trace.WithAttributes(
			attribute.String("Name", in.GetName()),
			attribute.Int64("TransactionType", int64(sphinxproxy.TransactionType_WalletNew)),
			attribute.String("TransactionID", uid),
		),
	)
	walletDoneChannel.Store(uid, done)
	signProxy.walletNew <- &sphinxproxy.ProxySignRequest{
		CoinType:        coinType,
		TransactionType: sphinxproxy.TransactionType_WalletNew,
		TransactionID:   uid,
		Payload:         payload,
	}

	// timeout, block wait done
	select {
	case <-time.After(sconst.GrpcTimeout):
		walletDoneChannel.Delete(uid)
		logger.Sugar().Error("create wallet wait response timeout")
		return out, status.Error(codes.Internal, "internal server error")
	case info := <-done:
		walletDoneChannel.Delete(uid)
		if !info.success {
			logger.Sugar().Errorf("wait create wallet done error: %v", info.message)
			return out, status.Error(codes.Internal, "internal server error")
		}
		out = &sphinxproxy.CreateWalletResponse{
			Info: &sphinxproxy.WalletInfo{
				Address: info.address,
			},
		}
	}

	return out, nil
}
