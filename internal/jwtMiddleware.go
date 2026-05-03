package main

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (rolePerm RolePermMap) RequireJWTPermission(hndlr http.HandlerFunc, reqPerm Permission) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Extract Bearer token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		tokenString, found := strings.CutPrefix(authHeader, "Bearer ")
		if !found || tokenString == "" {
			http.Error(w, "Bearer token required", http.StatusUnauthorized)
			return
		}

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			},
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		)

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		for _, role := range claims.Roles {
			if rolePerm[role][reqPerm] {
				hndlr(w, r)
				return
			}
		}

		http.Error(w, "Forbidden", http.StatusForbidden)

	}
}
