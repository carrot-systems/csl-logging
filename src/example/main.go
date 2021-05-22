package main

import (
	"errors"
	logging "github.com/carrot-systems/csl-logging"
)

func main() {
	log := logging.LoadLogger(logging.LoggerConfig{
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
