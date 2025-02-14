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
	scoreUseCase *usecase.ScoreUseCase
}

// NewScoreHandler は ScoreHandler のインスタンスを生成する
func NewScoreHandler(scoreUseCase *usecase.ScoreUseCase) *ScoreHandler {
	return &ScoreHandler{scoreUseCase: scoreUseCase}
}

// GetScoresRanking はスコアランキングを取得するエンドポイント
func (h *ScoreHandler) GetScoresRanking(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	sortBy := r.URL.Query().Get("sort_by")
	startStr := r.URL.Query().Get("start")
	limitStr := r.URL.Query().Get("limit")

	// パラメータのバリデーション
	if sortBy != "keystrokes" && sortBy != "accuracy" {
		http.Error(w, "Invalid sort_by parameter", http.StatusBadRequest)
		return
	}

	start, err := strconv.Atoi(startStr)
	if err != nil || start <= 0 {
		http.Error(w, "Invalid start parameter", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
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
		http.Error(w, "Failed to fetch ranking", http.StatusInternalServerError)
		return
	}

	// JSON レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
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
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// UUID のバリデーション
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		http.Error(w, "Invalid user_id format", http.StatusBadRequest)
		return
	}

	// ユースケースを呼び出し
	err = h.scoreUseCase.RegisterScore(r.Context(), userID, req.Keystrokes, req.Accuracy)
	if err != nil {
		http.Error(w, "Failed to register score", http.StatusInternalServerError)
		return
	}

	// 成功時のレスポンス
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Score registered successfully"))
}
