package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/api/service"
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

	rankings, err := service.GetScoresRanking(ctx, sortBy, start, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rankings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostScore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIDStr := r.URL.Query().Get("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	keystrokesStr := r.URL.Query().Get("keystrokes")
	keystrokes, err := strconv.Atoi(keystrokesStr)
	if err != nil {
		http.Error(w, "Invalid keystrokes", http.StatusBadRequest)
		return
	}

	accuracyStr := r.URL.Query().Get("accuracy")
	accuracy, err := strconv.ParseFloat(accuracyStr, 64)
	if err != nil {
		http.Error(w, "Invalid accuracy", http.StatusBadRequest)
		return
	}

	if err := service.CreateScore(ctx, userID, keystrokes, accuracy); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
