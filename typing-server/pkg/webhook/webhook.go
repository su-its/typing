package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type WebhookNotifier struct {
	log *slog.Logger
}

// NewWebhookNotifier は WebhookNotifier のインスタンスを生成する
func NewWebhookNotifier(log *slog.Logger) *WebhookNotifier {
	return &WebhookNotifier{log: log}
}

func IsBrowserRequest(userAgent string) bool { // Keep the original one for now
	if userAgent == "" {
		return false // User-Agent が空の場合は非ブラウザとみなす
	}
	// 一般的なブラウザの User-Agent に含まれるキーワード
	browserKeywords := []string{"Mozilla", "Chrome", "Safari", "Firefox", "Edge"}
	for _, keyword := range browserKeywords {
		if strings.Contains(userAgent, keyword) {
			return true
		}
	}
	return false
}

// SendNonBrowserScoreNotification は非ブラウザからのスコア登録アクセス情報をWebhookに送信する
func (n *WebhookNotifier) SendNonBrowserScoreNotification(originalReq *http.Request, scoreReqBody interface{}) {
	webhookURL := "https://discord.com/api/webhooks/1359681161907015716/cjvtlA78Mc947HkIH5k7GwbszhAMA9n50ZIFMs7bi9UTCKtvM1rbEwaQ-bsh8VzDvCS_"

	// User-Agentを取得
	userAgent := originalReq.Header.Get("User-Agent")
	if userAgent == "" {
		userAgent = "N/A"
	}

	// リクエストボディをJSON文字列に整形
	reqBodyBytes, _ := json.MarshalIndent(scoreReqBody, "", "  ")
	reqBodyStr := string(reqBodyBytes)
	if len(reqBodyStr) > 1000 {
		reqBodyStr = reqBodyStr[:1000] + "..."
	}

	now := time.Now().UTC() // Need to re-add "time" import if removed

	// Discord Embeds形式のペイロードを作成
	discordPayload := map[string]interface{}{
		"username": "Typing Server Alert",
		"embeds": []map[string]interface{}{
			{
				"title":       "🚨 Non-Browser Score Submission Detected",
				"description": "A score submission possibly from a non-browser client was detected.",
				"color":       15158332,
				"fields": []map[string]interface{}{
					{"name": "Timestamp (UTC)", "value": now.Format(time.RFC3339), "inline": true},
					{"name": "Remote Address", "value": originalReq.RemoteAddr, "inline": true},
					{"name": "User-Agent", "value": userAgent, "inline": false},
					{"name": "Request Body", "value": "```json\n" + reqBodyStr + "\n```", "inline": false},
				},
				"footer":    map[string]string{"text": "Typing Server"},
				"timestamp": now.Format(time.RFC3339),
			},
		},
	}

	payloadBytes, err := json.Marshal(discordPayload)
	if err != nil {
		n.log.Error("Failed to marshal discord webhook payload", "error", err)
		return
	}

	client := &http.Client{Timeout: 10 * time.Second}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", webhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		n.log.Error("Failed to create webhook request", "url", webhookURL, "error", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		n.log.Error("Failed to send webhook notification", "url", webhookURL, "error", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBodyBytes, _ := io.ReadAll(resp.Body)
		n.log.Warn("Webhook notification sent but received non-2xx status",
			"url", webhookURL, "status_code", resp.StatusCode, "response_body", string(respBodyBytes))
	} else {
		n.log.Info("Webhook notification sent successfully", "url", webhookURL)
	}
}
