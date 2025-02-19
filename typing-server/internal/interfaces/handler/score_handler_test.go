package handler

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"
	"strings"

	"net/http/httptest"
	"github.com/google/uuid"

	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/usecase"
)

type mockScoreUseCase struct {
	getScoresRanking func(ctx context.Context, request *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error)
    registerScore func(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error
}

func (m *mockScoreUseCase) GetScoresRanking(ctx context.Context, req *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error) {
	return m.getScoresRanking(ctx,req)
}

// RegisterScore は今回は使わないため簡易実装
func (m *mockScoreUseCase) RegisterScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error {
	return m.registerScore(ctx,userID,keystrokes,accuracy)
}

func TestNewScoreHandler(t *testing.T) {
	type args struct {
		scoreUseCase usecase.IScoreUseCase
	}
	tests := []struct {
		name string
		args args
		want *ScoreHandler
	}{
		{
			name: "正常系: ScoreHandlerが正しく生成される",
			args: args{
				scoreUseCase: &mockScoreUseCase{},
			},
			want: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{},
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
		wantStatus int
		wantBody   string
	}{
		// TODO: Add test cases.
		{
			name: "正常系: スコアランキングが取得できる場合",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					getScoresRanking: func(ctx context.Context, request *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error) {
						return &model.GetScoresRankingResponse{
							Rankings: []*model.ScoreRanking{
								{
									Rank: 1,
									Score: model.Score{
										ID:         "score-1",
										UserID:     "user-1",
										Keystrokes: 300,
										Accuracy:   0.95,
										CreatedAt:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
										User: model.User{
											ID:            "1",
											StudentNumber: "k20000",
											HandleName:    "テストユーザー",
											CreatedAt:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
											UpdatedAt:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
										},
									},
								},
							},
							TotalCount: 1,
						}, nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/scores/ranking?sort_by=keystrokes&start=1&limit=10", nil),
			},
			wantStatus: http.StatusOK,
			wantBody: `{"rankings":[{"rank":1,"score":{"id":"score-1","user_id":"user-1","keystrokes":300,"accuracy":0.95,"created_at":"2021-01-01T00:00:00Z","user":{"id":"1","student_number":"k20000","handle_name":"テストユーザー","created_at":"2021-01-01T00:00:00Z","updated_at":"2021-01-01T00:00:00Z"}}}],"total_count":1}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetScoresRanking(tt.args.w, tt.args.r)
			rr := tt.args.w.(*httptest.ResponseRecorder)

			// ステータスコードの検証
			if rr.Code != tt.wantStatus {
				t.Errorf("GetScoresRanking() status code = %v, want %v",
					rr.Code, tt.wantStatus)
			}
			gotBody := strings.TrimSpace(rr.Body.String())
			if gotBody != tt.wantBody {
				t.Errorf("GetScoresRanking() body = %q, want %q", gotBody, tt.wantBody)
			}
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
