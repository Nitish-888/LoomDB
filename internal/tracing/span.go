package tracing

import (
	"sync"
	"time"
)

type Event struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}

type Span struct {
	mu        sync.Mutex        // Protects the span from concurrent writes
	TraceID   string            `json:"trace_id"`
	SpanID    string            `json:"span_id"`
	ParentID  string            `json:"parent_id"`
	Name      string            `json:"name"`
	StartTime time.Time         `json:"start_time"`
	EndTime   time.Time         `json:"end_time"`
	Tags      map[string]string `json:"tags"`
	Events    []Event           `json:"events"`
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

func (s *Span) AddEvent(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Events = append(s.Events, Event{
		Name:      name,
		Timestamp: time.Now(),
	})
}

// RecordError is a helper to quickly tag a span as failed
func (s *Span) RecordError(err error) {
	if err == nil {
		return
	}
	s.SetTag("error", "true")
	s.AddEvent("error: " + err.Error())
}