package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiMiddleWare(t *testing.T) {
	validKeys := map[string]string{
		"testkey123": "valid",
	}

	tests := []struct {
		name       string
		authHeader string
		wantStatus int
	}{
		{"Valid key", "testkey123", http.StatusOK},
		{"Invalid key", "nope", http.StatusUnauthorized},
		{"No key", "", http.StatusUnauthorized},
	}

	mw := ApiKeyMiddleware(validKeys)

	mockHandler := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			rec := httptest.NewRecorder()

			mockHandler.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("got status %d, want status %d", rec.Code, tt.wantStatus)
			}

		})
	}

}
