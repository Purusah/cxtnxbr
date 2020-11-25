package config

import (
	"log"
	"os"
	"strconv"
)

type templatesConfig struct {
	CounterFile string
}

type redisConfig struct {
	Url string
	Db  int
}

type appConfig struct {
	Name string
	Url  string
}

type Config struct {
	App       appConfig
	Redis     redisConfig
	Templates templatesConfig
}

func GetConfig() (Config, error) {
	redisURL := os.Getenv("COUNTER_REDIS_URL")
	receivedRedisDb := os.Getenv("COUNTER_REDIS_DB")
	redisDB, err := strconv.Atoi(receivedRedisDb)
	if err != nil {
		log.Print("error COUNTER_REDIS_DB", receivedRedisDb)
		redisDB = 0
	}
	return Config{
		App: appConfig{
			Name: "cxtnxbr",
			Url:  "0.0.0.0:8080",
		},
		Redis: redisConfig{
			Url: redisURL,
			Db:  redisDB,
		},
		Templates: templatesConfig{CounterFile: "assets/counter.gohtml"},
	}, nil
}
