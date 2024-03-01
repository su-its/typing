package presenter

import (
	"net/http"

	"github.com/su-its/typing/typing-server/api/controller/system"
)

func RegisterRoutes() {
	http.HandleFunc("/health", system.HealthCheck)
}
