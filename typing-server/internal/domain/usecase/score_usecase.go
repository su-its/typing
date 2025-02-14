package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
	"github.com/su-its/typing/typing-server/internal/domain/service"
)

// ScoreUseCase はスコア関連のユースケース
type ScoreUseCase struct {
	scoreRepo    repository.ScoreRepository
	scoreService *service.ScoreService
}

// NewScoreUseCase は ScoreUseCase のインスタンスを生成する
func NewScoreUseCase(scoreRepo repository.ScoreRepository, scoreService *service.ScoreService) *ScoreUseCase {
	return &ScoreUseCase{
		scoreRepo:    scoreRepo,
		scoreService: scoreService,
	}
}

// GetScoresRanking はスコアランキングを取得する
func (uc *ScoreUseCase) GetScoresRanking(ctx context.Context, request *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error) {
	// DB からスコアを取得
	scores, totalCount, err := uc.scoreRepo.GetScores(ctx, request.SortBy, request.Start, request.Limit)
	if err != nil {
		return nil, err
	}

	// サービスでランキング計算
	rankings := uc.scoreService.ComputeRanking(scores, request.SortBy, request.Start)

	return &model.GetScoresRankingResponse{
		Rankings:   rankings,
		TotalCount: totalCount,
	}, nil
}

// RegisterScore はスコアを登録する
func (uc *ScoreUseCase) RegisterScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error {
	// スコアのバリデーション
	if err := uc.scoreService.ValidateScore(userID, keystrokes, accuracy); err != nil {
		return err
	}

	// スコアの作成
	newScore := &model.Score{
		UserID:     userID.String(),
		Keystrokes: keystrokes,
		Accuracy:   accuracy,
	}

	// 最大スコアの判定
	isMaxKeystrokes, isMaxAccuracy, err := uc.scoreService.ShouldUpdateMaxScore(ctx, userID, newScore)
	if err != nil {
		return err
	}

	// DB にスコアを保存
	return uc.scoreRepo.CreateScore(ctx, userID, keystrokes, accuracy, isMaxKeystrokes, isMaxAccuracy)
}
