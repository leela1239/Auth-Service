package main

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserName string `json:"sub"`
	Roles    []Role `json:"roles"`
	jwt.RegisteredClaims
}
