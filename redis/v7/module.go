package redis

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v7"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewClient), fx.Invoke(Init), fx.Provide(Check))

func NewClient() *redis.Client {
	addr := os.Getenv("REDIS_URI")

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

type CheckOut struct {
	fx.Out
	Check func(ctx context.Context) error `group:"checkers"`
}

func Check(client *redis.Client) CheckOut {
	return CheckOut{
		Check: func(ctx context.Context) error {
			cmd := client.WithTimeout(3 * time.Second).Ping()
			if _, err := cmd.Result(); err != nil {
				return err
			}
			return nil
		},
	}
}

func Init(lifecycle fx.Lifecycle, client *redis.Client) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			cmd := client.Ping()
			_, err := cmd.Result()
			return err
		},
		OnStop: func(ctx context.Context) error {
			cmd := client.Shutdown()
			_, err := cmd.Result()
			return err
		},
	})
}
