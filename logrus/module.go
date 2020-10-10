package logrus

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(new),
)

func new() *logrus.Logger {
	loglevel := os.Getenv("LOGLEVEL")
	lvl, err := logrus.ParseLevel(loglevel)
	if err != nil {
		log.Println("failed to parse LOGLEVEL")
	}
	logrus.SetLevel(lvl)
	return logrus.StandardLogger()
}
