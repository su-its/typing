package model

import "time"

type ScoreRanking struct {
	Rank      int       `json:"rank"`
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Score     int       `json:"score"`
	Accuracy  float64   `json:"accuracy"`
	CreatedAt time.Time `json:"created_at"`
}
