package domain

import (
	"context"

	"github.com/fadedreams/xclone"
	"github.com/fadedreams/xclone/uuid"
)

type UserService struct {
	UserRepo xclone.UserRepo
}

func NewUserService(ur xclone.UserRepo) *UserService {
	return &UserService{
		UserRepo: ur,
	}
}

func (u *UserService) GetByID(ctx context.Context, id string) (xclone.User, error) {
	if !uuid.Validate(id) {
		return xclone.User{}, xclone.ErrInvalidUUID
	}

	return u.UserRepo.GetByID(ctx, id)
}
