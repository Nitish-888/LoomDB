package tracing

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"
)

var GlobalExporter interface {
	Export(span *Span)
}

// StartSpan creates a new span and attaches it to the context.
func StartSpan(ctx context.Context, name string) (context.Context, *Span) {
	// Look for an existing span in the context to find the ParentID
	parent, _ := ctx.Value("span").(*Span)

	newSpan := &Span{
		Name:      name,
		StartTime: time.Now(),
		SpanID:    generateID(),
		Tags:      make(map[string]string),
	}

	if parent != nil {
		// Link this span to the existing trace
		newSpan.TraceID = parent.TraceID
		newSpan.ParentID = parent.SpanID
	} else {
		// This is the beginning of a brand new trace
		newSpan.TraceID = generateID()
	}

	// Return a new context containing the current span
	newCtx := context.WithValue(ctx, "span", newSpan)
	return newCtx, newSpan
}

// generateID creates a simple random hex string for IDs
func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// End sets the end time for the span. 
// In a real system, this is where you'd send the data to a database.


func (s *Span) End() {
	s.EndTime = time.Now()
	if GlobalExporter != nil {
		GlobalExporter.Export(s)
	}
}