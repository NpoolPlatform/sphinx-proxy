package api

import (
	"context"
	"encoding/json"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/trading"
	pconst "github.com/NpoolPlatform/sphinx-plugin/pkg/message/const"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	msgcli "github.com/NpoolPlatform/sphinx-proxy/pkg/message/message"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/unit"
	sconst "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
	msgproducer "github.com/NpoolPlatform/sphinx-service/pkg/message/message"
)

func (s *Server) Transaction(stream signproxy.SignProxy_TransactionServer) error {
	consumerInfo, err := msgcli.ConsumeTrans()
	if err != nil {
		logger.Sugar().Errorf("fail to consume %s: %v", msgproducer.GetQueueName(), err)
		return err
	}

	for info := range consumerInfo {
		var (
			isOk       bool
			cid        string
			mpResponse *sphinxplugin.MpoolPushResponse
		)
		// if use go routine should limit go routine numbers
		tinfo := &msgproducer.NotificationTransaction{}
		if err := json.Unmarshal(info.Body, tinfo); err != nil {
			logger.Sugar().Errorf("Unmarshal consumer info error: %v ", err)
			goto callback
		}

		if mpResponse, err = mpoolPush(stream, tinfo); err != nil {
			logger.Sugar().Errorf("deal consumer info: %v error: %v", info, err)
			goto callback
		}

		isOk = true
		logger.Sugar().Infof(
			"deal consumer info TranID: %v CoinType: %v From: %v To: %v ok",
			tinfo.TransactionIDInsite,
			tinfo.CoinType,
			tinfo.AddressFrom,
			tinfo.AddressTo,
			err,
		)

	callback:
		{
			// callback
			if err := ack(isOk, cid, mpResponse, tinfo); err != nil {
				logger.Sugar().Errorf(
					"call ack ServiceName: %v error: %v",
					sconst.ServiceName,
					err,
				)
			}
		}
	}

	return nil
}

func mpoolPush(stream signproxy.SignProxy_TransactionServer, info *msgproducer.NotificationTransaction) (*sphinxplugin.MpoolPushResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.GrpcTimeout)
	defer cancel()
	pluginConn, err := grpc2.GetGRPCConn(pconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		logger.Sugar().Errorf(
			"call GetGRPCConn ServiceName: %v error: %v",
			pconst.ServiceName,
			err,
		)
		return nil, err
	}

	pluginClient := sphinxplugin.NewPluginClient(pluginConn)
	nonce, err := pluginClient.MpoolGetNonce(ctx, &sphinxplugin.MpoolGetNonceRequest{
		CoinType: info.CoinType,
		Address:  info.AddressFrom,
	})
	if err != nil {
		logger.Sugar().Errorf(
			"call MpoolGetNonce CoinType: %v Address: %v error: %v",
			info.CoinType,
			info.AddressFrom,
			err,
		)
		return nil, err
	}

	if err := stream.Send(&signproxy.TransactionRequest{
		CoinType: info.CoinType,
		Message: &sphinxplugin.UnsignedMessage{
			To:         info.AddressTo,
			From:       info.AddressFrom,
			Nonce:      nonce.GetInfo().GetNonce(),
			Value:      unit.FIL2AttoFIL(info.AmountFloat64),
			GasLimit:   11,
			GasFeeCap:  "",
			GasPremium: "",
			Method:     0, // 0 send
		},
	}); err != nil {
		logger.Sugar().Errorf(
			"call Send CoinType: %v error: %v",
			info.CoinType,
			err,
		)
		return nil, err
	}

	signResponse, err := stream.Recv()
	if err != nil {
		logger.Sugar().Errorf(
			"call Recv error: %v",
			err,
		)
		return nil, err
	}

	mpResponse, err := pluginClient.MpoolPush(ctx, &sphinxplugin.MpoolPushRequest{
		Message:   signResponse.GetInfo().GetMessage(),
		Signature: signResponse.GetInfo().GetSignature(),
	})
	if err != nil {
		logger.Sugar().Errorf("call MpoolPush info: %v error: %v", info, err)
		return nil, err
	}
	return mpResponse, nil
}

func ack(
	isOk bool,
	cid string,
	mpResponse *sphinxplugin.MpoolPushResponse,
	tinfo *msgproducer.NotificationTransaction,
) error {
	ackConn, err := grpc2.GetGRPCConn(sconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		logger.Sugar().Errorf(
			"call GetGRPCConn ServiceName: %v error: %v",
			sconst.ServiceName,
			err,
		)
		return err
	}
	ackClient := trading.NewTradingClient(ackConn)
	ctx, cancel := context.WithTimeout(context.Background(), constant.GrpcTimeout)
	defer cancel()
	if isOk && mpResponse.GetInfo() != nil {
		cid = mpResponse.GetInfo().GetCID()
	}
	_, err = ackClient.ACK(ctx, &trading.ACKRequest{
		TransactionIdInsite: tinfo.TransactionIDInsite,
		TransactionChainId:  cid,
		IsOkay:              isOk,
	})
	if err != nil {
		logger.Sugar().Errorf(
			"call ACK ServiceName: %v TranID: %v error: %v",
			sconst.ServiceName,
			tinfo.TransactionIDInsite,
			err,
		)
	}
	return nil
}
