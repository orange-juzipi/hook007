package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Token interface {
		JwtSign(userId int64, role, openId string, expiredDuration time.Duration) (string, error)
		JwtParse(tokenString string) (*claims, error)
	}
	token struct {
		Secret string
	}

	claims struct {
		UserId int64
		Role   string
		OpenId string
		jwt.RegisteredClaims
	}
)

func New(secret string) Token {
	return &token{
		Secret: secret,
	}
}
