package handler

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/su-its/typing/typing-server/internal/domain/usecase"
)

func TestNewUserHandler(t *testing.T) {
	type args struct {
		userUseCase *usecase.UserUseCase
	}
	tests := []struct {
		name string
		args args
		want *UserHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.userUseCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHandler_GetUserByStudentNumber(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *UserHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetUserByStudentNumber(tt.args.w, tt.args.r)
		})
	}
}
