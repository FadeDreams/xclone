package domain

import (
	"context"
	"errors"
	"fmt"
	"github.com/fadedreams/xclone"
	"golang.org/x/crypto/bcrypt"
)

var passwordCost = bcrypt.DefaultCost

type AuthService struct {
	AuthTokenService xclone.AuthTokenService
	UserRepo         xclone.UserRepo
}

func NewAuthService(ur xclone.UserRepo, ats xclone.AuthTokenService) *AuthService {
	return &AuthService{
		AuthTokenService: ats,
		UserRepo:         ur,
	}
}

func (as *AuthService) Register(ctx context.Context, input xclone.RegisterInput) (xclone.AuthResponse, error) {
	input.Sanitize()

	if err := input.Validate(); err != nil {
		return xclone.AuthResponse{}, err
	}

	// check if username is already taken
	if _, err := as.UserRepo.GetByUsername(ctx, input.Username); !errors.Is(err, xclone.ErrNotFound) {
		return xclone.AuthResponse{}, xclone.ErrUsernameTaken
	}

	// check if email is already taken
	if _, err := as.UserRepo.GetByEmail(ctx, input.Email); !errors.Is(err, xclone.ErrNotFound) {
		return xclone.AuthResponse{}, xclone.ErrEmailTaken
	}

	user := xclone.User{
		Email:    input.Email,
		Username: input.Username,
	}

	// hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), passwordCost)
	if err != nil {
		return xclone.AuthResponse{}, fmt.Errorf("error hashing password: %v", err)
	}

	user.Password = string(hashPassword)

	// create the user
	user, err = as.UserRepo.Create(ctx, user)
	if err != nil {
		return xclone.AuthResponse{}, fmt.Errorf("error creating user: %v", err)
	}

	accessToken, err := as.AuthTokenService.CreateAccessToken(ctx, user)
	if err != nil {
		return xclone.AuthResponse{}, xclone.ErrGenAccessToken
	}

	// return accessToken and user
	return xclone.AuthResponse{
		AccessToken: accessToken,
		User:        user,
	}, nil
}
