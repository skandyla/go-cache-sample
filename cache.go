package cache

type Cache struct {
	items map[string]interface{}
}

func (c *Cache) Get(key string) interface{} {
	return c.items[key]
}

func (c *Cache) Set(key string, value interface{}) error {
	c.items[key] = value
	// not sure is it a right way to handle errors for methods ?
	return nil
}

func (c *Cache) Delete(key string) error {
	delete(c.items, key)
	return nil
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}
