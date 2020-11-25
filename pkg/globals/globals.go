package globals

import (
	"github.com/gorilla/websocket"
	"html/template"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/purusah/cxtnxbr/pkg/config"
)

type globalTemplates struct {
	Counter *template.Template
}

type Global struct {
	C         KVStorage
	T         globalTemplates
	WsUpgrade websocket.Upgrader
}

func NewCache(c config.Config) *Global {
	var cache KVStorage
	t, err := template.New("counter.gohtml").ParseFiles(c.Templates.CounterFile)
	if err != nil {
		log.Printf("template not set")
		t = nil
	}

	if c.Redis.Url != "" {
		cache = &CacheRedis{redis.NewClient(&redis.Options{
			Addr:     c.Redis.Url,
			Username: "",
			Password: "",
			DB:       c.Redis.Db,
		})}
	} else {
		cache = &CacheMemory{
			m: make(map[string]int),
		}
		log.Print("fall back to memory cache")
	}

	return &Global{
		C:         cache,
		T:         globalTemplates{Counter: t},
		WsUpgrade: websocket.Upgrader{},
	}
}
