package system

import (
	"net/http"
)

// HealthCheck はヘルスチェックのためのハンドラー関数です。
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is running"))
}
