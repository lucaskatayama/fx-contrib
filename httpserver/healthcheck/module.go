package healthcheck

import (
	"context"
	"net/http"
	"os"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(new),
)

type Checkers struct {
	fx.In
	Checkers []func(ctx context.Context) error `group:"checkers"`
}

type HealthCheck struct {
	Checkers []func(ctx context.Context) error
}

func new(checkers Checkers) *HealthCheck {
	return &HealthCheck{
		Checkers: checkers.Checkers,
	}
}

func (c HealthCheck) Add(check func(ctx context.Context) error) {
	c.Checkers = append(c.Checkers, check)
}

func (c HealthCheck) CheckAll(ctx context.Context) error {
	for _, val := range c.Checkers {
		if err := val(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c HealthCheck) HandleReadinessCheck() (string, http.HandlerFunc) {
	path := os.Getenv("HEALTHCHECK_READINESS_PATH")
	if path == "" {
		path = "/readiness"
	}
	return path, func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if err := c.CheckAll(ctx); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}

}

func (c HealthCheck) HandleHealthzCheck() (string, http.HandlerFunc) {
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
