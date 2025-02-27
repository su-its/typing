package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"net/http/httptest"

	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/usecase"
	"github.com/su-its/typing/typing-server/internal/testutils"
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
	createUser             func(ctx context.Context, studentNumber string, handleName string) (*model.User, error)
}

func (m *mockUserUseCase) GetUserByStudentNumber(ctx context.Context, studentNumber string) (*model.User, error) {
	return m.getUserByStudentNumber(ctx, studentNumber)
}

func (m *mockUserUseCase) CreateUser(ctx context.Context, studentNumber string, handleName string) (*model.User, error) {
	if m.createUser != nil {
		return m.createUser(ctx, studentNumber, handleName)
	}
	return &model.User{
		StudentNumber: studentNumber,
		HandleName:    handleName,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
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
			wantBody:   ErrMsgStudentNumberRequired,
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
			wantBody:   ErrMsgUserNotFound,
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
			wantBody:   ErrMsgInternalServer,
		},
		{
			name: "異常系: userがnilのとき",
			h: &UserHandler{
				userUseCase: &mockUserUseCase{
					getUserByStudentNumber: func(ctx context.Context, studentNumber string) (*model.User, error) {
						return nil, nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/?student_number=k99999", nil),
			},
			wantStatus: http.StatusNotFound,
			wantBody:   ErrMsgUserNotFound,
		},
		{
			name: "異常系: レスポンスのエンコードが失敗したとき",
			h: &UserHandler{
				userUseCase: &mockUserUseCase{
					getUserByStudentNumber: func(ctx context.Context, studentNumber string) (*model.User, error) {
						return &model.User{
							ID:            "1",
							StudentNumber: "k20000",
							HandleName:    "テストユーザー",
							CreatedAt:     time.Now(),
							UpdatedAt:     time.Now(),
						}, nil
					},
				},
			},
			args: args{
				w: testutils.NewFakeResponseWriter(),
				r: httptest.NewRequest("GET", "/?student_number=k99999", nil),
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   ErrMsgEncodeResponse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetUserByStudentNumber(tt.args.w, tt.args.r)

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

			if code != tt.wantStatus {
				t.Errorf("GetUserByStudentNumber() status = %v, want %v", code, tt.wantStatus)
			}

			if tt.wantStatus == http.StatusOK {
				var got, want model.User
				if err := json.Unmarshal([]byte(body), &got); err != nil {
					t.Errorf("Failed to parse response body: %v", err)
				}
				if err := json.Unmarshal([]byte(tt.wantBody), &want); err != nil {
					t.Errorf("Failed to parse expected body: %v", err)
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("GetUserByStudentNumber() = %+v, want %+v", got, want)
				}
			} else {
				if !strings.Contains(body, strings.TrimSpace(tt.wantBody)) {
					t.Errorf("GetUserByStudentNumber() body = %q, want to contain %q", body, tt.wantBody)
				}
			}
		})
	}
}
