package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated"
)

func TestNewEntScoreRepository(t *testing.T) {
	type args struct {
		client *ent_generated.Client
	}
	tests := []struct {
		name string
		args args
		want *EntScoreRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEntScoreRepository(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEntScoreRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntScoreRepository_GetScores(t *testing.T) {
	type args struct {
		ctx    context.Context
		sortBy string
		start  int
		limit  int
	}
	tests := []struct {
		name    string
		r       *EntScoreRepository
		args    args
		want    []*model.Score
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.r.GetScores(tt.args.ctx, tt.args.sortBy, tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("EntScoreRepository.GetScores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EntScoreRepository.GetScores() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("EntScoreRepository.GetScores() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEntScoreRepository_GetMaxScores(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *EntScoreRepository
		args    args
		want    *model.Score
		want1   *model.Score
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.r.GetMaxScores(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("EntScoreRepository.GetMaxScores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EntScoreRepository.GetMaxScores() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("EntScoreRepository.GetMaxScores() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEntScoreRepository_CreateScore(t *testing.T) {
	type args struct {
		ctx             context.Context
		userID          uuid.UUID
		keystrokes      int
		accuracy        float64
		isMaxKeystrokes bool
		isMaxAccuracy   bool
	}
	tests := []struct {
		name    string
		r       *EntScoreRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.CreateScore(tt.args.ctx, tt.args.userID, tt.args.keystrokes, tt.args.accuracy, tt.args.isMaxKeystrokes, tt.args.isMaxAccuracy); (err != nil) != tt.wantErr {
				t.Errorf("EntScoreRepository.CreateScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
