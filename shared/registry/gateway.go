package registry

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AutoRegister(db *pgxpool.Pool, slugName string, pageURL string) {
	query := `
		INSERT INTO endpoint_list (slug_name, page_url) 
		VALUES ($1, $2) 
		ON CONFLICT (slug_name) 
		DO UPDATE SET page_url = EXCLUDED.page_url;
	`

	_, err := db.Exec(context.Background(), query, slugName, pageURL)
	if err != nil {
		log.Printf("Gagal mendaftarkan '%s': %v\n", slugName, err)
	} else {
		fmt.Printf("Berhasil! Gateway merouting '/%s/*' ke '%s'\n", slugName, pageURL)
	}
}