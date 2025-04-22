package router

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
)

func MwCsrf(next http.Handler) http.Handler {
	csrfKey := []byte(os.Getenv("CSRF_KEY"))
	inDevMode := os.Getenv("ENVIRONMENT") == "development"

	csrfOptions := []csrf.Option{
		csrf.SameSite(csrf.SameSiteStrictMode),
		csrf.HttpOnly(true),
		csrf.Path("/"),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slog.Error(
				fmt.Sprintf(
					"CSRF Error: %v, Origin: %v, Referer: %v, Host: %v",
					csrf.FailureReason(r),
					r.Header.Get("Origin"),
					r.Header.Get("Referer"),
					r.Host),
			)
			http.Error(w, "Forbidden - origin invalid", http.StatusForbidden)
		})),
	}

	if inDevMode {
		csrfOptions = append(csrfOptions, csrf.TrustedOrigins([]string{"localhost:8888"}))
	}

	csrfMw := csrf.Protect(csrfKey, csrfOptions...)

	return csrfMw(next)
}
