package logging

import (
	"github.com/carrot-systems/csl-logging/src/gateway/rollbar"
)

type Logger struct {
	config LoggerConfig
	engine Usecases
}

func LoadLogger(config LoggerConfig) Logger {
	var engine Usecases

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
