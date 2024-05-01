package cqrs

import (
	"context"
)

// Handler is the interface that must be implemented by a command or query handler
// It is used to execute a command or query
// The first parameter is the context
// The second parameter is the command or query
// The return value is the response and an error
// The error is an error
// The error is nil if the command or query is executed successfully
// The error is not nil if the command or query is not executed successfully
type Handler[TParams any, TResponse any] interface {
	Handle(context.Context, TParams) (TResponse, error)
}

// HandlerFunc is a function that can be used as a Handler
// It is used to convert a function to a Handler
type HandlerFunc[TParams any, TResponse any] func(context.Context, TParams) (TResponse, error)
