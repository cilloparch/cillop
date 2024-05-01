package chain

import (
	"context"
)

// Handler is a function that handles a request.
// It returns a result and an error.
// If the error is not nil, the chain will be stopped.
type Handler[Params any, Result any] func(ctx context.Context, params Params) (*Result, error)

// Chain is a chain of handlers.
// It can be used to execute a series of handlers in order.
// The result of the previous handler is passed to the next handler.
// If an error occurs, the chain will be stopped.
type Chain[Params any, Result any] interface {
	// Use adds a handler to the chain.
	// The handler will be executed in the order of use.
	// If the handler is nil, it will be ignored.
	// example:
	// chain.Use(handler1, handler2, handler3)
	Use(handler ...Handler[Params, Result]) Chain[Params, Result]

	// Run starts the chain.
	// It will execute all handlers in the chain.
	// If an error occurs, the chain will be stopped.
	// example:
	// result, err := chain.Run(ctx, params)
	Run(ctx context.Context, params Params) (*Result, error)

	// RunErr starts the chain.
	// It will execute all handlers in the chain.
	// If an error occurs, the chain will be stopped.
	// example:
	// result, err := chain.RunErr(ctx, params, err)
	RunErr(ctx context.Context, params Params, err error) (*Result, error)
}

type chain[Params any, Result any] struct {
	handlers []Handler[Params, Result]
}

// Make creates a new chain.
// example:
// chain := chain.Make[Params, Result]()
// chain.Use(handler1, handler2, handler3)
// result, err := chain.Start(ctx, params)
func New[Params any, Result any]() Chain[Params, Result] {
	return &chain[Params, Result]{}
}

// Use adds a handler to the chain.
// The handler will be executed in the order of use.
// If the handler is nil, it will be ignored.
// example:
// chain.Use(handler1, handler2, handler3)
func (c *chain[Params, Result]) Use(handler ...Handler[Params, Result]) Chain[Params, Result] {
	c.handlers = append(c.handlers, handler...)
	return c
}

// Run starts the chain.
// It will execute all handlers in the chain.
// If an error occurs, the chain will be stopped.
// example:
// result, err := chain.Run(ctx, params)
func (c *chain[Params, Result]) Run(ctx context.Context, params Params) (*Result, error) {
	var result *Result
	var err error
	for _, handler := range c.handlers {
		result, err = handler(ctx, params)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

// RunErr starts the chain.
// It will execute all handlers in the chain.
// If an error occurs, the chain will be stopped.
// example:
// result, err := chain.RunErr(ctx, params, err)
func (c *chain[Params, Result]) RunErr(ctx context.Context, params Params, err error) (*Result, error) {
	if err != nil {
		return nil, err
	}
	return c.Run(ctx, params)
}
