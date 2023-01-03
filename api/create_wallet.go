package api

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	ct "github.com/NpoolPlatform/sphinx-plugin/pkg/types"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	ocodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	coincli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	coinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	"github.com/NpoolPlatform/sphinx-plugin/pkg/coins/getter"
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

	// query coininfo
	coinInfo, err := coincli.GetCoinOnly(ctx, &coinpb.Conds{
		Name: &npool.StringVal{
			Op:    cruder.EQ,
			Value: in.GetName(),
		},
	})
	if err != nil {
		logger.Sugar().Errorf("check coin info %v error %v", in.GetName(), err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	logger.Sugar().Errorf("GetCoinOnly %v", coinInfo)

	coinType := utils.CoinName2Type(in.GetName())
	pcoinInfo := getter.GetTokenInfo(in.GetName())
	if pcoinInfo != nil || coinType == sphinxplugin.CoinType_CoinTypeUnKnow {
		coinType = pcoinInfo.CoinType
	}

	span.AddEvent("call getProxySign")
	signProxy, err := getProxySign(in.GetName())
	if err != nil {
		logger.Sugar().Errorf("Get ProxySign client not found")
		return out, status.Error(codes.Internal, "internal server error")
	}

	payload, err := json.Marshal(ct.NewAccountRequest{
		ENV: coinInfo.GetENV(),
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
		Name:            in.GetName(),
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
