package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/su-its/typing/typing-server/internal/infra/ent/ent_generated"
)

func TestNewEntTxManager(t *testing.T) {
	type args struct {
		client *ent_generated.Client
	}
	tests := []struct {
		name string
		args args
		want *EntTxManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEntTxManager(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEntTxManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntTxManager_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		fn  func(ctx context.Context) error
	}
	tests := []struct {
		name    string
		tm      *EntTxManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tm.Execute(tt.args.ctx, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("EntTxManager.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
