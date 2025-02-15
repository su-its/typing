package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
)

func TestNewUserUseCase(t *testing.T) {
	type args struct {
		userRepo repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want *UserUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUseCase(tt.args.userRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUseCase_GetUserByStudentNumber(t *testing.T) {
	type args struct {
		ctx           context.Context
		studentNumber string
	}
	tests := []struct {
		name    string
		uc      *UserUseCase
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetUserByStudentNumber(tt.args.ctx, tt.args.studentNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUseCase.GetUserByStudentNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUseCase.GetUserByStudentNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
