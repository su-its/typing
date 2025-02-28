package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
	"github.com/su-its/typing/typing-server/internal/domain/service"
)

func TestNewScoreUseCase(t *testing.T) {
	type args struct {
		txManager    repository.TxManager
		scoreRepo    repository.ScoreRepository
		scoreService *service.ScoreService
	}
	tests := []struct {
		name string
		args args
		want *ScoreUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScoreUseCase(tt.args.txManager, tt.args.scoreRepo, tt.args.scoreService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScoreUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScoreUseCase_GetScoresRanking(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *model.GetScoresRankingRequest
	}
	tests := []struct {
		name    string
		uc      *ScoreUseCase
		args    args
		want    *model.GetScoresRankingResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetScoresRanking(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScoreUseCase.GetScoresRanking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScoreUseCase.GetScoresRanking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScoreUseCase_RegisterScore(t *testing.T) {
	type args struct {
		ctx        context.Context
		userID     uuid.UUID
		keystrokes int
		accuracy   float64
	}
	tests := []struct {
		name    string
		uc      *ScoreUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uc.RegisterScore(tt.args.ctx, tt.args.userID, tt.args.keystrokes, tt.args.accuracy); (err != nil) != tt.wantErr {
				t.Errorf("ScoreUseCase.RegisterScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
