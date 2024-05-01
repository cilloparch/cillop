package log

import "fmt"

type Service interface {
	Msg(msg string)
	Error(err error, msg ...string)
	DebugMode() bool
}

type Config struct {
	Debug bool
}

type defaultLogger struct {
	debug bool
}

func Default(cfg Config) Service {
	return &defaultLogger{
		debug: cfg.Debug,
	}
}

func (l *defaultLogger) Msg(msg string) {
	if l.debug {
		fmt.Printf("acorn: %s\n", msg)
	}
}

func (l *defaultLogger) Error(err error, msg ...string) {
	if l.debug {
		if len(msg) > 0 {
			fmt.Printf("acorn: %s: %s\n", msg[0], err.Error())
			return
		}
		l.Msg(err.Error())
	}
}

func (l *defaultLogger) DebugMode() bool {
	return l.debug
}
