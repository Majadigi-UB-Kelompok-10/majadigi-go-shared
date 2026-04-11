package init_helper

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Majadigi-UB-Kelompok-10/majadigi-go-shared/shared/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializePostgreDB(db_url string) *pgxpool.Pool {
	dbURL := db_url

	config, errConf := pgxpool.ParseConfig(dbURL)
	if errConf != nil {
		log.Fatalf("❌ Gagal parse config DB: %s\n", util.MaskDBSensitiveData(errConf.Error()))
	}

	config.MaxConns = 20
	config.MinConns = 5
	config.MaxConnLifetime = 1 * time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s\n", util.MaskDBSensitiveData(err.Error()))
	}
	fmt.Println("Database PostgreSQL connected")

	return pool
}
