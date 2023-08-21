package cache

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
)

type (
	Handler[Entity any] func() (Entity, *i18np.Error) // Handler is a function that returns an entity and an error. It's running when the entity is not found in the cache
	Creator[Entity any] func() Entity                 // Creator is a function that returns an entity. It's running when the entity is found in the cache and make a pointer to it
)

// Client is an interface that defines the methods of a cache client
type Client[Entity any] interface {
	// Get returns an entity and an error.
	Get(context.Context, string) (Entity, *i18np.Error)

	// Set sets an entity and returns an error.
	Handler(Handler[Entity]) Client[Entity]

	// Set sets an entity and returns an error.
	Creator(creator Creator[Entity]) Client[Entity]

	// Set sets an entity and returns an error.
	Timeout(time.Duration) Client[Entity]
}

// New returns a new cache client
// It receives a service that implements the Service interface
// and returns a Client interface
func New[Entity any](service Service) Client[Entity] {
	return new[Entity](service)
}
