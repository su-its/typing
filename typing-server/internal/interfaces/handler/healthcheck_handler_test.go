package handler

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewHealthCheckHandler(t *testing.T) {
	tests := []struct {
		name string
		want *HealthCheckHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHealthCheckHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthCheckHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHealthCheckHandler_LivenessProbe(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *HealthCheckHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.LivenessProbe(tt.args.w, tt.args.r)
		})
	}
}
