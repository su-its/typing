package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const traceIDKey contextKey = "trace_id"

func Trace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := uuid.New().String()
		ctx := context.WithValue(r.Context(), traceIDKey, traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
