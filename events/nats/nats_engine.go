package nats

import (
	"github.com/cilloparch/cillop/events"
	"github.com/nats-io/nats.go"
)

type natsEngine struct {
	config Config
	C      *nats.Conn
	JS     nats.JetStreamContext
	subs   []*nats.Subscription
}

type Config struct {
	Url     string
	Streams []string
}

func New(cnf Config) events.Engine {
	nc, js := initDefault(cnf.Url)
	return &natsEngine{
		config: cnf,
		C:      nc,
		JS:     js,
		subs:   []*nats.Subscription{},
	}
}

func initDefault(url string) (*nats.Conn, nats.JetStreamContext) {
	nc := connectNats(url)
	js := connectJS(nc)
	return nc, js
}

func connectNats(url string) *nats.Conn {
	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}
	return nc
}

func connectJS(nc *nats.Conn) nats.JetStreamContext {
	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}
	return js
}
