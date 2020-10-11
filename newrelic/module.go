package newrelic

import (
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
)

// Module provides a newrelic module for fx.
var Module = fx.Options(
	fx.Provide(new),
)

func new() (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("NEWRELIC_APPNAME")),
		newrelic.ConfigLicense(os.Getenv("NEWRELIC_LICENSEKEY")),
	)

}
