package api

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	coinconst "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/stretchr/testify/assert"
)

func TestRegisterCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	coinConn, err := grpc2.GetGRPCConn(coinconst.ServiceName, grpc2.GRPCTAG)
	if assert.Nil(t, err) {
		t.Errorf("Coin Service Not Running: %v", err)
	}

	coinClinet := coininfo.NewSphinxCoininfoClient(coinConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err = coinClinet.GetCoinInfo(ctx, &coininfo.GetCoinInfoRequest{
		CoinType: sphinxplugin.CoinType_CoinTypeFIL,
	})
	if assert.Nil(t, err) {
		_, err = coinClinet.RegisterCoin(ctx, &coininfo.RegisterCoinRequest{
			CoinType: sphinxplugin.CoinType_CoinTypeFIL,
		})
		assert.Nil(t, err)
	} else {
		_, err = coinClinet.RegisterCoin(ctx, &coininfo.RegisterCoinRequest{
			CoinType: sphinxplugin.CoinType_CoinTypeFIL,
		})
		assert.NotNil(t, err)
	}
}
