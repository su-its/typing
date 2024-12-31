package handler

import (
	"net/http"
)

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("API is running")); err != nil {
		h.log.Error("failed to write response",
			"error", err,
			"path", r.URL.Path)
	}
}
