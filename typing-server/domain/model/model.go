package model

import "time"

type ScoreRanking struct {
	Rank  int   `json:"rank"`
	Score Score `json:"score"`
}

type User struct {
	ID            string    `json:"id"`
	StudentNumber string    `json:"student_number"`
	HandleName    string    `json:"handle_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Score struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Keystrokes int       `json:"keystrokes"`
	Accuracy   float64   `json:"accuracy"`
	CreatedAt  time.Time `json:"created_at"`
	User       User      `json:"user"`
}
