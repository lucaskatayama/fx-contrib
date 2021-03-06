package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
)

// Module provides a fx module
//
//
// Use on fx app declation
//     app := fx.New(
//        httpserver.Module,
//     )
var Module = fx.Options(
	fx.Provide(new),
	fx.Invoke(start),
)

type params struct {
	fx.In
	Router http.Handler
	Check  *healthcheck `optional:"true"`
}

func new(params params) http.Server {
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
		mux.HandleFunc(params.Check.handleReadinessCheck())
		mux.HandleFunc(params.Check.handleHealthzCheck())
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
