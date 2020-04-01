package context

import (
	"net/http"
	"net/url"
	"sync"
)

type context struct {
	opts       Options
	writer     http.ResponseWriter
	request    *http.Request
	keysMutex  *sync.RWMutex
	keys       map[string]interface{}
	queryCache url.Values
}

func NewContext(opts Options) *context {
	return &context{
		opts:      opts,
		keysMutex: &sync.RWMutex{},
	}
}

func (c *context) Options() Options {
	return c.opts
}

func (c *context) Writer() http.ResponseWriter {
	return c.writer
}

func (c *context) Request() *http.Request {
	return c.request
}

func (c *context) Reset(w http.ResponseWriter, r *http.Request) {
	c.writer = w
	c.request = r
	c.keysMutex = &sync.RWMutex{}
	c.keys = nil
	c.queryCache = nil
}
