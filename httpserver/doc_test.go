package httpserver_test

import (
	"go.uber.org/fx"

	"github.com/lucaskatayama/fx-contrib/httpserver"
)

// This is a simple usage of httpserver.Module with fx
func Example() {
	app := fx.New(
		httpserver.Module,
	)
	app.Run()
}

// Setting HOST and PORT
func ExampleModule_example1() {
	// export HOST=0.0.0.0
	// export PORT=8080
	app := fx.New(
		httpserver.Module,
	)
	app.Run()
}
