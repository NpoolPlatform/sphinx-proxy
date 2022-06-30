package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxproxy"
)

func (s *Server) UpdateTransaction(ctx context.Context, in *sphinxproxy.UpdateTransactionRequest) (out *sphinxproxy.UpdateTransactionResponse, err error) {
	return &sphinxproxy.UpdateTransactionResponse{}, nil
}
