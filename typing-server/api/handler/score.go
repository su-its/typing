package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/api/service"
	"github.com/su-its/typing/typing-server/domain/model"
)

func GetScoresRanking(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sortBy := r.URL.Query().Get("sort_by")
	if sortBy == "" {
		sortBy = "keystrokes"
	}

	startStr := r.URL.Query().Get("start")
	start, err := strconv.Atoi(startStr)
	if err != nil {
		start = 1
	}

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	request := model.GetScoresRankingRequest{
		SortBy: sortBy,
		Start:  start,
		Limit:  limit,
	}

	response, err := service.GetScoresRanking(ctx, entClient, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostScore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// リクエストボディから値を取得
	var requestBody struct {
		UserID     string  `json:"user_id"`
		Keystrokes int     `json:"keystrokes"`
		Accuracy   float64 `json:"accuracy"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// user_idをUUIDに変換
	userID, err := uuid.Parse(requestBody.UserID)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// スコアを作成
	if err := service.CreateScore(ctx, entClient, userID, requestBody.Keystrokes, requestBody.Accuracy); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
