package middlewares

import (
	"context"
	"net/http"
	"newproject/database"
	"newproject/helpers"
	"newproject/models"
	"strings"
)

//AuthMiddleware middleware for authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		tokenString := strings.SplitN(auth, " ", 2)[1]
		claims, ok, err := helpers.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !ok {
			http.Error(w, "You are unauthorized", http.StatusUnauthorized)
			return
		}
		db := database.DB
		u := models.User{}
		db.Where("Username = ?", claims.Username).First(&u)

		if (models.User{}) == u {
			http.Error(w, "You are unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, helpers.ContextKey("user"), u)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
