package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/su-its/typing/typing-server/api/service"
	"github.com/su-its/typing/typing-server/domain/repository/ent"
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
		limit = 50
	}

	rankings, err := service.GetScoresRanking(ctx, sortBy, start, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rankings)
}

func PostScore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var score ent.Score
	if err := json.NewDecoder(r.Body).Decode(&score); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.CreateScore(ctx, &score); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
