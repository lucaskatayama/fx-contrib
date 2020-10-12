package etcd_test

import (
	"context"
	"testing"

	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"

	"github.com/lucaskatayama/fx-contrib/etcd"
)

func TestModule(t *testing.T) {
	assert := assert.New(t)
	var client *clientv3.Client
	app := fxtest.New(t, etcd.Module, fx.Populate(&client))
	app.RequireStart()

	key := "/test"
	value := "hello"
	client.Put(context.Background(), key, value)

	get, err := client.Get(context.Background(), key)
	if err != nil {
		t.FailNow()
	}

	assert.Len(get.Kvs, 1)
	for _, kv := range get.Kvs {
		assert.Equal(key, string(kv.Key))
		assert.Equal(value, string(kv.Value))
	}

}
