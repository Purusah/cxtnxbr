package globals

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"
)

type KVStorage interface {
	Get(ctx context.Context, key string) (int, error)
	Incr(ctx context.Context, key string) error
}

type CacheMemory struct {
	m map[string]int
	sync.Mutex
}

func (c *CacheMemory) Get(_ context.Context, key string) (int, error) {
	val, ok := c.m[key]
	if !ok {
		return val, KeyNotFound{key}
	}
	return val, nil
}

func (c *CacheMemory) Incr(_ context.Context, key string) error {
	c.Lock()
	defer c.Unlock()
	val, ok := c.m[key]
	if !ok {
		c.m[key] = 1
		return nil
	}
	c.m[key] = val + 1
	return nil
}

type CacheRedis struct {
	m *redis.Client
}

func (c *CacheRedis) Get(ctx context.Context, key string) (int, error) {
	valNorm, err := c.m.Get(ctx, key).Result()
	if err != nil {
		log.Print("redis get ", err.Error())
		return 0, KeyNotFound{key}
	}
	val, err := strconv.Atoi(valNorm)
	if err != nil {
		return val, BrokenKey{key}
	}
	return val, nil
}

func (c *CacheRedis) Incr(ctx context.Context, key string) error {
	_, err := c.m.Incr(ctx, key).Result()
	return err
}
