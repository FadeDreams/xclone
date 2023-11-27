package domain

import (
	"context"
	"errors"
	"testing"

	"github.com/fadedreams/xclone"
	"github.com/fadedreams/xclone/faker"
	"github.com/fadedreams/xclone/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAuthService_Register(t *testing.T) {
	validInput := xclone.RegisterInput{
		Username:        "xxx",
		Email:           "xxx@example.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	t.Run("can register", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		userRepo.On("GetByUsername", mock.Anything, mock.Anything).
			Return(xclone.User{}, xclone.ErrNotFound)

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).
			Return(xclone.User{}, xclone.ErrNotFound)

		userRepo.On("Create", mock.Anything, mock.Anything).
			Return(xclone.User{
				ID:       "123",
				Username: validInput.Username,
				Email:    validInput.Email,
			}, nil)

		authTokenService := &mocks.AuthTokenService{}

		authTokenService.On("CreateAccessToken", mock.Anything, mock.Anything).
			Return("a token", nil)

		service := NewAuthService(userRepo, authTokenService)

		res, err := service.Register(ctx, validInput)
		require.NoError(t, err)

		require.NotEmpty(t, res.AccessToken)
		require.NotEmpty(t, res.User.ID)
		require.NotEmpty(t, res.User.Email)
		require.NotEmpty(t, res.User.Username)

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})

	t.Run("username taken", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		userRepo.On("GetByUsername", mock.Anything, mock.Anything).
			Return(xclone.User{}, nil)

		authTokenService := &mocks.AuthTokenService{}

		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Register(ctx, validInput)
		require.ErrorIs(t, err, xclone.ErrUsernameTaken)

		userRepo.AssertNotCalled(t, "Create")

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})

	t.Run("email taken", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		userRepo.On("GetByUsername", mock.Anything, mock.Anything).
			Return(xclone.User{}, xclone.ErrNotFound)

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).
			Return(xclone.User{}, nil)

		authTokenService := &mocks.AuthTokenService{}

		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Register(ctx, validInput)
		require.ErrorIs(t, err, xclone.ErrEmailTaken)

		userRepo.AssertNotCalled(t, "Create")

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})

	t.Run("create error", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		userRepo.On("GetByUsername", mock.Anything, mock.Anything).
			Return(xclone.User{}, xclone.ErrNotFound)

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).
			Return(xclone.User{}, xclone.ErrNotFound)

		userRepo.On("Create", mock.Anything, mock.Anything).
			Return(xclone.User{}, errors.New("something"))

		authTokenService := &mocks.AuthTokenService{}

		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Register(ctx, validInput)
		require.Error(t, err)

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})

	t.Run("invalid input", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		authTokenService := &mocks.AuthTokenService{}

		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Register(ctx, xclone.RegisterInput{})
		require.ErrorIs(t, err, xclone.ErrValidation)

		userRepo.AssertNotCalled(t, "GetByUsername")
		userRepo.AssertNotCalled(t, "GetByEmail")
		userRepo.AssertNotCalled(t, "Create")

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})

	t.Run("can't generate access token", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		userRepo.On("GetByUsername", mock.Anything, mock.Anything).
			Return(xclone.User{}, xclone.ErrNotFound)

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).
			Return(xclone.User{}, xclone.ErrNotFound)

		userRepo.On("Create", mock.Anything, mock.Anything).
			Return(xclone.User{
				ID:       "123",
				Username: validInput.Username,
				Email:    validInput.Email,
			}, nil)

		authTokenService := &mocks.AuthTokenService{}

		authTokenService.On("CreateAccessToken", mock.Anything, mock.Anything).
			Return("", errors.New("error"))

		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Register(ctx, validInput)
		require.ErrorIs(t, err, xclone.ErrGenAccessToken)

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})
}

func TestAuthService_Login(t *testing.T) {
	validInput := xclone.LoginInput{
		Email:    "bob@gmail.com",
		Password: "password",
	}

	t.Run("can login", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).
			Return(xclone.User{
				Email:    validInput.Email,
				Password: faker.Password,
			}, nil)

		authTokenService := &mocks.AuthTokenService{}

		authTokenService.On("CreateAccessToken", mock.Anything, mock.Anything).
			Return("a token", nil)

		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Login(ctx, validInput)
		require.NoError(t, err)

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).
			Return(xclone.User{
				Email:    validInput.Email,
				Password: faker.Password,
			}, nil)

		authTokenService := &mocks.AuthTokenService{}

		service := NewAuthService(userRepo, authTokenService)

		input := xclone.LoginInput{
			Email:    validInput.Email,
			Password: "somethingelse",
		}

		_, err := service.Login(ctx, input)
		require.ErrorIs(t, err, xclone.ErrBadCredentials)

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})

	t.Run("email not found", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).
			Return(xclone.User{}, xclone.ErrNotFound)

		authTokenService := &mocks.AuthTokenService{}

		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Login(ctx, validInput)
		require.ErrorIs(t, err, xclone.ErrBadCredentials)

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})

	t.Run("get user by email error", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).
			Return(xclone.User{}, errors.New("something"))

		authTokenService := &mocks.AuthTokenService{}

		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Login(ctx, validInput)
		require.Error(t, err)

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})

	t.Run("invalid input", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		authTokenService := &mocks.AuthTokenService{}

		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Login(ctx, xclone.LoginInput{
			Email:    "bob",
			Password: "",
		})
		require.ErrorIs(t, err, xclone.ErrValidation)

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})

	t.Run("can't generate access token", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).
			Return(xclone.User{
				Email:    validInput.Email,
				Password: faker.Password,
			}, nil)

		authTokenService := &mocks.AuthTokenService{}

		authTokenService.On("CreateAccessToken", mock.Anything, mock.Anything).
			Return("", errors.New("error"))

		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Login(ctx, validInput)
		require.ErrorIs(t, err, xclone.ErrGenAccessToken)

		userRepo.AssertExpectations(t)
		authTokenService.AssertExpectations(t)
	})
}
