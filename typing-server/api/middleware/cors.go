package middleware

import (
	"log/slog"
	"net/http"

	"github.com/rs/cors"
)

func CORS(log *slog.Logger) func(http.Handler) http.Handler {
	allowedOrigins := []string{"*"}
	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders := []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}
	exposedHeaders := []string{"Link"}
	maxAge := 300

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: allowedMethods,
		AllowedHeaders: allowedHeaders,
		ExposedHeaders: exposedHeaders,
		MaxAge:         maxAge,
	})

	log.Info("CORS middleware configured",
		"allowedOrigins", allowedOrigins,
		"allowedMethods", allowedMethods,
		"maxAge", maxAge)

	return corsMiddleware.Handler
}
