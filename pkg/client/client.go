package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/sphinxproxy"

	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get balance connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewSphinxProxyClient(conn)

	return fn(_ctx, cli)
}

func GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (*npool.BalanceInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error) {
		resp, err := cli.GetBalance(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail get balances: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get balances: %v", err)
	}
	return info.(*npool.BalanceInfo), nil
}
