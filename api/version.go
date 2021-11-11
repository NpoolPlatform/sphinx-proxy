// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cv "github.com/NpoolPlatform/go-service-framework/pkg/version"
	"github.com/NpoolPlatform/sphinx-proxy/message/npool/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Version(ctx context.Context, in *emptypb.Empty) (*version.VersionResponse, error) {
	info, err := cv.GetVersion()
	if err != nil {
		logger.Sugar().Errorw("[Version] get service version error: %w", err)
		return &version.VersionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return &version.VersionResponse{Info: info}, nil
}
