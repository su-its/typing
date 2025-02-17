package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"
	"errors"

	"net/http/httptest"

	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/usecase"
)

func TestNewUserHandler(t *testing.T) {
	type args struct {
		userUseCase usecase.IUserUseCase
	}
	tests := []struct {
		name string
		args args
		want *UserHandler
	}{
		{
			name: "正常系: UserHandlerが正しく生成される",
			args: args{
				userUseCase: &mockUserUseCase{},
			},
			want: &UserHandler{
				userUseCase: &mockUserUseCase{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.userUseCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockUserUseCase struct {
	getUserByStudentNumber func(ctx context.Context, studentNumber string) (*model.User, error)
}

func (m *mockUserUseCase) GetUserByStudentNumber(ctx context.Context, studentNumber string) (*model.User, error) {
	return m.getUserByStudentNumber(ctx, studentNumber)
}

func TestUserHandler_GetUserByStudentNumber(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name       string
		h          *UserHandler
		args       args
		wantStatus int
		wantBody   string
	}{
		{
			name: "正常系: ユーザーが見つかる場合",
			h: &UserHandler{
				userUseCase: &mockUserUseCase{
					getUserByStudentNumber: func(ctx context.Context, studentNumber string) (*model.User, error) {
						return &model.User{
							ID:            "1",
							StudentNumber: "k20000",
							HandleName:    "テストユーザー",
							CreatedAt:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
							UpdatedAt:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						}, nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/?student_number=k20000", nil),
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"id":"1","student_number":"k20000","handle_name":"テストユーザー","created_at":"2021-01-01T00:00:00Z","updated_at":"2021-01-01T00:00:00Z"}`,
		},
		{
			name: "異常系: student_numberが指定されていない場合",
			h: &UserHandler{
				userUseCase: &usecase.UserUseCase{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/", nil),
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "student_numberが指定されていません\n",
		},
		{
			name: "異常系: ユーザーが見つからない場合",
			h: &UserHandler{
				userUseCase: &mockUserUseCase{
					getUserByStudentNumber: func(ctx context.Context, studentNumber string) (*model.User, error) {
						return nil, usecase.ErrUserNotFound
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/?student_number=k99999", nil),
			},
			wantStatus: http.StatusNotFound,
			wantBody:   "ユーザーが見つかりません\n",
		},
		{
			name: "異常系: useCaseが想定外のエラーを吐いたとき",
			h: &UserHandler{
				userUseCase: &mockUserUseCase{
					getUserByStudentNumber: func(ctx context.Context, studentNumber string) (*model.User, error) {
						return nil, errors.New("hoge")
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/?student_number=k99999", nil),
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   "内部サーバーエラーが発生しました\n",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			tt.h.GetUserByStudentNumber(rec, tt.args.r)

			if rec.Code != tt.wantStatus {
				t.Errorf("GetUserByStudentNumber() status = %v, want %v", rec.Code, tt.wantStatus)
			}

			if tt.wantStatus == http.StatusOK {
				var got, want model.User
				if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
					t.Errorf("Failed to parse response body: %v", err)
				}
				if err := json.Unmarshal([]byte(tt.wantBody), &want); err != nil {
					t.Errorf("Failed to parse expected body: %v", err)
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("GetUserByStudentNumber() = %+v, want %+v", got, want)
				}
			} else {
				if rec.Body.String() != tt.wantBody {
					t.Errorf("GetUserByStudentNumber() body = %v, want %v", rec.Body.String(), tt.wantBody)
				}
			}
		})
	}
}
