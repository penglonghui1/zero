package discov

import (
	"context"
	"testing"

	"github.com/pengcainiao2/zero/core/logx"
)

const authorizationFree = "/flyele/configs/user/authentication-free"

var (
	flyeleUserPhoneNumber []string
)

func TestEtcdClient_LoadOrStore(t1 *testing.T) {
	etcd := newEtcdClient()
	resp := etcd.LoadOrStore(authorizationFree, "")
	_ = resp.JSON(&flyeleUserPhoneNumber)
	if len(flyeleUserPhoneNumber) == 0 {
		logx.NewTraceLogger(context.Background()).Warn().Msg("authorizationFree is NULL")
	}

	go etcd.WatchKey(authorizationFree, func(event EtcdChangeType, value string) {
		_ = GetResponse(value).JSON(&flyeleUserPhoneNumber)
	})
}
