package jwtx

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	UserID string `json:"user_id,omitempty"`
	Email  string `json:"email,omitempty"`
}

type Claims struct {
	Payload
	jwt.RegisteredClaims
}

func Sign(payload Payload, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Payload: Payload{
			UserID: payload.UserID,
			Email:  payload.Email,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	return token.SignedString([]byte("fsdhfkdsjlkfjdslfjdslkfjdlsfkdjslj"))
}
