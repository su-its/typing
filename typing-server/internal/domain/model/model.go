package model

import "time"

type ScoreRanking struct {
	Rank  int   `json:"rank"`
	Score Score `json:"score"`
}

type User struct {
	ID            string    `json:"id"`
	// NOTE: 学籍番号はユニークであるべき．
	StudentNumber string    `json:"student_number"`
	HandleName    string    `json:"handle_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Score struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	// NOTE: キーストロークは正の値
	Keystrokes int       `json:"keystrokes"`
	// NOTE: 正確性は0.0から1.0の範囲であるべき．
	Accuracy   float64   `json:"accuracy"`
	CreatedAt  time.Time `json:"created_at"`
	User       User      `json:"user"`
}

type GetScoresRankingRequest struct {
	SortBy string `json:"sort_by"`
	Start  int    `json:"start"`
	Limit  int    `json:"limit"`
}

type GetScoresRankingResponse struct {
	Rankings   []*ScoreRanking `json:"rankings"`
	TotalCount int             `json:"total_count"`
}
