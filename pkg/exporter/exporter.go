package exporter

import (
	"fmt"
	"github.com/Nitish_Thotakura/loomdb/internal/tracing"
)

// SpanExporter is an interface for sending spans to a backend.
type SpanExporter interface {
	Export(span *tracing.Span)
}

// ConsoleExporter simply prints spans to the terminal.
type ConsoleExporter struct{}

func (ce *ConsoleExporter) Export(span *tracing.Span) {
	duration := span.EndTime.Sub(span.StartTime)
	fmt.Printf("[LoomDB Export] Name: %s | TraceID: %s | SpanID: %s | Duration: %v\n", 
		span.Name, span.TraceID, span.SpanID, duration)
}