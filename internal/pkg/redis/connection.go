package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func SetupRedisClient() *redis.Client {
	host := os.Getenv("REDIS_HOSTNAME")
	port := os.Getenv("REDIS_PORT")

	address := fmt.Sprintf("%s:%s", host, port)

	pwd := os.Getenv("REDIS_PASSWORD")

	// only initializes if there's no previous connection
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: pwd,
			DB:       0, // Use default DB
		})
	}

	return rdb
}
