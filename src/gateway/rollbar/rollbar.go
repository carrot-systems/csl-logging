package rollbar

import (
	"github.com/carrot-systems/csl-logging/src/config"
	"github.com/carrot-systems/csl-logging/src/usecases"
	rb "github.com/rollbar/rollbar-go"
)

type rollbar struct {
	config config.LoggerConfig
}

func (r rollbar) Log(err error) {
	rb.Critical(err)
	rb.Wait()
}

func CreateLogger(config config.LoggerConfig) usecases.Logger {
	rb.SetToken(config.Token)
	rb.SetEnvironment(config.EnvironmentType) // defaults to "development"
	rb.SetCodeVersion(config.Version)         // optional Git hash/branch/tag (required for GitHub integration)
	rb.SetServerHost(config.Hostname)         // optional override; defaults to hostname
	rb.SetServerRoot(GetModuleName())         // path of project (required for GitHub integration and non-project stacktrace collapsing)

	return rollbar{
		config: config,
	}
}
