package handler

import (
	"net/http"

	"github.com/stwile/go_todo_app/auth"
)

func AuthMiddleware(j *auth.JWTer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req, err := j.FillContext(r)
			if err != nil {
				RespondJson(
					r.Context(),
					w,
					ErrResponse{
						Message: "not find auth info",
						Details: []string{err.Error()},
					},
					http.StatusUnauthorized,
				)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if !auth.IsAdmin(ctx) {
			RespondJson(
				ctx,
				w,
				ErrResponse{Message: "not admin"},
				http.StatusUnauthorized,
			)
			return
		}
		next.ServeHTTP(w, r)
	})
}
