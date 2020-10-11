package httpserver

import (
	"context"
	"net/http"
	"os"

	"go.uber.org/fx"
)

// ModuleHealthCheck provides a healthcheck module for fx.
var ModuleHealthCheck = fx.Options(
	fx.Provide(newHealthCheck),
)

type Checkers struct {
	fx.In
	Checkers []func(ctx context.Context) error `group:"checkers"`
}

type healthcheck struct {
	Checkers []func(ctx context.Context) error
}

func newHealthCheck(checkers Checkers) *healthcheck {
	return &healthcheck{
		Checkers: checkers.Checkers,
	}
}

func (c healthcheck) checkAll(ctx context.Context) error {
	for _, val := range c.Checkers {
		if err := val(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c healthcheck) handleReadinessCheck() (string, http.HandlerFunc) {
	path := os.Getenv("HEALTHCHECK_READINESS_PATH")
	if path == "" {
		path = "/readiness"
	}
	return path, func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if err := c.checkAll(ctx); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}

}

func (c healthcheck) handleHealthzCheck() (string, http.HandlerFunc) {
	path := os.Getenv("HEALTHCHECK_HEALTHZ_PATH")
	if path == "" {
		path = "/healthz"
	}
	return path, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}

}
