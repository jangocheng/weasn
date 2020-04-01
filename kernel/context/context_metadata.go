package context

func (c *context) Set(key string, value interface{}) {
	c.keysMutex.Lock()
	if c.keys == nil {
		c.keys = make(map[string]interface{})
	}

	c.keys[key] = value
	c.keysMutex.Unlock()
}

func (c *context) Get(key string) (value interface{}, exists bool) {
	c.keysMutex.RLock()
	value, exists = c.keys[key]
	c.keysMutex.RUnlock()
	return
}
