package web

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getAuthCookie(r *http.Request) (string, error) {
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "Authorization" {
			return cookie.Value, nil
		}
	}

	return "", errors.New("Cookies not found")
}

type MyClaims struct {
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
	jwt.RegisteredClaims
}

func CreateJWT(userID string, userEmail string) (string, error) {
	jwtKey := []byte("some-key")
	expirationTime := time.Now().Add(24 * 60 * time.Hour)

	claims := &MyClaims{
		UserID:    userID,
		UserEmail: userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*MyClaims, error) {
	jwtKey := []byte("some-key")
	claims := &MyClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func createCookie(exp time.Time, key string, value string) http.Cookie {
	return http.Cookie{
		Name:     key,
		Value:    value,
		Expires:  exp,
		HttpOnly: false,
		Path:     "/",
	}

}
