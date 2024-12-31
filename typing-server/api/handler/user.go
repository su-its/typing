package handler

import (
	"encoding/json"
	"net/http"

	"github.com/su-its/typing/typing-server/api/service"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 学籍番号の取得と検証
	studentNumber := r.URL.Query().Get("student_number")
	if studentNumber == "" {
		h.log.Error("student_number is missing")
		http.Error(w, "student_number is required", http.StatusBadRequest)
		return
	}

	// ユーザー情報の取得
	user, err := service.GetUserByStudentNumber(ctx, h.entClient, studentNumber)
	if err != nil {
		h.log.Error("failed to get user",
			"error", err,
			"student_number", studentNumber)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// レスポンスの返却
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		h.log.Error("failed to encode response",
			"error", err,
			"student_number", studentNumber)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.log.Debug("user retrieved successfully",
		"student_number", studentNumber)
}
