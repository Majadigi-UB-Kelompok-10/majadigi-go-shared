package init_helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Majadigi-UB-Kelompok-10/majadigi-go-shared/shared/cache"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeGracefulShutdownListener(app *fiber.App, pool *pgxpool.Pool) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	fmt.Println("\nShutdown signal received, starting graceful shutdown...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	fmt.Println("Shutting down HTTP server...")
	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v\n", err)
	}
	fmt.Println("HTTP server shutdown complete")

	fmt.Println("Closing database connections...")
	pool.Close()
	fmt.Println("Database connections closed")

	if redisCache, ok := cache.GlobalCache.(*cache.RedisCache); ok {
		fmt.Println("Closing Redis connection...")
		redisCache.Close()
		fmt.Println("Redis connection closed")
	}

	fmt.Println("Graceful shutdown complete")
}
