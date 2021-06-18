package jwtx

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(claims jwt.MapClaims, key string, tmout time.Duration) (string, error) {
	claims["times"] = time.Now()
	if tmout > 0 {
		claims["timeout"] = time.Now().Add(tmout)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokens, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tokens, nil
}

func GetToken(s string, key string) jwt.MapClaims {
	if s == "" {
		return nil
	}
	tk, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err == nil {
		claim, ok := tk.Claims.(jwt.MapClaims)
		if ok {
			return claim
		}
	}
	return nil
}
