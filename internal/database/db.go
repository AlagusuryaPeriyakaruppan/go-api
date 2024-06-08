package db

import (
	"log"

	"github.com/AlagusuryaPeriyakaruppan/go-api/internal/ent"
	_ "github.com/lib/pq"
)

var EntClient *ent.Client

func initDB() {
	client, err := ent.Open("postgres", "host=localhost port=5433 user=surya dbname=postgres password=surya123 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// defer client.Close()

	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }
	EntClient = client
}
