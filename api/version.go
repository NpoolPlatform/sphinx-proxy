//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/version"
	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Version(ctx context.Context, in *emptypb.Empty) (*sphinxproxy.VersionResponse, error) {
	resp, err := version.Version()
	if err != nil {
		logger.Sugar().Errorw("[Version] get service version error: %w", err)
		return &sphinxproxy.VersionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
