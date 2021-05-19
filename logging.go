package logging

import (
	"github.com/carrot-systems/csl-logging/src/config"
	"github.com/carrot-systems/csl-logging/src/gateway/rollbar"
	"github.com/carrot-systems/csl-logging/src/usecases"
)

type Logger struct {
	config config.LoggerConfig
	engine usecases.Logger
}

func LoadLogger(config config.LoggerConfig) Logger {
	var engine usecases.Logger

	switch config.Engine {
	case "rollbar":
		engine = rollbar.CreateLogger(config)
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
