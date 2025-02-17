package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/su-its/typing/typing-server/internal/domain/usecase"
)

// UserHandler はユーザー関連の HTTP ハンドラ
type UserHandler struct {
	userUseCase usecase.IUserUseCase
}

// NewUserHandler は UserHandler のインスタンスを生成する
func NewUserHandler(userUseCase usecase.IUserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// GetUserByStudentNumber は学籍番号をクエリパラメータとして受け取り、該当するユーザー情報を取得する
// ユーザーが見つからない場合は404 Not Foundを返す
// エラーが発生した場合は500 Internal Server Errorを返す
// クエリパラメータが指定されていない場合は400 Bad Requestを返す
func (h *UserHandler) GetUserByStudentNumber(w http.ResponseWriter, r *http.Request) {
	studentNumber := r.URL.Query().Get("student_number")

	if studentNumber == "" {
		http.Error(w, "student_numberが指定されていません", http.StatusBadRequest)
		return
	}

	user, err := h.userUseCase.GetUserByStudentNumber(r.Context(), studentNumber)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrUserNotFound):
			http.Error(w, "ユーザーが見つかりません", http.StatusNotFound)
		default:
			http.Error(w, "内部サーバーエラーが発生しました", http.StatusInternalServerError)
		}
		return
	}

	if user == nil {
		http.Error(w, "ユーザーが見つかりません", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "レスポンスのエンコードに失敗しました", http.StatusInternalServerError)
	}
}
