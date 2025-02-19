package middleware

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDefaultCORSConfig(t *testing.T) {
	tests := []struct {
		name string
		want CORSConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultCORSConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultCORSConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCORSMiddleware(t *testing.T) {
	type args struct {
		config CORSConfig
	}
	tests := []struct {
		name string
		args args
		want func(http.Handler) http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CORSMiddleware(tt.args.config); fmt.Sprintf("%T", got) != fmt.Sprintf("%T", tt.want) {
				t.Errorf("CORSMiddleware() = %T, want %T", got, tt.want)
			}
		})
	}
}
