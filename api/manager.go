package api

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"

	coincli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	coinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
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

func getProxySign(name string) (*mSign, error) {
	slk.RLock()
	defer slk.RUnlock()

	if len(lmSign) == 0 {
		return nil, ErrNoSignServiceFound
	}

	signs := lmSign[0].ctype
	for _, s := range lmSign[1:] {
		signs += " | " + s.ctype
	}
	logger.Sugar().Infof("signs: %v | %v", signs, len(lmSign))

	if len(name) > 0 {
		for _, s := range lmSign {
			if strings.EqualFold(s.ctype, name) {
				return s, nil
			}
		}
		logger.Sugar().Infof("fallback: %v", name)
	}

	lSigns := make([]*mSign, 0)
	for _, s := range lmSign {
		if s.ctype != "" {
			continue
		}
		lSigns = append(lSigns, s)
	}

	if len(lSigns) == 0 {
		return nil, ErrNoSignServiceFound
	}

	return lSigns[rnd.Intn(len(lSigns))], nil
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
	_, total, err := coincli.GetCoins(context.Background(), &coinpb.Conds{
		Name: &v1.StringVal{
			Op:    cruder.EQ,
			Value: name,
		},
	}, 0, 1)
	if err != nil {
		return false, err
	} else if total > 0 {
		return true, nil
	}
	return false, nil
}

func registerCoin(coinInfo *coinpb.CoinReq) error {
	_, err := coincli.CreateCoin(context.Background(), coinInfo)
	return err
}
