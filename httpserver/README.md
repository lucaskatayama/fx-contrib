# HTTP Server Module

[![PkgGoDev](https://pkg.go.dev/badge/github.com/lucaskatayama/fx-contrib/httpserver)](https://pkg.go.dev/github.com/lucaskatayama/fx-contrib/httpserver)
[![goreport](https://goreportcard.com/badge/github.com/lucaskatayama/fx-contrib/httpserver)](https://goreportcard.com/report/github.com/lucaskatayama/fx-contrib)

## Environment Variables

```
HOST (Default: localhost)
PORT (Default: 5000)
```

## Getting Started

Provide a `http.Handler` for `httpserver` fx Module Simple usage using `echo`

```go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"

	"github.com/lucaskatayama/fx-contrib/httpserver"
)

func router() http.Handler {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	g := e.Group("/api")

	g.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	return e
}

func main() {
	app := fx.New(
		httpserver.Module,
		fx.Provide(router),
	)
	app.Run()
}
```

# HealthCheck Module

## Environment Variables

```
HEALTHCHECK_HEALTHZ_PATH (Default: /healthz)
HEALTHCHECK_READINESS_PATH (Default: /readiness)
```
