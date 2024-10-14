package middlewares

import (
	"fmt"
	"net/http"
)

func (s *Middlewares) AllowOnlyBDFL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(int32)
		isBDFL, err := s.Services.Db.IsBDFL(s.Services.Ctx, userID)
		if err != nil {
			s.Services.HttpError(w, fmt.Sprintf("Unable to check for bdfl: %v", err), http.StatusInternalServerError)
			return
		}
		if !isBDFL.Bool {
			s.Services.HttpError(w, "This route is restricted to BDFL only", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
