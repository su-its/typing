package service

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
)

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
			if got := tt.s.ComputeRanking(tt.args.scores, tt.args.sortBy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScoreService.ComputeRanking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScoreService_LimitRankings(t *testing.T) {
	type args struct {
		rankings []*model.ScoreRanking
		start    int
		limit    int
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
			if got := tt.s.LimitRankings(tt.args.rankings, tt.args.start, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScoreService.LimitRankings() = %v, want %v", got, tt.want)
			}
		})
	}
}
