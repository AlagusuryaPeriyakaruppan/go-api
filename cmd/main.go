package main

import (
	"context"
	"fmt"
	"log"

	db "github.com/AlagusuryaPeriyakaruppan/go-api/internal/database"
)

func main() {
	fmt.Println("Surya loves everyone!")
	ctx := context.Background()
	db.InitDB()
	if err := db.EntClient.Schema.Create(ctx); err != nil {
		log.Println("Failed to create schema")
	}
	user, err := db.EntClient.User.Create().SetID(29).SetUsername("PP").Save(ctx)
	if err != nil {
		log.Println("Error creating User")
	}
	fmt.Println(user, "Creted User")
}
