package client

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/sphinxproxy"

	servicename "github.com/NpoolPlatform/sphinx-proxy/pkg/servicename"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewSphinxProxyClient(conn)

	return fn(_ctx, cli)
}

func GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (*npool.BalanceInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error) {
		resp, err := cli.GetBalance(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.BalanceInfo), nil
}

func CreateTransaction(ctx context.Context, in *npool.CreateTransactionRequest) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error) {
		_, err := cli.CreateTransaction(ctx, in)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateTransaction(ctx context.Context, in *npool.UpdateTransactionRequest) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error) {
		_, err := cli.UpdateTransaction(ctx, in)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetTransaction(ctx context.Context, id string) (*npool.TransactionInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error) {
		resp, err := cli.GetTransaction(ctx, &npool.GetTransactionRequest{
			TransactionID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.TransactionInfo), nil
}

func CreateAddress(ctx context.Context, coinName string) (*npool.WalletInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.SphinxProxyClient) (cruder.Any, error) {
		resp, err := cli.CreateWallet(ctx, &npool.CreateWalletRequest{
			Name: coinName,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.WalletInfo), nil
}
