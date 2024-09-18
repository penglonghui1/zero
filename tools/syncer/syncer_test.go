package syncer

//
//import (
//	"context"
//	"testing"
//
//	"github.com/pengcainiao2/zero/core/bloom"
//	"github.com/stretchr/testify/assert"
//
//	"github.com/pengcainiao2/zero/core/queue/nsqueue"
//)
//
//func TestNSQ(t *testing.T) {
//	var t1 = QueueCloudDisk(nsqueue.NoOrderNormal)
//	t1.Priority(nsqueue.OrderedNormal)
//}
//
//func TestStores(t *testing.T) {
//	var m []uint8
//	if err := MySQL().Get(context.Background(), &m, "SELECT UNIX_TIMESTAMP()"); err != nil {
//		t.Error(err)
//	}
//	r, err := Redis().Get(context.Background(), "weather:gzs").Result()
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(r)
//}
//
//func TestRedisBitSet_Add(t *testing.T) {
//	filter := bloom.New(Redis(), "test_key", 8092)
//	assert.Nil(t, filter.Add("1426376105656372"))
//	assert.Nil(t, filter.Add("1509307838628047"))
//	ok, err := filter.Exists("1364931407773974")
//	assert.Nil(t, err)
//	assert.True(t, ok)
//}
