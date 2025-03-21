package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/su-its/typing/typing-server/internal/domain/usecase"
)

// UserHandler はユーザー関連の HTTP ハンドラ
type UserHandler struct {
	userUseCase usecase.IUserUseCase
	log         *slog.Logger
}

// NewUserHandler は UserHandler のインスタンスを生成する
func NewUserHandler(userUseCase usecase.IUserUseCase, log *slog.Logger) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
		log:         log,
	}
}

// GetUserByStudentNumber は学籍番号をクエリパラメータとして受け取り、該当するユーザー情報を取得する
// ユーザーが見つからない場合は404 Not Foundを返す
// エラーが発生した場合は500 Internal Server Errorを返す
// クエリパラメータが指定されていない場合は400 Bad Requestを返す
func (h *UserHandler) GetUserByStudentNumber(w http.ResponseWriter, r *http.Request) {
	studentNumber := r.URL.Query().Get("student_number")

	if studentNumber == "" {
		http.Error(w, ErrMsgStudentNumberRequired, http.StatusBadRequest)
		return
	}

	user, err := h.userUseCase.GetUserByStudentNumber(r.Context(), studentNumber)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrUserNotFound):
			http.Error(w, ErrUserNotFound, http.StatusNotFound)
		default:
			h.log.Error("GetUserByStudentNumber failed", "error", err)
			http.Error(w, ErrInternalServer, http.StatusInternalServerError)
		}
		return
	}

	if user == nil {
		http.Error(w, ErrUserNotFound, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		h.log.Error("Failed to encode JSON response", "error", err)
		http.Error(w, ErrFailedToEncodeResponse, http.StatusInternalServerError)
	}
}
