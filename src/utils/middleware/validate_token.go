package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Define a custom type for context keys to avoid collisions
type contextKey string

const UserIDKey contextKey = "userID"

// AuthMiddleware checks if the request contains a valid JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization header missing", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix to extract the token string
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// If parsing failed or token is invalid, return 401 Unauthorized
		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return 
		}

		// Extract user_id claim from the token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if userID, ok := claims["user_id"].(float64); ok {
				// Store userID into request context for further usage
				ctx := context.WithValue(r.Context(), UserIDKey, int(userID))
				next.ServeHTTP(w, r.WithContext(ctx))
				return 
			}
		}

		http.Error(w, "invalid userID", http.StatusUnauthorized)
	})
}