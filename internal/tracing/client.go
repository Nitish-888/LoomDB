package tracing

import (
	"context"
	"net/http"
)

// InjectTrace adds the current TraceID from the context into the HTTP request headers
func InjectTrace(ctx context.Context, req *http.Request) {
	if span, ok := ctx.Value("span").(*Span); ok {
		req.Header.Set(TraceHeader, span.TraceID)
	}
}