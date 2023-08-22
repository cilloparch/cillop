package events

type Handler func(data []byte)

type Engine interface {
	Open() error
	Publish(event string, data interface{}) error
	CheckSubAndPublish(fullEvent string, data interface{}) error
	Subscribe(event string, handler Handler) error
	Unsubscribe(event string, handler Handler) error
	Marshal(data interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Close() error
}

type Publisher interface {
	Publish(event string, data interface{}) error
	CheckSubAndPublish(fullEvent string, data interface{}) error
	Marshal(data interface{}) ([]byte, error)
}

type Subscriber interface {
	Subscribe(event string, handler Handler) error
	Unsubscribe(event string, handler Handler) error
	Unmarshal(data []byte, v interface{}) error
}
