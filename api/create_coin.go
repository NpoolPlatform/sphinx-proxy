// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/sphinx-proxy/message/npool/signproxy"
)

func (s *Server) CreateCoin(ctx context.Context, in *signproxy.CreateCoinRequest) (*signproxy.CreateCoinResponse, error) {
	return nil, nil
}
