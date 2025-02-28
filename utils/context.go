package utils

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/szmcdull/go-forceexport"
)

type (
	// 用event模拟的Context,实验性质,请勿使用
	Context struct {
		exitEvent *Event
	}

	// 封装标准库context.WithCancel
	CancelCtx struct {
		context.Context
		cancelFunc context.CancelFunc
		isDone     int32
	}
)

var (
	ContextDoneError = context.Canceled
)

func NewContext() Context {
	return Context{
		exitEvent: NewEvent(),
	}
}

func NewCancelCtx(parent context.Context) *CancelCtx {
	return newCancelCtx(parent)
}

// 将多个Context聚合在一起,任意一个parent Done,聚合Context都会Done
func (me *CancelCtx) NewLinkedCancelCtx(contexts ...context.Context) *CancelCtx {
	count := len(contexts)
	if count == 0 {
		panic(`at least 1 ctx expected`)
	}

	withCancel := newCancelCtx(me)
	for i := 0; i < count; i++ {
		propagateCancel(contexts[i], withCancel)
	}

	return withCancel
}

type (
	_LinkedCancelCtx struct {
		CancelCtx
		parent *CancelCtx
	}

	cancelCtx = struct{}
	canceler  interface {
		cancel(removeFromParent bool, err error)
		Done() <-chan struct{}
	}
)

var (
	propagateCancel func(parent context.Context, child canceler)
)

func (c Context) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (c Context) Done() <-chan struct{} {
	return c.exitEvent.Done()
}

func (c Context) Err() error {
	if c.exitEvent.IsSet() {
		return ContextDoneError
	} else {
		return nil
	}
}

func (c Context) Value(key interface{}) interface{} {
	return nil
}

func (c Context) Close() {
	c.exitEvent.Set()
}

func (me *CancelCtx) Cancel() bool {
	if atomic.CompareAndSwapInt32(&me.isDone, 0, 1) {
		me.cancelFunc()
		return true
	}
	return false
}

func (me *CancelCtx) cancel(removeFromParent bool, err error) {
	if atomic.CompareAndSwapInt32(&me.isDone, 0, 1) {
		me.cancelFunc()
	}
}

func (me *CancelCtx) Err() error {
	if me.isDone != 0 {
		return ContextDoneError
	} else {
		return me.Context.Err()
	}
}

func init() {
	if err := forceexport.GetFunc(&propagateCancel, `context.propagateCancel`); err != nil {
		panic(err)
	}
}

func newCancelCtx(parent context.Context) *CancelCtx {
	c, f := context.WithCancel(parent)
	return &CancelCtx{
		Context:    c,
		cancelFunc: f,
	}
}
