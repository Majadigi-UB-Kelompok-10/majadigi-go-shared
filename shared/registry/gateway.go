package registry

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AutoRegister(gatewayDBUrl string, slugName string, pageURL string) {
	if gatewayDBUrl == "" {
		log.Println("⚠️ AutoRegister dibatalkan: URL DB Gateway tidak ditemukan")
		return
	}

	dbGateway, err := pgxpool.New(context.Background(), gatewayDBUrl)
	if err != nil {
		log.Printf("❌ Gagal konek ke DB Gateway: %v\n", err)
		return
	}
	defer dbGateway.Close()

	query := `
		INSERT INTO endpoint_list (slug_name, page_url) 
		VALUES ($1, $2) 
		ON CONFLICT (slug_name) 
		DO UPDATE SET page_url = EXCLUDED.page_url;
	`

	_, err = dbGateway.Exec(context.Background(), query, slugName, pageURL)
	if err != nil {
		log.Printf("Gagal mendaftarkan '%s': %v\n", slugName, err)
	} else {
		fmt.Printf("Berhasil! Service '%s' mendaftarkan diri ke Gateway\n", slugName)
	}
}