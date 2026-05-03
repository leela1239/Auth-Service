package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_ultra_secret_key") // Keep this safe in env/config outside a POC.

func registerRoutes(mux *http.ServeMux) {
	// Map/header based RBAC POC.
	mux.HandleFunc("/view", RequirePermission(handleAuth, Permission("read")))
	mux.HandleFunc("/edit", RequirePermission(handleAuth, Permission("write")))
	mux.HandleFunc("/delete", RequirePermission(handleAuth, Permission("delete")))

	// JWT based RBAC POC.
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/jwtView", RolePerm.RequireJWTPermission(handleAuth, Permission("read")))
	mux.HandleFunc("/jwtEdit", RolePerm.RequireJWTPermission(handleAuth, Permission("write")))
	mux.HandleFunc("/jwtDelete", RolePerm.RequireJWTPermission(handleAuth, Permission("delete")))
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User has permission to access this resource."))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")

	roles, exists := UserRole[username]
	if !exists {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	claims := Claims{
		UserName: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))
}
