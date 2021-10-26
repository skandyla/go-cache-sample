package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Item struct {
	name interface{}
	ttl  time.Duration
}

type Cache struct {
	items map[string]Item
	mu    *sync.Mutex
	close chan bool // close cache cache.close <- true
}

// Get - get item from cache
func (c *Cache) Get(key string) (interface{}, error) {
	value, ok := c.items[key]
	if !ok {
		return nil, errors.New("item not exist")
	}
	return value.name, nil
}

// Set - set new item into cache
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.items[key] = Item{value, ttl}
	c.mu.Unlock()
}

// Delete - delete item from cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.Unlock()
}

// New - return new Cache object
func New() *Cache {
	cache := &Cache{
		items: make(map[string]Item),
		mu:    &sync.Mutex{},
		close: make(chan bool),
	}

	timeStart := time.Now()
	//fmt.Println("timeStart:", timeStart)
	ticker := time.NewTicker(100 * time.Millisecond)

	go func() {
		for {
			select {
			case <-cache.close:
				return
			case <-ticker.C:
				timeNow := time.Now()
				for key, v := range cache.items {
					timeExpired := timeStart.Add(time.Duration(+v.ttl))
					if v.ttl > 0 && timeNow.Unix() > timeExpired.Unix() {
						fmt.Println("cache expired, cleaning key:", key)
						fmt.Printf("timeNow.Unix:%v timeExpired.Unit:%v\n", timeNow.Unix(), timeExpired.Unix())
						delete(cache.items, key)
					}
				}
			}
		}
	}()

	return cache
}

// Stop cache goroutine
func (c *Cache) Close() error {
	c.close <- true
	return nil
}
