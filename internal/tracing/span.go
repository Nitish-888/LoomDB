package tracing

import (
	"sync"
	"time"
)

type Span struct {
	mu        sync.Mutex        // Add this line
	TraceID   string            `json:"trace_id"`
	SpanID    string            `json:"span_id"`
	ParentID  string            `json:"parent_id"`
	Name      string            `json:"name"`
	StartTime time.Time         `json:"start_time"`
	EndTime   time.Time         `json:"end_time"`
	Tags      map[string]string `json:"tags"`
}

// Add a method to safely add tags
func (s *Span) SetTag(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Tags[key] = value
}