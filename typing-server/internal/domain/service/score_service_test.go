package service

import (
	"context"
	"reflect"
	"strings"
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
		name       string
		s          *ScoreService
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "正常系: 正しいパラメータの場合",
			s: &ScoreService{
				scoreRepo: &mockScoreRepository{},
			},
			args: args{
				userID:     uuid.New(),
				keystrokes: 100,
				accuracy:   0.5,
			},
			wantErr: false,
		},
		{
			name: "異常系: keystrokes が負の場合",
			s: &ScoreService{
				scoreRepo: &mockScoreRepository{},
			},
			args: args{
				userID:     uuid.New(),
				keystrokes: -10,
				accuracy:   0.5,
			},
			wantErr:    true,
			wantErrMsg: "keystrokes must be non-negative",
		},
		{
			name: "異常系: accuracy が 0未満の場合",
			s: &ScoreService{
				scoreRepo: &mockScoreRepository{},
			},
			args: args{
				userID:     uuid.New(),
				keystrokes: 100,
				accuracy:   -0.1,
			},
			wantErr: true,
		},
		{
			name: "異常系: accuracy が 1 を超える場合",
			s: &ScoreService{
				scoreRepo: &mockScoreRepository{},
			},
			args: args{
				userID:     uuid.New(),
				keystrokes: 100,
				accuracy:   1.1,
			},
			wantErr:    true,
			wantErrMsg: "accuracy must be between 0 and 1",
		},
		{
			name: "異常系: userID が uuid.Nilの場合",
			s: &ScoreService{
				scoreRepo: &mockScoreRepository{},
			},
			args: args{
				userID:     uuid.Nil,
				keystrokes: 100,
				accuracy:   0.5,
			},
			wantErr:    true,
			wantErrMsg: "invalid user ID\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.ValidateScore(tt.args.userID, tt.args.keystrokes, tt.args.accuracy)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateScore() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErrMsg != "" {
				if err == nil {
					t.Fatalf("expected error %q, but got nil", tt.wantErrMsg)
				}

				if !strings.Contains(err.Error(), strings.TrimSpace(tt.wantErrMsg)) {
					t.Errorf("expected error message %q, but got %q", tt.wantErrMsg, err.Error())
				}
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
		{
			name: "正常系: keystrokes で降順ソート、 start=1",
			s:    &ScoreService{},
			args: args{
				scores: []*model.Score{
					{
						ID:         "s1",
						Keystrokes: 200,
						Accuracy:   0.90,
					},
					{
						ID:         "s2",
						Keystrokes: 300,
						Accuracy:   0.85,
					},
					{
						ID:         "s3",
						Keystrokes: 100,
						Accuracy:   0.95,
					},
				},
				sortBy: "keystrokes",
				start:  1,
			},
			want: []*model.ScoreRanking{
				// s2: keystrokes=300
				{Rank: 1, Score: model.Score{ID: "s2", Keystrokes: 300, Accuracy: 0.85}},
				// s1: keystrokes=200
				{Rank: 2, Score: model.Score{ID: "s1", Keystrokes: 200, Accuracy: 0.90}},
				// s3: keystrokes=100
				{Rank: 3, Score: model.Score{ID: "s3", Keystrokes: 100, Accuracy: 0.95}},
			},
		},
		{
			name: "正常系: accuracy で降順ソート、 start=1, 重複accuracyあり",
			s:    &ScoreService{},
			args: args{
				scores: []*model.Score{
					{
						ID:         "s1",
						Keystrokes: 200,
						Accuracy:   0.90,
					},
					{
						ID:         "s2",
						Keystrokes: 500,
						Accuracy:   0.90, // s1 と同じ accuracy
					},
					{
						ID:         "s3",
						Keystrokes: 100,
						Accuracy:   0.95,
					},
				},
				sortBy: "accuracy",
				start:  1,
			},
			want: []*model.ScoreRanking{
				// s3: accuracy=0.95
				{Rank: 1, Score: model.Score{ID: "s3", Keystrokes: 100, Accuracy: 0.95}},
				// s1: accuracy=0.90 (上から2番目)
				{Rank: 2, Score: model.Score{ID: "s1", Keystrokes: 200, Accuracy: 0.90}},
				// s2: accuracy=0.90 (上と同じ accuracy => 現状の実装ではスコアが同じでも rank を別にしているので rank=3)
				{Rank: 2, Score: model.Score{ID: "s2", Keystrokes: 500, Accuracy: 0.90}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ComputeRanking(tt.args.scores, tt.args.sortBy, tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScoreService.ComputeRanking() = %+v, want %+v", got, tt.want)
				for i := range got {
					t.Logf("got[%d] = %+v", i, got[i])
				}
				for i := range tt.want {
					t.Logf("want[%d] = %+v", i, tt.want[i])
				}
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
