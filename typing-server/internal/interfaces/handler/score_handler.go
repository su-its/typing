package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/usecase"
)

// ScoreHandler はスコア関連の HTTP ハンドラ
type ScoreHandler struct {
	scoreUseCase usecase.IScoreUseCase
}

// NewScoreHandler は ScoreHandler のインスタンスを生成する
func NewScoreHandler(scoreUseCase usecase.IScoreUseCase) *ScoreHandler {
	return &ScoreHandler{scoreUseCase: scoreUseCase}
}

const (
	ErrMsgInvalidSortbyParam	= "Invalid sort_by parameter"
	ErrMsgInvalidStartParam		= "Invalid start parameter"
	ErrMsgInvalidLimitParam		= "Invalid limit parameter"
	ErrMsgFetchRanking			= "Failed to fetch ranking"
	ErrMsgScoreEncodeResponse	= "Failed to encode response"

	ErrMsgInvalidReqBody		= "Invalid request body"
	ErrMsgInvalidUserIdFormat	= "Invalid user_id format"
	ErrMsgRegisterScore			= "Failed to register score"
	MsgRegisteredSuccessfully	= "Score registered successfully"
	ErrMsgWriteResponse			= "Failed to write response"
)
// GetScoresRanking はスコアランキングを取得するエンドポイント
func (h *ScoreHandler) GetScoresRanking(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	sortBy := r.URL.Query().Get("sort_by")
	startStr := r.URL.Query().Get("start")
	limitStr := r.URL.Query().Get("limit")

	// パラメータのバリデーション
	if sortBy != "keystrokes" && sortBy != "accuracy" {
		http.Error(w, ErrMsgInvalidSortbyParam, http.StatusBadRequest)
		return
	}

	start, err := strconv.Atoi(startStr)
	if err != nil || start <= 0 {
		http.Error(w, ErrMsgInvalidStartParam, http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		http.Error(w, ErrMsgInvalidLimitParam, http.StatusBadRequest)
		return
	}

	// ユースケースを呼び出し
	req := &model.GetScoresRankingRequest{
		SortBy: sortBy,
		Start:  start,
		Limit:  limit,
	}

	resp, err := h.scoreUseCase.GetScoresRanking(r.Context(), req)
	if err != nil {
		http.Error(w, ErrMsgFetchRanking, http.StatusInternalServerError)
		return
	}

	// JSON レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, ErrMsgScoreEncodeResponse, http.StatusInternalServerError)
	}
}

// RegisterScore はスコアを登録するエンドポイント
func (h *ScoreHandler) RegisterScore(w http.ResponseWriter, r *http.Request) {
	// リクエストボディをパース
	var req struct {
		UserID     string  `json:"user_id"`
		Keystrokes int     `json:"keystrokes"`
		Accuracy   float64 `json:"accuracy"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrMsgInvalidReqBody, http.StatusBadRequest)
		return
	}

	// UUID のバリデーション
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		http.Error(w, ErrMsgInvalidUserIdFormat, http.StatusBadRequest)
		return
	}

	// ユースケースを呼び出し
	err = h.scoreUseCase.RegisterScore(r.Context(), userID, req.Keystrokes, req.Accuracy)
	if err != nil {
		http.Error(w, ErrMsgRegisterScore, http.StatusInternalServerError)
		return
	}

	// 成功時のレスポンス
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(MsgRegisteredSuccessfully)); err != nil {
		http.Error(w, ErrMsgWriteResponse, http.StatusInternalServerError)
	}
}
