package zap

import (
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module provides zap module for fx.
var Module = fx.Options(
	fx.Provide(newLogger),
	fx.Provide(newSugared),
)

func newLogger() (*zap.Logger, error) {
	env := os.Getenv("ENV")
	logger, err := zap.NewProduction()
	if env == "DEV" {
		logger, err = zap.NewDevelopment()
	}

	defer logger.Sync()
	return logger, err
}

func newSugared(logger *zap.Logger) *zap.SugaredLogger {
	return logger.Sugar()
}
