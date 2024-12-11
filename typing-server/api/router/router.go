package router

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/su-its/typing/typing-server/api/handler"
	"github.com/su-its/typing/typing-server/api/middleware"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
)

func SetupRouter(log *slog.Logger, entClient *ent.Client) http.Handler {
	r := chi.NewRouter()

	// ハンドラーの初期化
	h := handler.New(log, entClient)

	// ミドルウェアの設定
	r.Use(middleware.Trace)
	r.Use(middleware.CORS(log))

	// ルートの設定
	routes := []struct {
		method  string
		path    string
		handler http.HandlerFunc
	}{
		{"GET", "/health", h.HealthCheck},
		{"GET", "/users", h.GetUser},
		{"GET", "/scores/ranking", h.GetScoresRanking},
		{"POST", "/scores", h.PostScore},
	}

	for _, route := range routes {
		switch route.method {
		case "GET":
			r.Get(route.path, route.handler)
		case "POST":
			r.Post(route.path, route.handler)
		}
		log.Debug("route configured",
			"method", route.method,
			"path", route.path)
	}

	log.Info("routes configured")
	return r
}
