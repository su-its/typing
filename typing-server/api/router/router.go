package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/su-its/typing/typing-server/api/handler"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/health", handler.HealthCheck)

	r.Get("/users", handler.GetUsers)

	r.Get("/scores/ranking", handler.GetScoresRanking)
	r.Post("/scores", handler.PostScore)

	return r
}
