package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/app/service"
	"github.com/su-its/typing/typing-server/internal/domain/model"
)

func (h *Handler) GetScoresRanking(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// クエリパラメータの取得と検証
	sortBy := r.URL.Query().Get("sort_by")
	if sortBy == "" {
		sortBy = "keystrokes"
	}

	start, err := strconv.Atoi(r.URL.Query().Get("start"))
	if err != nil {
		start = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	// リクエストの構築
	request := model.GetScoresRankingRequest{
		SortBy: sortBy,
		Start:  start,
		Limit:  limit,
	}

	// サービス呼び出し
	response, err := service.GetScoresRanking(ctx, h.entClient, &request)
	if err != nil {
		h.log.Error("failed to get scores ranking",
			"error", err,
			"request", request)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// レスポンス返却
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.log.Error("failed to encode response",
			"error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) PostScore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// リクエストのデコード
	var requestBody struct {
		UserID     string  `json:"user_id"`
		Keystrokes int     `json:"keystrokes"`
		Accuracy   float64 `json:"accuracy"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		h.log.Error("failed to decode request body",
			"error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// UUIDの検証
	userID, err := uuid.Parse(requestBody.UserID)
	if err != nil {
		h.log.Error("invalid user_id",
			"error", err,
			"user_id", requestBody.UserID)
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// スコアの作成
	if err := service.CreateScore(ctx, h.entClient, userID, requestBody.Keystrokes, requestBody.Accuracy); err != nil {
		h.log.Error("failed to create score",
			"error", err,
			"user_id", userID,
			"keystrokes", requestBody.Keystrokes,
			"accuracy", requestBody.Accuracy)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
