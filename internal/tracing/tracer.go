package tracing

import (
	"context"
	"fmt"
	"math/rand"
)

// Exporter defines the interface for sending spans to a destination
type Exporter interface {
	Export(span *Span)
}

// Global Variables shared across the tracing package
var (
	GlobalExporter Exporter
	SamplingRate   = 0.2 // 20% of traces are kept
)

// contextKey is a private type to prevent collisions in context.Context
type contextKey string
const spanKey contextKey = "span"

// generateID creates a simple random hex string for IDs (backup helper)
func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// SpanFromContext safely retrieves a span from a Go context
func SpanFromContext(ctx context.Context) *Span {
	if span, ok := ctx.Value(spanKey).(*Span); ok {
		return span
	}
	return nil
}