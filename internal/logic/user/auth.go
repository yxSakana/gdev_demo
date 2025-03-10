package user

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/yxSakana/gdev_demo/internal/consts"
)

type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

// var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var jwtSecret = []byte("test")

func GenerateToken(userId uint64) (tokenStr string, err error) {
	claims := Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(consts.JwtExpire)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return
}

func ParseToken(tokenStr string) (claims *Claims, err error) {
	claims = &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return
}
