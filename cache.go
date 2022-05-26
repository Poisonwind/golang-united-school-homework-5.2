package cache

import "time"

type Cache struct {
	CacheMap map[string]string
	TimerMap map[string]time.Time
}

func NewCache() Cache {

	cacheMap := make(map[string]string)
	timerMap := make(map[string]time.Time) 

	return Cache{cacheMap, timerMap}
}

func (c *Cache) Get(key string) (string, bool) {

	if c.TimerMap[key].Before(time.Now()) {
		delete(c.TimerMap, key)
		delete(c.CacheMap, key)
	}

	res, ok := c.CacheMap[key]

	return res, ok
}

func (c *Cache) Put(key, value string) {

	c.CacheMap[key] = value

	_, ok := c.TimerMap[key]

	if ok {
		delete(c.TimerMap, key)
	}

}

func (c *Cache) Keys() []string {

	result := []string{}

	c.ClearCache()

	for key := range (c.CacheMap) {
		result = append(result, key)
	}

	return result

}

func (c *Cache) PutTill(key, value string, deadline time.Time) {

	c.CacheMap[key] = value
	c.TimerMap[key] = deadline

}

func (c *Cache) ClearCache() {

	for key, val := range(c.TimerMap) {
		if val.Before(time.Now()) {
			delete(c.TimerMap, key)
			delete(c.CacheMap, key)
		}
	}

}
