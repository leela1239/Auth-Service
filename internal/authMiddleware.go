package main

import "net/http"

func RequirePermission(hndlr http.HandlerFunc, reqPerm Permission) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")

		if IsAuthorized(userID, reqPerm) {
			hndlr(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	}
}
