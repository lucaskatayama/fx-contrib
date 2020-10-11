package nats

import (
	"context"
	"errors"
	"os"

	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
)

var (
	ErrNatsURIRequired = errors.New("NATS_URI is required")
)

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
		return nil, ErrNatsURIRequired
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
