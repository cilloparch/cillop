package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cilloparch/cillop/i18np"
)

type client[Entity any] struct {
	handler Handler[Entity]
	creator Creator[Entity]
	service Service
	timeout time.Duration
}

func new[Entity any](service Service) Client[Entity] {
	return &client[Entity]{
		service: service,
		timeout: 5 * time.Minute,
	}
}

func (c *client[Entity]) Handler(handler Handler[Entity]) Client[Entity] {
	c.handler = handler
	return c
}

func (c *client[Entity]) Creator(creator Creator[Entity]) Client[Entity] {
	c.creator = creator
	return c
}

func (c *client[Entity]) Timeout(timeout time.Duration) Client[Entity] {
	c.timeout = timeout
	return c
}

func (c *client[Entity]) Get(ctx context.Context, key string) (Entity, error) {
	if !c.validate() {
		var e Entity
		return e, i18np.NewError(errorMessages.NotRunnable)
	}
	e := c.creator()
	isExist, err := c.service.Exist(ctx, key)
	if err != nil {
		return e, i18np.NewError(errorMessages.AnErrorOnExist)
	}
	if isExist {
		return c.get(ctx, key, e)
	}
	return c.handleAndSet(ctx, key, e)
}

func (c *client[Entity]) validate() bool {
	return c.handler != nil && c.creator != nil
}

func (c *client[Entity]) get(ctx context.Context, key string, e Entity) (Entity, error) {
	bytes, _err := c.service.Get(ctx, key)
	if _err != nil {
		return e, i18np.NewError(errorMessages.AnErrorOnGet)
	}
	return c.unmarshal(bytes)
}

func (c *client[Entity]) handleAndSet(ctx context.Context, key string, e Entity) (Entity, error) {
	entity, _error := c.handler()
	if _error != nil {
		return e, _error
	}
	err := c.service.SetEx(ctx, key, c.marshal(entity), c.timeout)
	if err != nil {
		return e, i18np.NewError(errorMessages.AnErrorOnSet)
	}
	return entity, nil
}

func (c *client[Entity]) unmarshal(bytes string) (Entity, error) {
	var entity Entity
	err := json.Unmarshal([]byte(bytes), &entity)
	if err != nil {
		return entity, i18np.NewError(errorMessages.AnErrorOnGet)
	}
	return entity, nil
}

func (c *client[Entity]) marshal(entity Entity) interface{} {
	bytes, _ := json.Marshal(entity)
	return bytes
}
