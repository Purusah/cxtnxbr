package globals

import (
	"context"
	"os"
	"testing"

	"github.com/purusah/cxtnxbr/pkg/config"
)

func TestCacheMemoryGetExisted(t *testing.T) {
	key := "key"
	value := 137
	conf, _ := config.GetConfig()
	conf.Redis.Url = ""
	conf.Redis.Db = 0
	ci := NewCache(conf).C
	c, ok := ci.(*CacheMemory)
	if !ok {
		t.Fatalf("expected CacheMemory")
	}
	c.m[key] = value
	receivedValue, err := c.Get(context.Background(), key)
	if err != nil {
		t.Fatalf("something wrong %s", err.Error())
	}
	if receivedValue != value {
		t.Fatalf("values must be equal")
	}
}

func TestCacheMemoryGetNotExisted(t *testing.T) {
	key := "key"
	expectedValue := 0
	conf, _ := config.GetConfig()
	conf.Redis.Url = ""
	conf.Redis.Db = 0
	ci := NewCache(conf).C
	c, ok := ci.(*CacheMemory)
	if !ok {
		t.Fatalf("expected CacheMemory")
	}
	receivedValue, err := c.Get(context.Background(), key)
	switch err.(type) {
	case KeyNotFound:
	default:
		t.Fatalf("cache error %s", err.Error())
	}
	if receivedValue != expectedValue {
		t.Fatalf("expected value should be zeroed")
	}
}

func TestCacheRedisGetExisted(t *testing.T) {
	redisUrl := os.Getenv("COUNTER_REDIS_URL")
	if redisUrl == "" {
		t.Skipf("redis url not set")
	}
	key := "key"
	value := 137
	conf, _ := config.GetConfig()
	ci := NewCache(conf).C
	c, ok := ci.(*CacheRedis)
	if !ok {
		t.Fatalf("expected CacheRedis")
	}
	_ = c.m.Set(context.Background(), key, value, 0)
	receivedValue, err := c.Get(context.Background(), key)
	if err != nil {
		t.Fatalf("cache error %s", err.Error())
	}
	if receivedValue != value {
		t.Fatalf("values must be equal")
	}
	c.m.Del(context.Background(), key)
}

func TestCacheRedisGetNotExisted(t *testing.T) {
	redisUrl := os.Getenv("COUNTER_REDIS_URL")
	if redisUrl == "" {
		t.Skipf("redis url not set")
	}
	key := "key"
	value := 0
	conf, _ := config.GetConfig()
	ci := NewCache(conf).C
	c, ok := ci.(*CacheRedis)
	if !ok {
		t.Fatalf("expected CacheRedis")
	}
	receivedValue, err := c.Get(context.Background(), key)
	switch err.(type) {
	case KeyNotFound:
	default:
		t.Fatalf("cache error %s", err.Error())
	}
	if receivedValue != value {
		t.Fatalf("values must be equal")
	}
	c.m.Del(context.Background(), key)
}
