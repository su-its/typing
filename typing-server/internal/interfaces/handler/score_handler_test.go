package handler

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/su-its/typing/typing-server/internal/domain/usecase"
)

func TestNewScoreHandler(t *testing.T) {
	type args struct {
		scoreUseCase *usecase.ScoreUseCase
	}
	fakeUseCase := &usecase.ScoreUseCase{}
	tests := []struct {
		name string
		args args
		want *ScoreHandler
	}{
		// TODO: Add test cases.
		{
			name: "正常系: ScoreHandlerが正しく生成される",
			args: args{
				scoreUseCase: fakeUseCase,
			},
			want: &ScoreHandler{
				scoreUseCase: fakeUseCase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScoreHandler(tt.args.scoreUseCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScoreHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScoreHandler_GetScoresRanking(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *ScoreHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetScoresRanking(tt.args.w, tt.args.r)
		})
	}
}

func TestScoreHandler_RegisterScore(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *ScoreHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.RegisterScore(tt.args.w, tt.args.r)
		})
	}
}
