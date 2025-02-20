package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"net/http/httptest"

	"github.com/google/uuid"

	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/usecase"
	"github.com/su-its/typing/typing-server/internal/testutils"
)

type mockScoreUseCase struct {
	getScoresRanking func(ctx context.Context, request *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error)
	registerScore    func(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error
}

func (m *mockScoreUseCase) GetScoresRanking(ctx context.Context, req *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error) {
	return m.getScoresRanking(ctx, req)
}

// RegisterScore は今回は使わないため簡易実装
func (m *mockScoreUseCase) RegisterScore(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error {
	return m.registerScore(ctx, userID, keystrokes, accuracy)
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
		name       string
		h          *ScoreHandler
		args       args
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
			wantBody:   `{"rankings":[{"rank":1,"score":{"id":"score-1","user_id":"user-1","keystrokes":300,"accuracy":0.95,"created_at":"2021-01-01T00:00:00Z","user":{"id":"1","student_number":"k20000","handle_name":"テストユーザー","created_at":"2021-01-01T00:00:00Z","updated_at":"2021-01-01T00:00:00Z"}}}],"total_count":1}`,
		},
		{
			name: "異常系: sort_byが不正な場合",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					getScoresRanking: func(ctx context.Context, request *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error) {
						t.Error("不正なsort_byの場合、ユースケースは呼び出されるべきではない")
						return nil, nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/scores/ranking?sort_by=invalid&start=1&limit=10", nil),
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid sort_by parameter",
		},
		{
			name: "異常系: startが不正(数字変換エラー)の場合",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					getScoresRanking: func(ctx context.Context, request *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error) {
						t.Error("不正なstartの場合、ユースケースは呼び出されるべきではない")
						return nil, nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				// start=abc は数字変換失敗
				r: httptest.NewRequest("GET", "/scores/ranking?sort_by=accuracy&start=abc&limit=10", nil),
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid start parameter",
		},
		{
			name: "異常系: startが不正(0以下)の場合",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					getScoresRanking: func(ctx context.Context, request *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error) {
						t.Error("不正なstartの場合、ユースケースは呼び出されるべきではない")
						return nil, nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				// start=-1 は 0以下
				r: httptest.NewRequest("GET", "/scores/ranking?sort_by=accuracy&start=-1&limit=10", nil),
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid start parameter",
		},
		{
			name: "異常系: limitが不正(数字変換エラー)の場合",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					getScoresRanking: func(ctx context.Context, request *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error) {
						t.Error("不正なlimitの場合、ユースケースは呼び出されるべきではない")
						return nil, nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/scores/ranking?sort_by=accuracy&start=1&limit=abc", nil),
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid limit parameter",
		},
		{
			name: "異常系: limitが不正(0以下)の場合",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					getScoresRanking: func(ctx context.Context, request *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error) {
						t.Error("不正なlimitの場合、ユースケースは呼び出されるべきではない")
						return nil, nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/scores/ranking?sort_by=accuracy&start=1&limit=0", nil),
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid limit parameter",
		},
		{
			name: "異常系: GetScoresRankingをしたときユースケースからエラーが返る場合",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					getScoresRanking: func(ctx context.Context, request *model.GetScoresRankingRequest) (*model.GetScoresRankingResponse, error) {
						return nil, errors.New("ErrGetScoresRanking") // ここで適当なエラーを返す
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/scores?sort_by=accuracy&start=1&limit=10", nil),
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   "Failed to fetch ranking",
		},
		{
			name: "異常系: レスポンスのエンコードが失敗したとき",
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
				w: testutils.NewFakeResponseWriter(),
				r: httptest.NewRequest("GET", "/scores/ranking?sort_by=keystrokes&start=1&limit=10", nil),
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   "Failed to encode response",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetScoresRanking(tt.args.w, tt.args.r)
			var code int
			var body string
			switch w := tt.args.w.(type) {
			case *testutils.FakeResponseWriter:
				code = w.StatusCode
				body = w.Body.String()
			case *httptest.ResponseRecorder:
				code = w.Code
				body = w.Body.String()
			default:
				t.Fatal("unknown ResponseWriter type")
			}
			// ステータスコードの検証
			if code != tt.wantStatus {
				t.Errorf("GetScoresRanking() status code = %v, want %v",
					code, tt.wantStatus)
			}
			if tt.wantStatus == http.StatusOK {
				var got, want model.Score
				if err := json.Unmarshal([]byte(body), &got); err != nil {
					t.Errorf("Failed to parse response body: %v", err)
				}
				if err := json.Unmarshal([]byte(tt.wantBody), &want); err != nil {
					t.Errorf("Failed to parse expected body: %v", err)
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("GetScoresRanking() = %+v, want %+v", got, want)
				}
			} else {
				if !strings.Contains(body, strings.TrimSpace(tt.wantBody)) {
					t.Errorf("GetScoresRanking() body = %q, want to contain %q", body, tt.wantBody)
				}
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
		name       string
		h          *ScoreHandler
		args       args
		wantStatus int
		wantBody   string
	}{
		{
			name: "正常系: スコア登録が成功する場合",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					registerScore: func(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error {
						return nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					bodyMap := map[string]interface{}{
						"user_id":    "b110a730-93d6-4dac-a8b4-c9a9fc5cb1bf", // 正しい UUID の例
						"keystrokes": 300,
						"accuracy":   0.95,
					}
					jsonBody, _ := json.Marshal(bodyMap)
					req := httptest.NewRequest("POST", "/scores", bytes.NewReader(jsonBody))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusCreated,
			wantBody:   "Score registered successfully",
		},
		{
			name: "異常系: JSONパースエラーの場合",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					registerScore: func(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error {
						t.Error("JSONパースに失敗した場合、ユースケースは呼び出されるべきではない")
						return nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				// 不正なJSON (キーがダブルクォートで囲まれていないなど)
				r: httptest.NewRequest("POST", "/scores", strings.NewReader(`{user_id:"xxx"}`)),
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid request body\n",
		},
		{
			name: "異常系: UUIDのバリデーションエラーの場合",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					registerScore: func(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error {
						t.Error("UUID不正の場合、ユースケースは呼び出されるべきではない")
						return nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				// 不正な user_id
				r: func() *http.Request {
					bodyMap := map[string]interface{}{
						"user_id":    "invalid-uuid",
						"keystrokes": 300,
						"accuracy":   0.95,
					}
					jsonBody, _ := json.Marshal(bodyMap)
					req := httptest.NewRequest("POST", "/scores", bytes.NewReader(jsonBody))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid user_id format\n",
		},
		{
			name: "異常系: ユースケースがエラーを返す場合",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					registerScore: func(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error {
						return errors.New("hoge")
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					bodyMap := map[string]interface{}{
						"user_id":    "b110a730-93d6-4dac-a8b4-c9a9fc5cb1bf",
						"keystrokes": 300,
						"accuracy":   0.95,
					}
					jsonBody, _ := json.Marshal(bodyMap)
					req := httptest.NewRequest("POST", "/scores", bytes.NewReader(jsonBody))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   "Failed to register score\n",
		},
		{
			name: "異常系: レスポンスの書き込みが失敗したとき",
			h: &ScoreHandler{
				scoreUseCase: &mockScoreUseCase{
					registerScore: func(ctx context.Context, userID uuid.UUID, keystrokes int, accuracy float64) error {
						return nil
					},
				},
			},
			args: args{
				w: testutils.NewFakeResponseWriter(),
				r: func() *http.Request {
					bodyMap := map[string]interface{}{
						"user_id":    "b110a730-93d6-4dac-a8b4-c9a9fc5cb1bf", // 正しい UUID の例
						"keystrokes": 300,
						"accuracy":   0.95,
					}
					jsonBody, _ := json.Marshal(bodyMap)
					req := httptest.NewRequest("POST", "/scores", bytes.NewReader(jsonBody))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   "Failed to write response\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.RegisterScore(tt.args.w, tt.args.r)
			var code int
			var body string
			switch w := tt.args.w.(type) {
			case *testutils.FakeResponseWriter:
				code = w.StatusCode
				body = w.Body.String()
			case *httptest.ResponseRecorder:
				code = w.Code
				body = w.Body.String()
			default:
				t.Fatal("unknown ResponseWriter type")
			}
			// ステータスコードの検証
			if code != tt.wantStatus {
				t.Errorf("RegisterScore() status code = %v, want %v",
					code, tt.wantStatus)
			}
			if tt.wantStatus == http.StatusOK {
				var got, want model.Score
				if err := json.Unmarshal([]byte(body), &got); err != nil {
					t.Errorf("Failed to parse response body: %v", err)
				}
				if err := json.Unmarshal([]byte(tt.wantBody), &want); err != nil {
					t.Errorf("Failed to parse expected body: %v", err)
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("RegisterScore() = %+v, want %+v", got, want)
				}
			} else {
				if !strings.Contains(body, strings.TrimSpace(tt.wantBody)) {
					t.Errorf("RegisterScore() body = %q, want to contain %q", body, tt.wantBody)
				}
			}
		})
	}
}
