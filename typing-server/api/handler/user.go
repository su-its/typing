package handler

import (
	"encoding/json"
	"net/http"

	"github.com/su-its/typing/typing-server/api/service"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	studentNumber := r.URL.Query().Get("student_number")
	if studentNumber == "" {
		http.Error(w, "student_number is required", http.StatusBadRequest)
		return
	}

	user, err := service.GetUserByStudentNumber(ctx, studentNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
