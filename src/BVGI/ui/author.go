package ui

import (
	"net/http"
)

// author is a HTTP middleware for authorization
func author(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/login" {
				if getUsername(r) == "" {
					http.Redirect(w, r, "/login", http.StatusFound)
					return
				}	
			}
			next.ServeHTTP(w, r)
	})
}