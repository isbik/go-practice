package middleware

import (
	"context"
	"main/internal/mocks"
	"net/http"
	"os"
)

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, password, ok := r.BasicAuth()

		if name != mocks.Admin.Name || password != mocks.Admin.Password || !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		h.ServeHTTP(w, r)
	})
}

func RequestIDMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := os.Getuid()
		print(id)
		ctx := context.WithValue(context.Background(), "ID", id)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
