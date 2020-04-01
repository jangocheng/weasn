package context

func (c *context) Query(key string) string {
	value, _ := c.GetQuery(key)
	return value
}

func (c *context) GetQuery(key string) (string, bool) {
	if values, ok := c.GetQueryArray(key); ok {
		return values[0], ok
	}
	return "", false
}

func (c *context) getQueryCache() {
	if c.queryCache == nil {
		c.queryCache = c.request.URL.Query()
	}
}

func (c *context) GetQueryArray(key string) ([]string, bool) {
	c.getQueryCache()
	if values, ok := c.queryCache[key]; ok && len(values) > 0 {
		return values, true
	}
	return []string{}, false
}
