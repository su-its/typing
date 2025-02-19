package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated"
)

func TestNewEntUserRepository(t *testing.T) {
	type args struct {
		client *ent_generated.Client
	}
	tests := []struct {
		name string
		args args
		want *EntUserRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEntUserRepository(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEntUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntUserRepository_GetUserByStudentNumber(t *testing.T) {
	type args struct {
		ctx           context.Context
		studentNumber string
	}
	tests := []struct {
		name    string
		r       *EntUserRepository
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetUserByStudentNumber(tt.args.ctx, tt.args.studentNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("EntUserRepository.GetUserByStudentNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EntUserRepository.GetUserByStudentNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
