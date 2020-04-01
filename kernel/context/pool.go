package context

import (
	"net/http"
	"sync"
)

// context临时对象池
type Pool struct {
	pool    *sync.Pool
	newFunc func() Context
}

func New(newFunc func() Context) *Pool {
	p := &Pool{pool: &sync.Pool{}, newFunc: newFunc}
	p.pool.New = func() interface{} { return p.newFunc() }
	return p
}

func (p *Pool) Attach(newFunc func() Context) {
	p.newFunc = newFunc
}

// Acquire 获取context
func (p *Pool) Acquire(w http.ResponseWriter, r *http.Request) Context {
	c := p.pool.Get().(Context)
	c.Reset(w, r)
	return c
}

// Release 释放context
func (c *Pool) Release(ctx Context) {
	c.pool.Put(ctx)
}
