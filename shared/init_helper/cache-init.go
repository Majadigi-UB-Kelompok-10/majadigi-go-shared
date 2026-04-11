package init_helper

import (
	"fmt"
	"log"

	"github.com/Majadigi-UB-Kelompok-10/majadigi-go-shared/shared/cache"
)

func InitializeRedisCache(redis_url string) {
	redisURL := redis_url

	if redisURL != "" {
		redisCache, err := cache.NewRedisCache(redisURL)
		if err != nil {
			log.Fatalf("Failed to initialize Redis cache: %v\n", err)
		}
		cache.GlobalCache = redisCache
		fmt.Println("Redis cache initialized")
	} else {
		fmt.Println("Using SimpleCache (in-memory)")
	}
}
