package mongodb_test

import (
	"go.uber.org/fx"

	"github.com/lucaskatayama/fx-contrib/mongodb"
)

// This is a simple usage of mongodb.Module with fx
func Example() {
	// export MONGODB_DATABASE=db
	app := fx.New(
		mongodb.Module,
	)
	app.Run()
}

// Setting MONGODB_URI
func Example_example1() {
	// export MONGODB_URI=mongodb://localhost:27017
	app := fx.New(
		mongodb.Module,
	)
	app.Run()
}
