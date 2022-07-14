package api

import (
	"context"
	"errors"
	"io"
	"math/rand"
	"sync"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	scconst "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
!important
now record sign and plugin conn in service memory, cause not start multi pod,
next we can record the connect in db or other service(eg: redis), then
we can start multi pod
*/

type lmPluginType map[sphinxplugin.CoinType][]*mPlugin

var (
	ErrNoSignServiceFound   = errors.New("no sign service conn found")
	ErrNoPluginServiceFound = errors.New("no plugin service conn found")
)

var (
	// rand stream client
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	slk            sync.RWMutex
	plk            sync.RWMutex
	lmSign         = make([]*mSign, 0)
	lmPlugin       = make(lmPluginType)
	channelBufSize = 100
)

func getProxySign() (*mSign, error) {
	slk.RLock()
	defer slk.RUnlock()
	logger.Sugar().Infof("get proxy sign length: %v", len(lmSign))
	if len(lmSign) == 0 {
		return nil, ErrNoSignServiceFound
	}
	return lmSign[rnd.Intn(len(lmSign))], nil
}

func getProxyPlugin(coinType sphinxplugin.CoinType) (*mPlugin, error) {
	plk.RLock()
	defer plk.RUnlock()
	logger.Sugar().Infof("get coin %v proxy plugin length: %v", coinType, len(lmPlugin[coinType]))
	if len(lmPlugin[coinType]) == 0 {
		return nil, ErrNoPluginServiceFound
	}
	return lmPlugin[coinType][rnd.Intn(len(lmPlugin[coinType]))], nil
}

func haveCoin(name string) (bool, error) {
	ackConn, err := grpc2.GetGRPCConn(scconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return false, err
	}
	defer ackConn.Close()

	client := coininfo.NewSphinxCoinInfoClient(ackConn)

	ctx, cancel := context.WithTimeout(context.Background(), constant.GrpcTimeout)
	defer cancel()

	// define in plugin
	ret, err := client.GetCoinInfos(ctx, &coininfo.GetCoinInfosRequest{
		Name: name,
	})
	if err != nil {
		return false, err
	} else if ret.Total > 0 {
		return true, nil
	}
	return false, nil
}

func registerCoin(coinInfo *coininfo.CreateCoinInfoRequest) error {
	ackConn, err := grpc2.GetGRPCConn(scconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return err
	}
	defer ackConn.Close()

	client := coininfo.NewSphinxCoinInfoClient(ackConn)

	ctx, cancel := context.WithTimeout(context.Background(), constant.GrpcTimeout)
	defer cancel()

	// define in plugin
	_, err = client.CreateCoinInfo(ctx, coinInfo)
	return err
}

func checkCode(err error) bool {
	if err == io.EOF ||
		status.Code(err) == codes.Unavailable ||
		status.Code(err) == codes.Canceled {
		return true
	}
	return false
}
