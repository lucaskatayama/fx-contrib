package sql

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"go.uber.org/fx"
)

// Module provides a sql module for fx.
var Module = fx.Options(
	fx.Provide(new),
	fx.Invoke(start),
)

func new() (*sql.DB, error) {
	driver := os.Getenv("SQL_DRIVER")
	prefix := strings.ToUpper(driver)
	uri := os.Getenv(fmt.Sprintf("%s_URI", prefix))
	return sql.Open(driver, uri)
}

func start(lifecycle fx.Lifecycle, db *sql.DB) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.PingContext(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})
}
