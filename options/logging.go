package options

import (
	"github.com/sirupsen/logrus"
)

// Logging store configuration about log level
var Logging = struct {
	Level string `long:"log-level" env:"LOGGING_LEVEL" default:"info" choice:"trace" choice:"info" choice:"debug" choice:"warn" choice:"error" choice:"fatal"`
}{}

func configLogging() {
	lvl, err := logrus.ParseLevel(Logging.Level)
	if err != nil {
		logrus.Panicf("error parsing log level '%s': %v", Logging.Level, err)
	}

	logrus.SetLevel(lvl)
}

func init() {
	registry[group{"logging", "Logging configuration"}] = &Logging
}
