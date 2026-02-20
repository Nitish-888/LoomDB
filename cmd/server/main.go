package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"github.com/Nitish_Thotakura/loomdb/internal/tracing" // Use your actual module name
	"github.com/Nitish_Thotakura/loomdb/pkg/exporter"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Any logic inside here is now automatically traced!
	fmt.Fprintf(w, "Hello from LoomDB Traced Server!")
}

func main() {
	fmt.Println("Starting LoomDB Test Server...")

	// 1. Start a "Root" Span (the beginning of a request)
	ctx, rootSpan := tracing.StartSpan(context.Background(), "main-request")
	time.Sleep(100 * time.Millisecond) // Simulate work

	// 2. Start a "Child" Span (simulating a database call)
	_, childSpan := tracing.StartSpan(ctx, "database-query")
	time.Sleep(50 * time.Millisecond) // Simulate DB latency
	childSpan.End()

	tracing.GlobalExporter = &exporter.ConsoleExporter{}

	mux := http.NewServeMux()
	mux.Handle("/", tracing.TraceMiddleware(http.HandlerFunc(helloHandler)))

	fmt.Println("LoomDB Server listening on :8080...")
	http.ListenAndServe(":8080", mux)

	

	// 3. End the Root Span
	rootSpan.End()
}