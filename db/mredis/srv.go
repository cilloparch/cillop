package mredis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type client struct {
	c *redis.Client
}

func new(cnf *Config) Service {
	c := redis.NewClient(&redis.Options{
		Addr:     cnf.Host + ":" + cnf.Port,
		Password: cnf.Password,
		DB:       cnf.DB,
	})
	return &client{
		c: c,
	}
}

func (r *client) Set(ctx context.Context, k string, v interface{}) error {
	return r.c.Set(ctx, k, v, 0).Err()
}

func (r *client) HSet(ctx context.Context, k string, v ...interface{}) error {
	return r.c.HSet(ctx, k, v).Err()
}

func (r *client) Get(ctx context.Context, k string) (string, error) {
	return r.c.Get(ctx, k).Result()
}

func (r *client) HGet(ctx context.Context, k string, field string) (string, error) {
	return r.c.HGet(ctx, k, field).Result()
}

func (r *client) HGetAll(ctx context.Context, k string) (map[string]string, error) {
	return r.c.HGetAll(ctx, k).Result()
}

func (r *client) SetEx(ctx context.Context, k string, v interface{}, d time.Duration) error {
	_, err := r.c.Set(ctx, k, v, d).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *client) Del(ctx context.Context, k ...string) error {
	return r.c.Del(ctx, k...).Err()
}

func (r *client) Exist(ctx context.Context, k string) (bool, error) {
	res, err := r.c.Exists(ctx, k).Result()
	if err != nil {
		return false, err
	}
	return res == 1, nil
}
