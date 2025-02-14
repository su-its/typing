package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
	ExposedHeaders []string
	MaxAge         int
}

func DefaultCORSConfig() CORSConfig {
	return CORSConfig{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}
}

func CORSMiddleware(config CORSConfig) func(http.Handler) http.Handler {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: config.AllowedOrigins,
		AllowedMethods: config.AllowedMethods,
		AllowedHeaders: config.AllowedHeaders,
		ExposedHeaders: config.ExposedHeaders,
		MaxAge:         config.MaxAge,
	})
	return corsMiddleware.Handler
}
