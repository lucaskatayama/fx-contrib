package httpserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"

	"github.com/lucaskatayama/fx-contrib/httpserver/healthcheck"
)

// Module provides a `fx` module
// Use on `fx` app declation
//     app := fx.New(
//        httpserver.Module,
//     )
var Module = fx.Options(
	fx.Provide(new),
	fx.Invoke(start),
)

var (
	InvalidHostErr       = errors.New("invalid host")
	InvalidPortErr       = errors.New("invalid port")
	InvaidReadTimeoutErr = errors.New("invalid read timeout")
)

type Params struct {
	fx.In
	Router http.Handler
	Check  *healthcheck.HealthCheck `optional:"true"`
}

func new(params Params) http.Server {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	mux := http.DefaultServeMux
	mux.Handle("/", params.Router)
	if params.Check != nil {
		mux.HandleFunc(params.Check.HandleReadinessCheck())
		mux.HandleFunc(params.Check.HandleHealthzCheck())
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return srv
}

func start(lifecycle fx.Lifecycle, srv http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Printf("running server on %s\n", srv.Addr)
				if err := srv.ListenAndServe(); err != http.ErrServerClosed {
					log.Panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			return srv.Shutdown(ctx)
		},
	})
}
