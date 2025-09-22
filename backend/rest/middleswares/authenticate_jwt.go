package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// AuthenticateJWT validates JWT from Authorization header
func (m *Middleswares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Expected format: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Unauthorized: invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// Get secret key from config
		secretKey := []byte(m.cnf.JwtSecretKey)

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Make sure token uses HMAC SHA256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenUnverifiable
			}
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		// Optionally, you can extract claims and attach to context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Add claims to request context if needed
			// e.g., r = r.WithContext(context.WithValue(r.Context(), "userClaims", claims))
			_ = claims
		}

		// Token is valid, proceed to next handler
		next.ServeHTTP(w, r)
	})
}
