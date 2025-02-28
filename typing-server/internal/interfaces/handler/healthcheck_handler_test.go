package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHealthCheckHandler(t *testing.T) {
	tests := []struct {
		name string
		want *HealthCheckHandler
	}{
		{
			name: "正常にハンドラーが生成されること",
			want: &HealthCheckHandler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHealthCheckHandler()
			if got == nil {
				t.Error("NewHealthCheckHandler() returned nil")
			}
		})
	}
}

func TestHealthCheckHandler_LivenessProbe(t *testing.T) {
	tests := []struct {
		name       string
		wantStatus int
	}{
		{
			name:       "正常にヘルスチェックが返されること",
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// テスト用のリクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			w := httptest.NewRecorder()

			// ハンドラーを実行
			h := NewHealthCheckHandler()
			h.LivenessProbe(w, req)

			// レスポンスを検証
			if status := w.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}
		})
	}
}
