package xclone

import (
	"context"
	"errors"
	"time"
)

var (
	ErrUsernameTaken      = errors.New("username taken")
	ErrEmailTaken         = errors.New("email taken")
	ErrGenAccessToken     = errors.New("generate access token error")
	ErrNotFound           = errors.New("not found")
	ErrValidation         = errors.New("validation error")
	ErrBadCredentials     = errors.New("email/password wrong combination")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrNoUserIDInContext  = errors.New("no user id in context")
	ErrUnauthenticated    = errors.New("unauthenticated")
	ErrInvalidUUID        = errors.New("invalid uuid")
	ErrForbidden          = errors.New("forbidden")
)

type UserRepo interface {
	Create(ctx context.Context, user User) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id string) (User, error)
	GetByIds(ctx context.Context, ids []string) ([]User, error)
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserService interface {
	GetByID(ctx context.Context, id string) (User, error)
}
