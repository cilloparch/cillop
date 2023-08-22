// Package mredis provides a Redis client implementation for Go.
package mredis

import (
	"context"
	"time"
)

// Service is the interface that defines the Redis service methods.
type Service interface {
	// Set sets the value of a key in Redis.
	Set(ctx context.Context, k string, v interface{}) error

	// HSet sets the values of multiple fields in a hash stored at key in Redis.
	HSet(ctx context.Context, k string, v ...interface{}) error

	// Get returns the value of a key in Redis as a string.
	Get(ctx context.Context, k string) (string, error)

	// HGet returns the value associated with field in the hash stored at key in Redis as a string.
	HGet(ctx context.Context, k string, field string) (string, error)

	// HGetAll returns all fields and values of the hash stored at key in Redis as a map[string]string.
	HGetAll(ctx context.Context, k string) (map[string]string, error)

	// SetEx sets the value of a key in Redis with an expiration duration.
	SetEx(ctx context.Context, k string, v interface{}, d time.Duration) error

	// Del deletes one or more keys in Redis.
	Del(ctx context.Context, k ...string) error

	// Exist checks if a key exists in Redis.
	Exist(ctx context.Context, k string) (bool, error)
}

// Config holds the configuration settings for the Redis client.
type Config struct {
	Host     string // Redis server host
	Port     string // Redis server port
	Password string // Redis server password (optional)
	DB       int    // Redis database number
}

// New creates a new Redis client with the provided configuration.
func New(cnf *Config) Service {
	return new(cnf)
}
