package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/su-its/typing/typing-server/api/handler"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// CORSの設定
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // 許可するオリジンを指定
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
	r.Get("/scores/{user-id}/current-rank", handler.GetMyScoreRanking)
	return r
}
