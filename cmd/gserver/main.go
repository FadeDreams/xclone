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
	"github.com/fadedreams/xclone"
	"github.com/fadedreams/xclone/domain"
	"github.com/fadedreams/xclone/graph"
	"github.com/fadedreams/xclone/jwt"
)

type EmptyAuthTokenService struct{}

func (s *EmptyAuthTokenService) CreateAccessToken(ctx context.Context, user xclone.User) (string, error) {
	// Implement the logic to create an access token here.
	// This can be a placeholder or an empty implementation depending on your needs.
	return "sample_access_token", nil
}

func (s *EmptyAuthTokenService) CreateRefreshToken(ctx context.Context, user xclone.User, tokenID string) (string, error) {
	// Implement the logic to create a refresh token here.
	// This can be a placeholder or an empty implementation depending on your needs.
	return "sample_refresh_token", nil
}

func (s *EmptyAuthTokenService) ParseToken(ctx context.Context, payload string) (xclone.AuthToken, error) {
	// Implement the logic to parse a token here.
	// This can be a placeholder or an empty implementation depending on your needs.
	return xclone.AuthToken{}, nil
}

func (s *EmptyAuthTokenService) ParseTokenFromRequest(ctx context.Context, r *http.Request) (xclone.AuthToken, error) {
	// Implement the logic to parse a token from a request here.
	// This can be a placeholder or an empty implementation depending on your needs.
	return xclone.AuthToken{}, nil
}
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

	userRepo := postgres.NewUserRepo(db)
	tweetRepo := postgres.NewTweetRepo(db)

	// authTokenService := &EmptyAuthTokenService{}
	authTokenService := jwt.NewTokenService(conf)
	authService := domain.NewAuthService(userRepo, authTokenService)
	tweetService := domain.NewTweetService(tweetRepo)
	userService := domain.NewUserService(userRepo)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Timeout(time.Second * 60))

	router.Use(authMiddleware(authTokenService))
	router.Handle("/", playground.Handler("X clone", "/query"))
	router.Handle("/query", handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					AuthService:  authService,
					TweetService: tweetService,
					UserService:  userService,
				},
			},
		),
	))

	log.Fatal(http.ListenAndServe(":8080", router))
}
