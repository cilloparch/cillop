package cache

import (
	"context"
	"time"
)

// Service is an interface that defines the methods of a cache service
// It's implemented by the cache service
type Service interface {
	// Get returns a value and an error.
	Get(ctx context.Context, k string) (string, error)

	// Set sets a value and returns an error.
	SetEx(ctx context.Context, k string, v interface{}, d time.Duration) error

	// Set sets a value and returns an error.
	Exist(ctx context.Context, k string) (bool, error)
}
