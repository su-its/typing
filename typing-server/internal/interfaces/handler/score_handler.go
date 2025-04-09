package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/usecase"
	"github.com/su-its/typing/typing-server/pkg/webhook"
)

// ScoreHandler はスコア関連の HTTP ハンドラ
type ScoreHandler struct {
	scoreUseCase    *usecase.ScoreUseCase
	log             *slog.Logger
	webhookNotifier *webhook.WebhookNotifier
}

// NewScoreHandler は ScoreHandler のインスタンスを生成する
func NewScoreHandler(scoreUseCase *usecase.ScoreUseCase, log *slog.Logger, webhookNotifier *webhook.WebhookNotifier) *ScoreHandler {
	return &ScoreHandler{scoreUseCase: scoreUseCase, log: log, webhookNotifier: webhookNotifier}
}

// GetScoresRanking はスコアランキングを取得するエンドポイント
func (h *ScoreHandler) GetScoresRanking(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	sortBy := r.URL.Query().Get("sort_by")
	startStr := r.URL.Query().Get("start")
	limitStr := r.URL.Query().Get("limit")

	// パラメータのバリデーション
	if sortBy != "keystrokes" && sortBy != "accuracy" {
		http.Error(w, ErrMsgInvalidSortByParameter, http.StatusBadRequest)
		return
	}

	start, err := strconv.Atoi(startStr)
	if err != nil || start <= 0 {
		http.Error(w, ErrMsgInvalidStartParameter, http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		http.Error(w, ErrMsgInvalidLimitParameter, http.StatusBadRequest)
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
		h.log.Error("GetScoresRanking failed", "error", err)
		http.Error(w, ErrInternalServer, http.StatusInternalServerError)
		return
	}

	// JSON レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.log.Error("Failed to encode JSON response", "error", err)
		http.Error(w, ErrFailedToEncodeResponse, http.StatusInternalServerError)
	}
}

// RegisterScore はスコアを登録するエンドポイント
func (h *ScoreHandler) RegisterScore(w http.ResponseWriter, r *http.Request) {
	// リクエストボディをパース
	var reqBody struct {
		UserID     string  `json:"user_id"`
		Keystrokes int     `json:"keystrokes"`
		Accuracy   float64 `json:"accuracy"`
	}

	// リクエストボディを読み込む（後でWebhookで使用するため）
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Failed to read request body", "error", err)
		http.Error(w, ErrMsgInvalidRequestBody, http.StatusBadRequest) // Assume ErrMsgInvalidRequestBody exists
		return
	}
	// 一度読み込んだボディを再度設定し直す
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Restore the body

	if err := json.Unmarshal(bodyBytes, &reqBody); err != nil {
		http.Error(w, ErrMsgInvalidRequestBody, http.StatusBadRequest) // Assume ErrMsgInvalidRequestBody exists
		return
	}

	// --- Webhook通知判定 ---
	userAgent := r.Header.Get("User-Agent")
	if !webhook.IsBrowserRequest(userAgent) {
		// 非ブラウザからのアクセスと判定された場合、非同期でWebhook通知
		headersCopy := r.Header.Clone()
		go h.webhookNotifier.SendNonBrowserScoreNotification(&http.Request{Header: headersCopy, RemoteAddr: r.RemoteAddr}, reqBody)
	}
	// --- Webhook通知判定 ここまで ---

	// UUID のバリデーション
	userID, err := uuid.Parse(reqBody.UserID)
	if err != nil {
		http.Error(w, ErrMsgInvalidUserIDParameter, http.StatusBadRequest) // Assume ErrMsgInvalidUserIDParameter exists
		return
	}

	// ユースケースを呼び出し
	err = h.scoreUseCase.RegisterScore(r.Context(), userID, reqBody.Keystrokes, reqBody.Accuracy)
	if err != nil {
		h.log.Error("RegisterScore failed", "error", err)
		http.Error(w, ErrFailedToRegisterScore, http.StatusInternalServerError) // Assume ErrFailedToRegisterScore exists
		return
	}

	// 成功時のレスポンス
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": SuccessMsgScoreRegistered} // Assume SuccessMsgScoreRegistered exists
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.log.Error("Failed to encode JSON response", "error", err)
		http.Error(w, ErrFailedToEncodeResponse, http.StatusInternalServerError) // Assume ErrFailedToEncodeResponse exists
	}
}

// GetUserScores はユーザーのスコアを取得するエンドポイント
func (h *ScoreHandler) GetUserScores(w http.ResponseWriter, r *http.Request) {
	// TODO: 実装する https://github.com/su-its/typing/issues/184
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
