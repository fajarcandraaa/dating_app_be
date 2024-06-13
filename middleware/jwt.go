package middleware

import (
	"errors"
	"net/http"

	auth "github.com/fajarcandraaa/dating_app_be/config/auth"
	"github.com/fajarcandraaa/dating_app_be/helpers"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			helpers.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	jwt := SetMiddlewareJSON(SetMiddlewareAuthentication(next))
	return jwt
}
