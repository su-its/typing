package logger

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/su-its/typing/typing-server/api/config"
)

type TraceIDLogHandler struct {
	slog.Handler
}

func NewTraceIDLogHandler(h slog.Handler) *TraceIDLogHandler {
	return &TraceIDLogHandler{Handler: h}
}

func (h *TraceIDLogHandler) Handle(ctx context.Context, r slog.Record) error {
	if v, ok := ctx.Value("trace_id").(string); ok {
		r.AddAttrs(slog.String("trace_id", v))
	}
	return h.Handler.Handle(ctx, r)
}

func New(config *config.Config) *slog.Logger {
	logLevel := new(slog.LevelVar)
	if config.Environment == "production" {
		logLevel.Set(slog.LevelInfo)
	} else {
		logLevel.Set(slog.LevelDebug)
	}

	// JSON形式でログを出力するハンドラーを作成
	opts := &slog.HandlerOptions{
		Level:     logLevel,
		AddSource: true, // ソースコードの位置情報を追加
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// タイムスタンプをISO8601形式に変換
			if a.Key == slog.TimeKey {
				return slog.Attr{
					Key:   a.Key,
					Value: slog.StringValue(a.Value.Time().Format(time.RFC3339)),
				}
			}
			return a
		},
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	// TraceIDを追加するカスタムハンドラーでラップ
	traceHandler := NewTraceIDLogHandler(handler)

	// グローバルロガーを設定
	logger := slog.New(traceHandler)
	slog.SetDefault(logger)

	return logger
}
