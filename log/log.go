package log

import "fmt"

// Service is the interface that wraps the basic logging methods.
type Service interface {

	// Msg logs a message. It should be used for general logging.
	Msg(msg string)

	// Error logs an error. It should be used for logging errors.
	Error(err error, msg ...string)

	// DebugMode returns whether the logger is in debug mode.
	DebugMode() bool
}

type Config struct {
	Debug bool
}

type defaultLogger struct {
	debug bool
}

// Default returns a new logger with the default configuration.
func Default(cfg Config) Service {
	return &defaultLogger{
		debug: cfg.Debug,
	}
}

func (l *defaultLogger) Msg(msg string) {
	if l.debug {
		fmt.Printf("cillop: %s\n", msg)
	}
}

func (l *defaultLogger) Error(err error, msg ...string) {
	if l.debug {
		if len(msg) > 0 {
			fmt.Printf("cillop: %s: %s\n", msg[0], err.Error())
			return
		}
		l.Msg(err.Error())
	}
}

func (l *defaultLogger) DebugMode() bool {
	return l.debug
}
