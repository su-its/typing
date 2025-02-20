package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
)

type mockScoreRepository struct {
	getScores    func(ctx context.Context, sortBy string, start int, limit int) ([]*model.Score, int, error)
	getMaxScores func(ctx context.Context, userID uuid.UUID) (*model.Score, *model.Score, error)
	createScore  func(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64, isMaxKeystrokes bool, isMaxAccuracy bool) error
}

func (m *mockScoreRepository) GetScores(ctx context.Context, sortBy string, start int, limit int) ([]*model.Score, int, error) {
	return m.getScores(ctx, sortBy, start, limit)
}

func (m *mockScoreRepository) GetMaxScores(ctx context.Context, userID uuid.UUID) (*model.Score, *model.Score, error) {
	return m.getMaxScores(ctx, userID)
}

func (m *mockScoreRepository) CreateScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64, isMaxKeystrokes bool, isMaxAccuracy bool) error {
	return m.createScore(ctx, userID, keystrokes, accuracy, isMaxKeystrokes, isMaxAccuracy)
}

func TestNewScoreService(t *testing.T) {
	type args struct {
		scoreRepo repository.ScoreRepository
	}
	tests := []struct {
		name string
		args args
		want *ScoreService
	}{
		// TODO: Add test cases.
		{
			name: "正常系: インスタンスが正しく生成される",
			args: args{
				scoreRepo: &mockScoreRepository{},
			},
			want: &ScoreService{
				scoreRepo: &mockScoreRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScoreService(tt.args.scoreRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScoreService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScoreService_ValidateScore(t *testing.T) {
	type args struct {
		userID     uuid.UUID
		keystrokes int
		accuracy   float64
	}
	tests := []struct {
		name    string
		s       *ScoreService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.ValidateScore(tt.args.userID, tt.args.keystrokes, tt.args.accuracy); (err != nil) != tt.wantErr {
				t.Errorf("ScoreService.ValidateScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScoreService_ComputeRanking(t *testing.T) {
	type args struct {
		scores []*model.Score
		sortBy string
		start  int
	}
	tests := []struct {
		name string
		s    *ScoreService
		args args
		want []*model.ScoreRanking
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ComputeRanking(tt.args.scores, tt.args.sortBy, tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScoreService.ComputeRanking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScoreService_ShouldUpdateMaxScore(t *testing.T) {
	type args struct {
		ctx      context.Context
		userID   uuid.UUID
		newScore *model.Score
	}
	tests := []struct {
		name    string
		s       *ScoreService
		args    args
		want    bool
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.ShouldUpdateMaxScore(tt.args.ctx, tt.args.userID, tt.args.newScore)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScoreService.ShouldUpdateMaxScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ScoreService.ShouldUpdateMaxScore() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ScoreService.ShouldUpdateMaxScore() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
