package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache struct {
	cache *cache.Cache
}

func NewCache() *Cache {
	return &Cache{
		cache: cache.New(20*time.Minute, 5*time.Minute),
	}
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	c.cache.Set(key, data, expiration)
	return nil
}

func (c *Cache) Get(ctx context.Context, key string, value interface{}) error {
	data, found := c.cache.Get(key)
	if !found {
		return errors.New("cache not found")
	}
	return json.Unmarshal(data.([]byte), value)
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	c.cache.Delete(key)
	return nil
}
