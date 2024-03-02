package system

import (
	"log/slog"
	"net/http"
)

// HealthCheck はヘルスチェックのためのハンドラー関数です。
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("API is running"))
	if err != nil {
		// エラーログを記録し、処理を終了します。
		// 実際には、この時点でレスポンスヘッダーやボディがクライアントに送信されている可能性が高いため、
		// http.Errorを呼び出すことは推奨されません。
		// 代わりに、ログに記録するなどのサーバー側での対応が適切です。
		slog.Error("failed to write response: %v", err)
	}
}

