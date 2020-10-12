package etcd

import (
	"os"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"go.uber.org/fx"
)

// Module provides etcd module for fx.
var Module = fx.Options(
	fx.Provide(new),
)

func new() (*clientv3.Client, error) {
	uri := os.Getenv("ETCD_URI")
	if uri == "" {
		uri = "localhost:2379"
	}
	servers := strings.Split(uri, ",")
	timeout := os.Getenv("ETCD_DIAL_TIMEOUT")
	if timeout == "" {
		timeout = "3s"
	}
	dialTimeout, err := time.ParseDuration(timeout)
	if err != nil {
		return nil, err
	}
	return clientv3.New(clientv3.Config{
		Endpoints:   servers,
		DialTimeout: dialTimeout,
	})
}
