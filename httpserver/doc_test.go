package httpserver_test

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/lucaskatayama/fx-contrib/httpserver"
)

// This is a simple usage of httpserver.Module with fx.
// You must provide a http.Handler to the module.
func Example() {
	app := fx.New(
		httpserver.Module,
		fx.Provide(func() http.Handler {
			mux := http.NewServeMux()
			mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello"))
				return
			})
			return mux
		}),
	)
	app.Run()
}

// Setting HOST and PORT
func ExampleModule_example1() {
	// export HOST=0.0.0.0
	// export PORT=8080

	app := fx.New(
		httpserver.Module,
		fx.Provide(func() http.Handler {
			mux := http.NewServeMux()
			mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello"))
				return
			})
			return mux
		}),
	)
	app.Run()
}
