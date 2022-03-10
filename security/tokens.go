package security

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidToken = errors.New("invalid jwt")

	jwtSecretKey = []byte(os.Getenv(""))
)

func NewToken(userId string) (string, error) {
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		Issuer:    userId,
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func parseJwtCallBack(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return jwtSecretKey, nil
}
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, parseJwtCallBack)
}

type TokenPayload struct {
	UserId    string
	Created   time.Time
	ExpiresAt time.Time
}

func NewTokenPayload(tokenString string) (*TokenPayload, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !token.Valid || !ok {
		return nil, ErrInvalidToken
	}

	id, _ := claims["iss"].(string)
	createdAt, _ := claims["iat"].(float64)
	expiresAt, _ := claims["exp"].(float64)
	return &TokenPayload{
		UserId:    id,
		Created:   time.Unix(int64(createdAt), 0),
		ExpiresAt: time.Unix(int64(expiresAt), 0),
	}, nil
}
