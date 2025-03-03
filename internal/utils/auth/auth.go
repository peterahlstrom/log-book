package auth

import "net/http"

func ApiKeyMiddleware(validKeys map[string]string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return ApiKeyHandler(next, validKeys)
	}
}

func ApiKeyHandler(next http.Handler, validKeys map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientKey := r.Header.Get("Authorization")
		if !ValidateApiKey(clientKey, validKeys) {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func ValidateApiKey(key string, validKeys map[string]string) bool {
	return validKeys[key] != ""
}

