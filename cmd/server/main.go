package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Nitish_Thotakura/loomdb/internal/tracing"
	"github.com/Nitish_Thotakura/loomdb/pkg/exporter"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Start the span
	_, childSpan := tracing.StartSpan(r.Context(), "database-query")

	// 2. Add breadcrumb Events to the span
	childSpan.AddEvent("connecting to postgresql...")
	time.Sleep(20 * time.Millisecond) // Simulate network latency

	// 3. Use Tags for static metadata
	childSpan.SetTag("db.system", "postgresql")
	childSpan.SetTag("http.method", r.Method)

	// 4. Simulate an error check
	// Let's pretend the query fails if the user sends a specific header
	if r.Header.Get("X-Simulate-Error") == "true" {
		err := errors.New("database connection timeout")
		childSpan.RecordError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		childSpan.End()
		return
	}

	childSpan.AddEvent("query-result-received")
	fmt.Println("Processing query with events and metadata...")
	
	childSpan.End()
	fmt.Fprintf(w, "LoomDB: Events and Metadata Recorded!")
}

func main() {
	// FIX: Use BatchExporter (saves every 5 spans OR every 5 seconds)
	tracing.GlobalExporter = exporter.NewBatchExporter("traces.json", 5)

	mux := http.NewServeMux()
	mux.Handle("/", tracing.TraceMiddleware(http.HandlerFunc(helloHandler)))

	fmt.Println("🚀 LoomDB Server with Batch Exporting on :8080")
	http.ListenAndServe(":8080", mux)
}