package jwt

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fadedreams/xclone"
	"github.com/fadedreams/xclone/config"
	"github.com/lestrrat-go/jwx/jwa"
	jwtrat "github.com/lestrrat-go/jwx/jwt"
)

var signatureType = jwa.HS256

var now = time.Now

type TokenService struct {
	Conf *config.Config
}

func NewTokenService(conf *config.Config) *TokenService {
	return &TokenService{
		Conf: conf,
	}
}

func (ts *TokenService) ParseTokenFromRequest(ctx context.Context, r *http.Request) (xclone.AuthToken, error) {
	token, err := jwtrat.ParseRequest(
		r,
		jwtrat.WithValidate(true),
		jwtrat.WithIssuer(ts.Conf.JWT.Issuer),
		jwtrat.WithVerify(signatureType, []byte(ts.Conf.JWT.Secret)),
	)
	if err != nil {
		return xclone.AuthToken{}, xclone.ErrInvalidAccessToken
	}

	return buildToken(token), nil
}

func buildToken(token jwtrat.Token) xclone.AuthToken {
	return xclone.AuthToken{
		ID:  token.JwtID(),
		Sub: token.Subject(),
	}
}

func (ts *TokenService) ParseToken(ctx context.Context, payload string) (xclone.AuthToken, error) {
	token, err := jwtrat.Parse(
		[]byte(payload),
		jwtrat.WithValidate(true),
		jwtrat.WithIssuer(ts.Conf.JWT.Issuer),
		jwtrat.WithVerify(signatureType, []byte(ts.Conf.JWT.Secret)),
	)
	if err != nil {
		return xclone.AuthToken{}, xclone.ErrInvalidAccessToken
	}

	return buildToken(token), nil
}

func (ts *TokenService) CreateRefreshToken(ctx context.Context, user xclone.User, tokenID string) (string, error) {
	t := jwtrat.New()

	if err := setDefaultToken(t, user, xclone.RefreshTokenLifetime, ts.Conf); err != nil {
		return "", err
	}

	if err := t.Set(jwtrat.JwtIDKey, tokenID); err != nil {
		return "", fmt.Errorf("error set jwt id: %v", err)
	}

	token, err := jwtrat.Sign(t, signatureType, []byte(ts.Conf.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("error sign jwt: %v", err)
	}

	return string(token), nil
}

func (ts *TokenService) CreateAccessToken(ctx context.Context, user xclone.User) (string, error) {
	t := jwtrat.New()

	if err := setDefaultToken(t, user, xclone.AccessTokenLifetime, ts.Conf); err != nil {
		return "", err
	}

	token, err := jwtrat.Sign(t, signatureType, []byte(ts.Conf.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("error sign jwt: %v", err)
	}

	return string(token), nil
}

func setDefaultToken(t jwtrat.Token, user xclone.User, lifetime time.Duration, conf *config.Config) error {
	if err := t.Set(jwtrat.SubjectKey, user.ID); err != nil {
		return fmt.Errorf("error set jwt sub: %v", err)
	}

	if err := t.Set(jwtrat.IssuerKey, conf.JWT.Issuer); err != nil {
		return fmt.Errorf("error set jwt issuer key: %v", err)
	}

	if err := t.Set(jwtrat.IssuedAtKey, now().Unix()); err != nil {
		return fmt.Errorf("error set jwt issued at key: %v", err)
	}

	if err := t.Set(jwtrat.ExpirationKey, now().Add(lifetime).Unix()); err != nil {
		return fmt.Errorf("error set jwt expired at: %v", err)
	}

	return nil
}
