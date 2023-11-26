package graph

import (
	// "context"
	// "errors"
	// "net/http"

	// "github.com/99designs/gqlgen/graphql"
	"github.com/fadedreams/xclone"
	// "github.com/vektah/gqlparser/v2/gqlerror"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// type Resolver struct{}

type Resolver struct {
	AuthService xclone.AuthService
}
