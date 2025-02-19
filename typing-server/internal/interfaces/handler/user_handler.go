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

const (
	ErrMsgStudentNumberRequired = "student_numberが指定されていません"
	ErrMsgUserNotFound         = "ユーザーが見つかりません"
	ErrMsgInternalServer       = "内部サーバーエラーが発生しました"
	ErrMsgEncodeResponse       = "レスポンスのエンコードに失敗しました"
)

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
			http.Error(w, ErrMsgUserNotFound, http.StatusNotFound)
		default:
			http.Error(w, ErrMsgInternalServer, http.StatusInternalServerError)
		}
		return
	}

	if user == nil {
		http.Error(w, ErrMsgUserNotFound, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, ErrMsgEncodeResponse, http.StatusInternalServerError)
	}
}
