package middlewares

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	key []byte
}

func NewJWT(key []byte) JWT {
	return JWT{
		key: key,
	}
}

func (j JWT) Create(ttl time.Duration, content interface{}) (string, error) {
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["dat"] = content             // Our custom data.
	claims["exp"] = now.Add(ttl).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()          // The time at which the token was issued.
	claims["nbf"] = now.Unix()          // The time before which the token must be disregarded.

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.key)
	if err != nil {
		return "", fmt.Errorf("create: %w", err)
	}

	return token, nil
}

func (j JWT) Validate(token string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return j.key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return claims["dat"], nil
}
