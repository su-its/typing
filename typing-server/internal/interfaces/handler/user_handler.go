package handler

import (
	"encoding/json"
	"net/http"

	"github.com/su-its/typing/typing-server/internal/domain/usecase"
)

// UserHandler はユーザー関連の HTTP ハンドラ
type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

// NewUserHandler は UserHandler のインスタンスを生成する
func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

// GetUserByStudentNumber は学籍番号をクエリパラメータとして受け取り、該当するユーザー情報を取得する
func (h *UserHandler) GetUserByStudentNumber(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	studentNumber := r.URL.Query().Get("student_number")

	// student_number がない場合は 400 Bad Request を返す
	if studentNumber == "" {
		http.Error(w, "Missing student_number query parameter", http.StatusBadRequest)
		return
	}

	// ユースケースを呼び出してユーザーを取得
	user, err := h.userUseCase.GetUserByStudentNumber(r.Context(), studentNumber)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// JSON レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
