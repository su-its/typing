package router

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/su-its/typing/typing-server/api/handler"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: getAllowedOrigins(),
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Preflightリクエストの結果をキャッシュする時間（秒）
	})
	r.Use(cors.Handler)

	r.Get("/health", handler.HealthCheck)

	r.Get("/users", handler.GetUser)

	r.Get("/scores/ranking", handler.GetScoresRanking)
	r.Post("/scores", handler.PostScore)

	return r
}

func getAllowedOrigins() []string {
	if os.Getenv("DEV_MODE") == "true" {
		return []string{"http://localhost:*", "http://ty.inf.in.shizuoka.ac.jp"}
	}
	return []string{"http://ty.inf.in.shizuoka.ac.jp"}
}
