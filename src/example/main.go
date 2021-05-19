package main

import (
	"errors"
	logging "github.com/carrot-systems/csl-logging"
	config "github.com/carrot-systems/csl-logging/src/config"
)

func main() {
	log := logging.LoadLogger(config.LoggerConfig{
		Engine:          "rollbar",
		Token:           "",
		EnvironmentType: "development",
		Version:         "1.0",
		Hostname:        "testmain",
	})

	check(log)
}

func check(logger logging.Logger) {
	logger.LogError(errors.New("no user id"))
}
