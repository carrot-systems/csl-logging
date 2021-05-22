package logging

import (
	"github.com/carrot-systems/csl-logging/src/gateway/rollbar"
	"github.com/carrot-systems/csl-logging/src/usecases"
)

type Logger struct {
	config LoggerConfig
	engine usecases.Usecases
}

func LoadLogger(config LoggerConfig) Logger {
	var engine usecases.Usecases

	switch config.Engine {
	case "rollbar":
		engine = rollbar.CreateLogger(rollbar.RollbarConfig{
			Token:           config.Token,
			EnvironmentType: config.EnvironmentType,
			Version:         config.Version,
			Hostname:        config.Hostname,
		})
		break
	}

	return Logger{
		config: config,
		engine: engine,
	}
}

func (l Logger) LogError(err error) {
	l.engine.Log(err)
}
