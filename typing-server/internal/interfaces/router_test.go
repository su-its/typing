package interfaces

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/su-its/typing/typing-server/config"
	"github.com/su-its/typing/typing-server/internal/interfaces/handler"
)

func TestNewRouter(t *testing.T) {
	type args struct {
		healthHandler *handler.HealthCheckHandler
		userHandler   *handler.UserHandler
		scoreHandler  *handler.ScoreHandler
		config        *config.Config
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(tt.args.healthHandler, tt.args.userHandler, tt.args.scoreHandler, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAllowedOrigins(t *testing.T) {
	type args struct {
		enviroment string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllowedOrigins(tt.args.enviroment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllowedOrigins() = %v, want %v", got, tt.want)
			}
		})
	}
}
