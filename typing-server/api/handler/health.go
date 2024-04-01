package handler

import (
	"log/slog"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("API is running"))
	if err != nil {
		slog.Error("failed to write response: %v", err)
	}
}
