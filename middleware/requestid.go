package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

var RequestIDHeader string = "X-Request-Id"

const contextKeyRequestID string = "requestID"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			ctx = context.WithValue(ctx, contextKeyRequestID, uuid.New().String())
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
