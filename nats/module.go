package nats

import (
	"context"
	"os"

	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
)

// Module provides a nats module for fx
var Module = fx.Options(
	fx.Provide(new),
	fx.Invoke(start),
)

type params struct {
	fx.In
	Opts []nats.Option `optional:"true"`
}

func new(p params) (*nats.Conn, error) {
	ncURI := os.Getenv("NATS_URI")
	if ncURI == "" {
		ncURI = "nats://localhost:4222"
	}
	return nats.Connect(ncURI, p.Opts...)
}

func start(lifecycle fx.Lifecycle, nc *nats.Conn) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			nc.Close()
			return nil
		},
	})
}
