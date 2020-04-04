package context

import "github.com/prodbox/weasn/kernel/context/render"

func (c *context) Render(code int, r render.Render) error {
	c.writer.WriteHeader(code)
	return r.Render(c.writer)
}

func (c *context) XML(code int, obj interface{}) error {
	return c.Render(code, render.XML{Data: obj})
}

func (c *context) String(code int, format string, values ...interface{}) error {
	return c.Render(code, render.String{Format: format, Data: values})
}
