package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTrace(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		args args
		want bool // traceIDが設定されているかどうかを確認
	}{
		{
			name: "トレースIDが正しく設定される",
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// コンテキストからトレースIDを取得
					traceID := r.Context().Value(traceIDKey)
					if traceID == nil {
						t.Error("トレースIDがコンテキストに設定されていません")
					}
					if _, ok := traceID.(string); !ok {
						t.Error("トレースIDが文字列型ではありません")
					}
				}),
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// テストリクエストの作成
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rr := httptest.NewRecorder()

			// ミドルウェアの実行
			handler := Trace(tt.args.next)
			handler.ServeHTTP(rr, req)
		})
	}
}

// トレースIDが一意であることを確認するテスト
func TestTraceUniqueness(t *testing.T) {
	var firstTraceID string
	var secondTraceID string

	handler := Trace(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Context().Value(traceIDKey).(string)
		if firstTraceID == "" {
			firstTraceID = traceID
		} else {
			secondTraceID = traceID
		}
	}))

	// 1回目のリクエスト
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rr1 := httptest.NewRecorder()
	handler.ServeHTTP(rr1, req1)

	// 2回目のリクエスト
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	rr2 := httptest.NewRecorder()
	handler.ServeHTTP(rr2, req2)

	if firstTraceID == secondTraceID {
		t.Error("トレースIDが一意ではありません")
	}
}
