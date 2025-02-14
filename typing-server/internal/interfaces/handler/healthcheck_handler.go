package handler

import (
	"net/http"
)

// HealthCheckHandler はヘルスチェック用のハンドラ
type HealthCheckHandler struct{}

// NewHealthCheckHandler は HealthCheckHandler のインスタンスを生成する
func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

// LivenessProbe はアプリケーションが「生存」しているかを確認する
func (h *HealthCheckHandler) LivenessProbe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
