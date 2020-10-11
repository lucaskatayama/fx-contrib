# FX Contrib Modules

This repo has a bunch of Uber/FX modules

## Getting Started

To use, register the chosen module on FX app initialization:

```go
package main

import (
	"github.com/lucaskatayama/fx-contrib/httpserver"
	"github.com/lucaskatayama/fx-contrib/httpserver/healthcheck"
	"github.com/lucaskatayama/fx-contrib/mongodb"
	"github.com/lucaskatayama/fx-contrib/redis/v7"
	"github.com/lucaskatayama/fx-contrib/sql"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		httpserver.Module,
		logrus.Module,
		mongodb.Module,
		healthcheck.Module,
		redis.Module,
		sql.Module,
	)

	app.Run()
}
```

## Modules

- [httpserver](./httpserver/README.md)
- [nats](./nats/README.md)
- [newrelic](./newrelic/README.md)
- [redis](./redis/README.md)
- [sql](./sql/README.md)
