package tracing

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

// NOTE: GlobalExporter, SamplingRate, and spanKey are now 
// defined in tracer.go. Do not redeclare them here!

type Event struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}

type Span struct {
	mu        sync.Mutex        `json:"-"` // Exclude mutex from JSON
	TraceID   string            `json:"trace_id"`
	SpanID    string            `json:"span_id"`
	ParentID  string            `json:"parent_id"`
	Name      string            `json:"name"`
	StartTime time.Time         `json:"start_time"`
	EndTime   time.Time         `json:"end_time"`
	Tags      map[string]string `json:"tags"`
	Events    []Event           `json:"events"`
	Sampled   bool              `json:"sampled"` // The decision flag
}

func StartSpan(ctx context.Context, name string) (context.Context, *Span) {
	spanID := uuid.New().String()
	parentSpan, ok := ctx.Value(spanKey).(*Span)

	var traceID string
	var parentID string
	var sampled bool

	if ok {
		// Child Span: Inherit TraceID and the Parent's sampling decision
		traceID = parentSpan.TraceID
		parentID = parentSpan.SpanID
		sampled = parentSpan.Sampled
	} else {
		// Root Span: Create new TraceID and make a NEW sampling decision
		traceID = uuid.New().String()
		sampled = rand.Float64() < SamplingRate
	}

	span := &Span{
		TraceID:   traceID,
		SpanID:    spanID,
		ParentID:  parentID,
		Name:      name,
		StartTime: time.Now(),
		Tags:      make(map[string]string),
		Sampled:   sampled,
	}

	return context.WithValue(ctx, spanKey, span), span
}

func (s *Span) End() {
	s.EndTime = time.Now()
	// Only send to exporter if the trace was sampled
	if s.Sampled && GlobalExporter != nil {
		GlobalExporter.Export(s)
	}
}

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

func (s *Span) RecordError(err error) {
	if err == nil {
		return
	}
	s.SetTag("error", "true")
	s.AddEvent("error: " + err.Error())
}