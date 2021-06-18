package jwtx

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	claims := jwt.MapClaims{}
	m := map[string]string{}
	m["name"] = "123"
	claims["u"] = m
	token, err := CreateToken(claims, "u", time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)
	getToken := GetToken(token, "u")
	fmt.Println(getToken["u"])
}
