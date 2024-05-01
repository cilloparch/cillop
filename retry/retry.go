package retry

import (
	"time"

	"github.com/cilloparch/cillop/log"
)

// RetryFunc is a function that will be retried
// if it returns an error
type RetryFunc func() error

// Config is the configuration for the retry package
type Config struct {

	// MaxRetries is the maximum number of retries
	MaxRetries int

	// WaitTime is the time to wait between retries
	WaitTime time.Duration

	// Logger is the logger to use
	Logger log.Service
}

// DefaultConfig is the default configuration for the retry package
var DefaultConfig = Config{
	MaxRetries: 3,
	WaitTime:   1 * time.Second,
	Logger:     log.Default(log.Config{Debug: true}),
}

// Run runs the given RetryFunc with the given Config
func Run(fn RetryFunc, cfg Config) error {
	for {
		err := fn()
		if err == nil {
			break
		}
		cfg.MaxRetries--
		if cfg.MaxRetries == 0 {
			return err
		}
		if cfg.Logger != nil && cfg.Logger != log.Service(nil) {
			cfg.Logger.Error(err, "retrying...")
		}
		time.Sleep(cfg.WaitTime)
	}
	return nil
}
