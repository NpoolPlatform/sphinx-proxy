package version

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/version"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
)

func Version() (*sphinxproxy.VersionResponse, error) {
	info, err := version.GetVersion()
	if err != nil {
		logger.Sugar().Errorf("get service version error: %+w", err)
		return nil, fmt.Errorf("get service version error: %w", err)
	}
	return &sphinxproxy.VersionResponse{
		Info: info,
	}, nil
}
