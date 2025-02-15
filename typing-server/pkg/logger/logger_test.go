package logger

import (
	"context"
	"log/slog"
	"reflect"
	"testing"
)

func TestNewTraceIDLogHandler(t *testing.T) {
	type args struct {
		h slog.Handler
	}
	tests := []struct {
		name string
		args args
		want *TraceIDLogHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTraceIDLogHandler(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTraceIDLogHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTraceIDLogHandler_Handle(t *testing.T) {
	type args struct {
		ctx context.Context
		r   slog.Record
	}
	tests := []struct {
		name    string
		h       *TraceIDLogHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Handle(tt.args.ctx, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("TraceIDLogHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *slog.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
