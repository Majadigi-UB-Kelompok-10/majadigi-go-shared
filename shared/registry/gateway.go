package registry

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AutoRegisterFull(gatewayDBUrl, slugName, pageURL, title, iconURL, description string, categoryIDs []string) {
      if gatewayDBUrl == "" {                                                                                                                                                                                                         
          log.Println("AutoRegisterFull dibatalkan: URL DB Gateway tidak ditemukan")
          return                                                                                                                                                                                                                      
      }                                                                                                                                                                                                                               
                                                                                                                                                                                                                                      
      db, err := pgxpool.New(context.Background(), gatewayDBUrl)                                                                                                                                                                      
      if err != nil {
          log.Printf("Gagal konek ke DB Gateway: %v\n", err)
          return                                                                                                                                                                                                                      
      }
      defer db.Close()                                                                                                                                                                                                                
                  
      ctx := context.Background()

      _, err = db.Exec(ctx, `
          INSERT INTO endpoint_list (slug_name, page_url)                                                                                                                                                                             
          VALUES ($1, $2)                                                                                                                                                                                                             
          ON CONFLICT (slug_name) DO UPDATE SET page_url = EXCLUDED.page_url
      `, slugName, pageURL)                                                                                                                                                                                                           
      if err != nil {
          log.Printf("Gagal register endpoint '%s': %v\n", slugName, err)
          return                                                                                                                                                                                                                      
      }
                                                                                                                                                                                                                                      
      // 2. Upsert service ke service_list, ambil UUID-nya                                                                                                                                                                            
      var serviceID string
      err = db.QueryRow(ctx, `                                                                                                                                                                                                        
          INSERT INTO service_list (title, description, icon_url)
          VALUES ($1, $2, $3)                                                                                                                                                                                                         
          ON CONFLICT (title) DO UPDATE SET icon_url = EXCLUDED.icon_url, description = EXCLUDED.description
          RETURNING service_list_id                                                                                                                                                                                                   
      `, title, description, iconURL).Scan(&serviceID)
      if err != nil {                                                                                                                                                                                                                 
          log.Printf("Gagal register service '%s': %v\n", title, err)
          return                                                                                                                                                                                                                      
      }
                                                                                                                                                                                                                                      
      // 3. Assign ke kategori
      for _, catID := range categoryIDs {
          _, _ = db.Exec(ctx, `                                                                                                                                                                                                       
              INSERT INTO services_has_categories (service_list_id, category_list_id)
              VALUES ($1::uuid, $2::uuid)                                                                                                                                                                                             
              ON CONFLICT DO NOTHING                                                                                                                                                                                                  
          `, serviceID, catID)
      }                                                                                                                                                                                                                               
                  
      fmt.Printf("Service '%s' terdaftar ke gateway (id: %s, %d kategori)\n", title, serviceID, len(categoryIDs))                                                                                                                  
  }