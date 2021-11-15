package api

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-proxy/message/npool/signproxy"
	msgcli "github.com/NpoolPlatform/sphinx-proxy/pkg/message/client"
)

func (s *Server) Transaction(stream signproxy.SignProxy_TransactionServer) error {
	logger.Sugar().Infof("consume template example")
	consumerInfo, err := msgcli.ConsumeExample()
	if err != nil {
		logger.Sugar().Errorf("fail to consume example: %v", err)
		return err
	}

	for info := range consumerInfo {
		// if use go routine should limit go routine numbers
		logger.Sugar().Infof("consumer info: ", info.Body)
		if err := stream.Send(&signproxy.TransactionRequest{}); err != nil {
			logger.Sugar().Infof("consumer info: ", info.Body)
			continue
		}

		resp, err := stream.Recv()
		if err != nil {
			logger.Sugar().Infof("consumer info: ", info.Body)
			continue
		}

		logger.Sugar().Infof("", resp)
	}

	return nil
}
