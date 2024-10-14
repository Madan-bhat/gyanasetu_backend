package middlewares

import (
	"net/http"
	"strings"

	"context"
)

func (s *Middlewares) RestictedAccess(allowBDFL bool, onlyForBDFL bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				s.Services.HttpError(w, "Authorization header not found", http.StatusUnauthorized)
				return
			}
			slices := strings.SplitN(authHeader, " ", 2)
			if len(slices) != 2 || strings.ToLower(slices[0]) != "bearer" {
				s.Services.HttpError(w, "Bearer token not found", http.StatusUnauthorized)
				return
			}
			token := slices[1]
			id, err := s.Services.GetIdFromToken(token)
			if s.Services.ISEOnError(w, err) {
				return
			}

			exists, err := s.Services.Db.UserExistsByID(s.Services.Ctx, id)
			if s.Services.ISEOnError(w, err) {
				return
			}
			if exists {
				if id == s.Services.BDFLId {

					if !allowBDFL {

						s.Services.HttpError(w, "BDFL Operations are restricted on this route", http.StatusForbidden)
						return
					}

					ctx := context.WithValue(r.Context(), "user_id", int32(id))
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {

					if allowBDFL && onlyForBDFL {

						s.Services.HttpError(w, "Access restricted for non-BDFL users", http.StatusForbidden)
						return
					}

					ctx := context.WithValue(r.Context(), "user_id", int32(id))
					next.ServeHTTP(w, r.WithContext(ctx))
				}
			} else {
				s.Services.RespondJson(w, "User not found", http.StatusUnauthorized)
				return
			}
		})
	}
}
