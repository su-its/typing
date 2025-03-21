package service

import (
	"context"
	"errors"
	"sort"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
)

// ScoreService はスコアのビジネスロジックを管理する
type ScoreService struct {
	scoreRepo repository.ScoreRepository
}

// NewScoreService は ScoreService のインスタンスを作成する
func NewScoreService(scoreRepo repository.ScoreRepository) *ScoreService {
	return &ScoreService{scoreRepo: scoreRepo}
}

// ValidateScore はスコアが有効かどうかを判定する
func (s *ScoreService) ValidateScore(userID uuid.UUID, keystrokes int, accuracy float64) error {
	if keystrokes < 0 {
		return errors.New("keystrokes must be non-negative")
	}
	if accuracy < 0 || accuracy > 1 {
		return errors.New("accuracy must be between 0 and 1")
	}
	if userID == uuid.Nil {
		return errors.New("invalid user ID")
	}
	return nil
}

// ComputeRanking はスコアランキングの順位を計算する
func (s *ScoreService) ComputeRanking(scores []*model.Score, sortBy string, start int) []*model.ScoreRanking {
	// ソート処理
	switch sortBy {
	case "accuracy":
		sort.Slice(scores, func(i, j int) bool { return scores[i].Accuracy > scores[j].Accuracy })
	case "keystrokes":
		sort.Slice(scores, func(i, j int) bool { return scores[i].Keystrokes > scores[j].Keystrokes })
	}

	// ランキングを計算
	rankings := make([]*model.ScoreRanking, 0, len(scores))
	var prevScore float64
	var rank int

	for i, s := range scores {
		var currentScore float64
		switch sortBy {
		case "accuracy":
			currentScore = s.Accuracy
		case "keystrokes":
			currentScore = float64(s.Keystrokes)
		}

		if i == 0 || currentScore != prevScore {
			rank = start + i
		}
		prevScore = currentScore

		rankings = append(rankings, &model.ScoreRanking{
			Rank:  rank,
			Score: *s,
		})
	}

	return rankings
}

// ShouldUpdateMaxScore は新しいスコアが最大スコアかどうかを判定する
func (s *ScoreService) ShouldUpdateMaxScore(ctx context.Context, userID uuid.UUID, newScore *model.Score) (bool, bool, error) {
	// 現在の最大スコアを取得
	maxKeystrokeScore, maxAccuracyScore, err := s.scoreRepo.GetMaxScores(ctx, userID)
	if err != nil {
		return false, false, err
	}

	isMaxKeystrokes := maxKeystrokeScore == nil || newScore.Keystrokes > maxKeystrokeScore.Keystrokes
	isMaxAccuracy := maxAccuracyScore == nil || newScore.Accuracy > maxAccuracyScore.Accuracy

	return isMaxKeystrokes, isMaxAccuracy, nil
}
