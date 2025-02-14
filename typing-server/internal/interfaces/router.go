package interfaces

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/su-its/typing/typing-server/config"
	"github.com/su-its/typing/typing-server/internal/interfaces/handler"
	"github.com/su-its/typing/typing-server/pkg/middleware"
)

func NewRouter(healthHandler *handler.HealthCheckHandler, userHandler *handler.UserHandler, scoreHandler *handler.ScoreHandler, config *config.Config) http.Handler {
	r := chi.NewRouter()

	// ミドルウェアの設定
	r.Use(middleware.Trace)

	// CORSの設定
	corsConfig := middleware.DefaultCORSConfig()
	corsConfig.AllowedOrigins = getAllowedOrigins(config.Environment)
	r.Use(middleware.CORSMiddleware(corsConfig))

	// ルートの設定
	routes := []struct {
		method  string
		path    string
		handler http.HandlerFunc
	}{
		{"GET", "/health", healthHandler.LivenessProbe},
		{"GET", "/users", userHandler.GetUserByStudentNumber},
		{"GET", "/scores/ranking", scoreHandler.GetScoresRanking},
		{"POST", "/scores", scoreHandler.RegisterScore},
	}

	for _, route := range routes {
		switch route.method {
		case "GET":
			r.Get(route.path, route.handler)
		case "POST":
			r.Post(route.path, route.handler)
		}
	}

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
