package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fadedreams/xclone/config"
	"github.com/fadedreams/xclone/postgres"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fadedreams/xclone/graph"
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

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Timeout(time.Second * 60))

	router.Handle("/", playground.Handler("X clone", "/query"))
	router.Handle("/query", handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	))

	log.Fatal(http.ListenAndServe(":8080", router))
}
