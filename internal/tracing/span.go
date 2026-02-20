package tracing

import (
	"time"
)

// Span represents a single unit of work in a distributed system.
type Span struct {
	TraceID   string            `json:"trace_id"`
	SpanID    string            `json:"span_id"`
	ParentID  string            `json:"parent_id"`
	Name      string            `json:"name"`
	StartTime time.Time         `json:"start_time"`
	EndTime   time.Time         `json:"end_time"`
	Tags      map[string]string `json:"tags"`
}