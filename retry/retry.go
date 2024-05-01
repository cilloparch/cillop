package retry

import (
	"time"

	"github.com/cilloparch/cillop/log"
)

type RetryFunc func() error

type Config struct {
	MaxRetries int
	WaitTime   time.Duration
	Logger     log.Service
}

var DefaultConfig = Config{
	MaxRetries: 3,
	WaitTime:   1 * time.Second,
	Logger:     log.Default(log.Config{Debug: true}),
}

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
