package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrorInvalidToken = "token is invalid"
	ErrorExpiredToken = "token is expired"
)


var key = []byte(os.Getenv("SECRET_KEY"))

// GenerateToken ...
func GenerateToken(Email string) (string, error) {
	expirationTime := time.Now().Add(1440 *time.Minute)

	// set claims i.e (payload)
	claims := jwt.MapClaims{
		"email": Email,
		"exp": expirationTime.Unix(),
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//sign token with secret key
	ts, err := token.SignedString(key)

	return ts, err
}

func ValidateToken(bearer string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(ErrorInvalidToken)
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			// Check if the token has expired
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return nil, fmt.Errorf("token has expired")
			}
		}
		return claims, nil
	} else {
		return nil, errors.New(ErrorInvalidToken)
	}
}

