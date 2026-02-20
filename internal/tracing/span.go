package tracing

import (
	"sync"
	"time"
)

type Span struct {
	mu        sync.Mutex        // Protects the span from concurrent writes
	TraceID   string            `json:"trace_id"`
	SpanID    string            `json:"span_id"`
	ParentID  string            `json:"parent_id"`
	Name      string            `json:"name"`
	StartTime time.Time         `json:"start_time"`
	EndTime   time.Time         `json:"end_time"`
	Tags      map[string]string `json:"tags"`
}

// SetTag adds metadata to the span in a thread-safe way
func (s *Span) SetTag(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.Tags == nil {
		s.Tags = make(map[string]string)
	}
	s.Tags[key] = value
}