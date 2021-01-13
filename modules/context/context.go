package context

import (
	"context"
	"github.com/google/uuid"
)

var (
	_ context.Context = Context{}
	Generator generator = createUUID
)

type generator func() string

type Context struct {
	context.Context
	tracer Tracer
}

func (c *Context) TraceId() string {
	return c.tracer.Id
}

func WithTracer(ctx context.Context) *Context {
	return &Context{
		Context: ctx,
		tracer: NewTracer(Generator()),
	}
}

func createUUID() string {
	return uuid.New().String()
}