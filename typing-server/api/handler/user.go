package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/api/service"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	studentNumber := r.URL.Query().Get("student_number")
	if studentNumber == "" {
		http.Error(w, "student_number is required", http.StatusBadRequest)
		return
	}

	user, err := service.GetUserByStudentNumber(ctx, entClient, studentNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetMyScoreRanking(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := chi.URLParam(r, "user-id")
	if userID == "" {
		http.Error(w, "user-id is required", http.StatusBadRequest)
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "Invalid user-id", http.StatusBadRequest)
		return
	}

	sortBy := r.URL.Query().Get("sort_by")
	if sortBy == "" {
		sortBy = "keystrokes"
	}

	currentRank, err := service.GetMyScoreRanking(ctx, entClient, userUUID, sortBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		CurrentRank int `json:"current-rank"`
	}{
		CurrentRank: currentRank,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
