package main

import (
	"fmt"

	"context"

	db "github.com/AlagusuryaPeriyakaruppan/go-api/internal/database"
)

func main() {
	fmt.Println("Surya loves everyone!")
	ctx := context.Background()
	db.InitDB()
	db.EntClient.User.Create().SetID(29).SetUsername("PP").Save(ctx)
}
