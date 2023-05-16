package api

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	coinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"

	coincli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/coins/getter"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type esGasDoneInfo struct {
	success bool
	message string
	payload []byte
}

var esGasDoneChannel = sync.Map{}

func (s *Server) GetEstimateGas(ctx context.Context, in *sphinxproxy.GetEstimateGasRequest) (out *sphinxproxy.GetEstimateGasResponse, err error) {
	logger.Sugar().Infof("get estimate gas info coin name: %v", in.GetName())

	if in.GetName() == "" {
		logger.Sugar().Errorf("GetEstimateGas Name: %v empty", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name empty")
	}

	// query coininfo
	coinExist, err := coincli.GetCoinOnly(ctx, &coinpb.Conds{
		Name: &npool.StringVal{
			Op:    cruder.EQ,
			Value: in.GetName(),
		},
	})
	if err != nil {
		logger.Sugar().Errorf("check coin info %v error %v", in.GetName(), err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	if coinExist == nil {
		logger.Sugar().Errorf("check coin info %v not exist", in.GetName())
		return out, status.Errorf(codes.NotFound, "coin %v not found", in.GetName())
	}

	coinType := utils.CoinName2Type(in.GetName())
	pcoinInfo := getter.GetTokenInfo(in.GetName())
	if pcoinInfo != nil || coinType == sphinxplugin.CoinType_CoinTypeUnKnow {
		coinType = pcoinInfo.CoinType
	}

	pluginProxy, err := getProxyPlugin(coinType)
	if err != nil {
		logger.Sugar().Errorf("Get PluginProxy client not found for coinType: %v", coinType)
		return out, status.Error(codes.Internal, "internal server error")
	}

	var (
		uid  = uuid.NewString()
		done = make(chan esGasDoneInfo)
	)

	now := time.Now()
	payload, err := json.Marshal(sphinxproxy.GetEstimateGasRequest{
		Name: in.GetName(),
	})
	if err != nil {
		logger.Sugar().Errorf("Marshal estimate gas request name: %v error %v", in.GetName(), err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	esGasDoneChannel.Store(uid, done)
	pluginProxy.pluginReq <- &sphinxproxy.ProxyPluginRequest{
		Name:            in.GetName(),
		CoinType:        coinType,
		TransactionType: sphinxproxy.TransactionType_EstimateGas,
		TransactionID:   uid,
		Payload:         payload,
	}

	// timeout, block wait done
	select {
	case <-time.After(sconst.GrpcTimeout * 6):
		esGasDoneChannel.Delete(uid)
		logger.Sugar().Errorf("get transactionID: %v coin name: %v estimate gas wait response timeout", uid, in.GetName())
		return out, status.Error(codes.Internal, "internal server error")
	case info := <-done:
		elasp := time.Since(now).Seconds()
		logger.Sugar().Warnf("get transactionID: %v coin name: %v estimate gas wait response use: %vs", uid, in.GetName(), elasp)
		esGasDoneChannel.Delete(uid)
		if !info.success {
			logger.Sugar().Errorf("wait get coin name:%v estimate gas done error: %v", in.GetName(), info.message)
			return out, status.Error(codes.Internal, "internal server error")
		}

		esGas := sphinxproxy.GetEstimateGasResponse{}
		if err := json.Unmarshal(info.payload, &esGas); err != nil {
			logger.Sugar().Errorf("Unmarshal estimate info coin name: %v error: %v", in.GetName(), err)
			return out, status.Error(codes.Internal, "internal server error")
		}

		out = &esGas
	}

	return out, nil
}
