package mongodb

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/fx"
)

var (
	ErrDBNameRequired = errors.New("MONGODB_DATABASE is required")
	ErrURIRequired    = errors.New("MONGODB_URI is required")
)

var Module = fx.Options(
	fx.Provide(newClient),
	fx.Provide(newDatabase),
	fx.Provide(check),
	fx.Invoke(start),
)

func newClient() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, ErrURIRequired
	}
	return mongo.NewClient(options.Client().ApplyURI(uri))
}

func newDatabase(client *mongo.Client) (*mongo.Database, error) {
	name := os.Getenv("MONGODB_DATABASE")
	if name == "" {
		return nil, ErrDBNameRequired
	}
	return client.Database(name), nil
}

type CheckOut struct {
	fx.Out
	Check func(ctx context.Context) error `group:"checkers"`
}

func check(client *mongo.Client) CheckOut {
	return CheckOut{
		Check: func(ctx context.Context) error {
			ctxTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()
			if err := client.Ping(ctxTimeout, readpref.Primary()); err != nil {
				return err
			}
			return nil
		},
	}
}

func start(lifecycle fx.Lifecycle, client *mongo.Client) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := client.Connect(ctx); err != nil {
				return err
			}
			if err := client.Ping(ctx, readpref.Primary()); err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return client.Disconnect(ctx)
		},
	})
}
