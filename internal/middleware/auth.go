package middleware

import (
	"net/http"
	"os"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "Missing API key", http.StatusUnauthorized)
			return
		}

		expectedAPIKey := os.Getenv("API_KEY")
		if expectedAPIKey == "" {
			expectedAPIKey = "default-dev-key" // For development only
		}

		if apiKey != expectedAPIKey {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
