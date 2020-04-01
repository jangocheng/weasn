package context

import "github.com/prodbox/weasn/kernel/context/render"

func (c *context) Render(code int, r render.Render) {
	c.writer.WriteHeader(code)
	if err := r.Render(c.writer); err != nil {
		panic(err)
	}
}

func (c *context) XML(code int, obj interface{}) {
	c.Render(code, render.XML{Data: obj})
}

func (c *context) String(code int, format string, values ...interface{}) {
	c.Render(code, render.String{Format: format, Data: values})
}
