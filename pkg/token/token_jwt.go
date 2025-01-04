package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (t *token) JwtSign(userId int64, role, openId string, expiredDuration time.Duration) (string, error) {
	claims := claims{
		userId,
		role,
		openId,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredDuration)),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.Secret))
}

func (t *token) JwtParse(tokenString string) (*claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.Secret), nil
	})
	if tokenClaims != nil && tokenClaims.Valid {
		if claims, ok := tokenClaims.Claims.(*claims); ok {
			return claims, nil
		}
	}
	return nil, err
}
