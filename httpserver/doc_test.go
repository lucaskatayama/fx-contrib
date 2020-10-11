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
