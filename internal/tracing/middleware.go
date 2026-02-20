package tracing

import (
	"context"
	"net/http"
)

const TraceHeader = "X-Loom-Trace-ID"

// TraceMiddleware wraps an http.Handler to automatically start a span
func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Try to get the TraceID from the incoming request headers
		traceID := r.Header.Get(TraceHeader)
		
		ctx := r.Context()
		if traceID != "" {
			// If it exists, we inject a "dummy" span to act as the parent
			parent := &Span{TraceID: traceID, SpanID: "external-parent"}
			ctx = context.WithValue(ctx, "span", parent)
		}

		// 2. Start a new span for this specific HTTP request
		newCtx, span := StartSpan(ctx, r.Method+" "+r.URL.Path)
		defer span.End()

		// 3. Pass the new context (with the span) down to the actual handler
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}