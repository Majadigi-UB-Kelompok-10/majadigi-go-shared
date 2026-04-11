package init_helper

import (
	"fmt"
	"log"
	"os"

	"github.com/Majadigi-UB-Kelompok-10/majadigi-go-shared/shared/cache"
)

func InitializeCache() {
	redisURL := os.Getenv("REDIS_URL")

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
