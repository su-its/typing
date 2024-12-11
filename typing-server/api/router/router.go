package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/su-its/typing/typing-server/api/config"
	"github.com/su-its/typing/typing-server/api/handler"
)

func SetupRouter(config *config.Config) http.Handler {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: getAllowedOrigins(config.Environment),
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Preflightリクエストの結果をキャッシュする時間（秒）
	})
	fmt.Println(getAllowedOrigins(config.Environment))
	r.Use(cors.Handler)

	r.Get("/health", handler.HealthCheck)

	r.Get("/users", handler.GetUser)

	r.Get("/scores/ranking", handler.GetScoresRanking)
	r.Post("/scores", handler.PostScore)

	return r
}

func getAllowedOrigins(enviroment string) []string {
	if enviroment == "local" {
		return []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
		}
	}
	return []string{
		"http://ty.inf.in.shizuoka.ac.jp",
		"https://ty.inf.in.shizuoka.ac.jp",
	}
}
