package graph

import (
	"context"
	"errors"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fadedreams/xclone"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// type Resolver struct{}

type tweetResolver struct {
	*Resolver
}

func (r *Resolver) Tweet() TweetResolver {
	return &tweetResolver{r}
}

type Resolver struct {
	AuthService  xclone.AuthService
	TweetService xclone.TweetService
	UserService  xclone.UserService
}

func buildBadRequestError(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code": http.StatusBadRequest,
		},
	}
}

func buildUnauthenticatedError(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code": http.StatusUnauthorized,
		},
	}
}

func buildForbiddenError(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code": http.StatusForbidden,
		},
	}
}

func buildNotFoundError(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code": http.StatusForbidden,
		},
	}
}

func buildError(ctx context.Context, err error) error {
	switch {
	case errors.Is(err, xclone.ErrForbidden):
		return buildForbiddenError(ctx, err)
	case errors.Is(err, xclone.ErrUnauthenticated):
		return buildUnauthenticatedError(ctx, err)
	case errors.Is(err, xclone.ErrValidation):
		return buildBadRequestError(ctx, err)
	case errors.Is(err, xclone.ErrNotFound):
		return buildNotFoundError(ctx, err)
	default:
		return err
	}
}
