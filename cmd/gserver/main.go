package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fadedreams/xclone/config"
	"github.com/fadedreams/xclone/postgres"
)

func main() {
	ctx := context.Background()

	conf, err := config.New()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	db := postgres.New(ctx, conf)
	if db == nil {
		log.Fatal("Error creating database instance")
	}

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration success")
}
