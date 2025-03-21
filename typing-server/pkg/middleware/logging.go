package middleware

import (
	"net/http"
	"time"

	"github.com/su-its/typing/typing-server/pkg/logger"
)

// loggingResponseWriter wraps http.ResponseWriter to capture the status code.
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware logs the details of each request and response.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.New()
		start := time.Now()
		log.Info("Incoming request",
			"method", r.Method,
			"url", r.URL.String())

		lrw := newLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)

		duration := time.Since(start)
		log.Info("Completed request",
			"status", lrw.statusCode,
			"duration", duration)
	})
}
