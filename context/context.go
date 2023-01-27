package contextext

import (
	"context"
	"time"
)

var _ context.Context = (*detachedContext)(nil)

type detachedContext struct {
	parent context.Context
}

func Detach(parent context.Context) detachedContext {
	return detachedContext{parent: parent}
}

func (detachedContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (detachedContext) Done() <-chan struct{} {
	return nil
}

func (c detachedContext) Err() error {
	return c.parent.Err()
}

func (c detachedContext) Value(key interface{}) interface{} {
	return c.parent.Value(key)
}
