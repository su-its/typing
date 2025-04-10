package handler

import (
	"log/slog"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"net/http/httptest"

	"github.com/su-its/typing/typing-server/internal/domain/usecase"
	"github.com/su-its/typing/typing-server/pkg/webhook"
)

func TestNewScoreHandler(t *testing.T) {
	type args struct {
		scoreUseCase    *usecase.ScoreUseCase
		log             *slog.Logger
		webhookNotifier *webhook.WebhookNotifier
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
				scoreUseCase:    fakeUseCase,
				log:             slog.Default(),
				webhookNotifier: &webhook.WebhookNotifier{},
			},
			want: &ScoreHandler{
				scoreUseCase: fakeUseCase,
				log:          slog.Default(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScoreHandler(tt.args.scoreUseCase, tt.args.log, tt.args.webhookNotifier); !reflect.DeepEqual(got, tt.want) {
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
		name       string
		h          *ScoreHandler
		args       args
		wantStatus int
		wantBody   string
	}{
		// TODO: Add test cases.
		{
			name: "sort_by が無効な場合は 400 が返る",
			h: &ScoreHandler{
				scoreUseCase: &usecase.ScoreUseCase{}, // 必要に応じてモック等に差し替え
			},
			args: args{
				w: httptest.NewRecorder(),
				// 例: sort_by=invalid をセットし、不正パラメータにしている
				r: httptest.NewRequest("GET", "/scores?sort_by=invalid&start=1&limit=10", nil),
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   ErrMsgInvalidSortByParameter,
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
			gotBody := rr.Body.String()
			if strings.TrimSpace(gotBody) != strings.TrimSpace(tt.wantBody) {
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
