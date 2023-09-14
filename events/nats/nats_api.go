package nats

import (
	"encoding/json"
	"strings"

	"github.com/cilloparch/cillop/events"
	"github.com/nats-io/nats.go"
)

func splitEvent(fullEvent string) (string, string) {
	splitted := strings.Split(fullEvent, ".")
	return splitted[0], splitted[1]
}

func (e *natsEngine) checkStreamExist(stream string) bool {
	for _, s := range e.config.Streams {
		if s == stream {
			return true
		}
	}
	return false
}

func (e *natsEngine) addStream(stream string) error {
	e.config.Streams = append(e.config.Streams, stream)
	_, err := e.JS.AddStream(&nats.StreamConfig{
		Name:     stream,
		Subjects: []string{stream + ".*"},
	})
	return err
}

func (e *natsEngine) Open() error {
	for _, stream := range e.config.Streams {
		_, err := e.JS.AddStream(&nats.StreamConfig{
			Name:     stream,
			Subjects: []string{stream + ".*"},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *natsEngine) Publish(event string, data interface{}) error {
	p, err := e.Marshal(data)
	if err != nil {
		return err
	}
	return e.C.Publish(event, p)
}

func (e *natsEngine) CheckSubAndPublish(fullEvent string, data interface{}) error {
	channelName, _ := splitEvent(fullEvent)
	if !e.checkStreamExist(channelName) {
		err := e.addStream(channelName)
		if err != nil {
			return err
		}
	}
	return e.Publish(fullEvent, data)
}

func (e *natsEngine) Subscribe(event string, handler events.Handler) error {
	sub, err := e.C.Subscribe(event, func(msg *nats.Msg) {
		handler(msg.Data)
	})
	if err != nil {
		return err
	}
	e.subs = append(e.subs, sub)
	return nil
}

func (e *natsEngine) SubscribeQueue(event string, queue string, handler events.Handler) error {
	sub, err := e.C.QueueSubscribe(event, queue, func(msg *nats.Msg) {
		handler(msg.Data)
		msg.Ack()
	})
	if err != nil {
		return err
	}
	e.subs = append(e.subs, sub)
	return nil
}

func (e *natsEngine) Unsubscribe(event string, handler events.Handler) error {
	for i, sub := range e.subs {
		if sub.Subject == event {
			err := sub.Unsubscribe()
			if err != nil {
				return err
			}
			e.subs = append(e.subs[:i], e.subs[i+1:]...)
		}
	}
	return nil
}

func (e *natsEngine) Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(&data)
}

func (e *natsEngine) Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}

func (e *natsEngine) Close() error {
	for _, sub := range e.subs {
		err := sub.Unsubscribe()
		if err != nil {
			return err
		}
	}
	return nil
}
