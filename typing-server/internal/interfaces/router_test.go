package interfaces

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/su-its/typing/typing-server/config"
	"github.com/su-its/typing/typing-server/internal/interfaces/handler"
)

func TestNewRouter(t *testing.T) {
	// テストケースのセットアップ
	healthHandler := &handler.HealthCheckHandler{}
	userHandler := &handler.UserHandler{}
	scoreHandler := &handler.ScoreHandler{}
	cfg := &config.Config{
		Environment: "local",
	}

	router := NewRouter(healthHandler, userHandler, scoreHandler, cfg)

	// 各エンドポイントのテスト
	testCases := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "ヘルスチェックエンドポイント",
			method:         "GET",
			path:           "/health",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "存在しないエンドポイント",
			method:         "GET",
			path:           "/not-exists",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("期待したステータスコード %d, 実際のステータスコード %d", tc.expectedStatus, w.Code)
			}
		})
	}
}

func Test_getAllowedOrigins(t *testing.T) {
	tests := []struct {
		name        string
		environment string
		want        []string
	}{
		{
			name:        "ローカル環境",
			environment: "local",
			want: []string{
				"http://localhost:3000",
				"http://127.0.0.1:3000",
			},
		},
		{
			name:        "本番環境",
			environment: "production",
			want: []string{
				"http://ty.inf.in.shizuoka.ac.jp",
				"https://ty.inf.in.shizuoka.ac.jp",
			},
		},
		{
			name:        "その他の環境",
			environment: "staging",
			want: []string{
				"http://ty.inf.in.shizuoka.ac.jp",
				"https://ty.inf.in.shizuoka.ac.jp",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getAllowedOrigins(tt.environment)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllowedOrigins() = %v, want %v", got, tt.want)
			}
		})
	}
}
