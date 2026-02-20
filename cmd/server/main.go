package main

import (
	"fmt"
	"net/http"

	"github.com/Nitish_Thotakura/loomdb/internal/tracing"
	"github.com/Nitish_Thotakura/loomdb/pkg/exporter"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Start the span
	_, childSpan := tracing.StartSpan(r.Context(), "database-query")
	
	// 2. Use our new SetTag method!
	childSpan.SetTag("db.system", "postgresql")
	childSpan.SetTag("http.method", r.Method)
	
	fmt.Println("Processing query with metadata...")
	childSpan.End()

	fmt.Fprintf(w, "LoomDB: Metadata Recorded!")
}

func main() {
	// FIX 1: Use FileExporter so the viewer has a file to read
	tracing.GlobalExporter = exporter.NewFileExporter("traces.json")

	mux := http.NewServeMux()

	// FIX 2: Wrap the handler with Middleware
	mux.Handle("/", tracing.TraceMiddleware(http.HandlerFunc(helloHandler)))

	fmt.Println("🚀 LoomDB Server listening on :8080")
	fmt.Println("📝 Saving traces to traces.json")

	// This starts the server and waits for requests
	http.ListenAndServe(":8080", mux)
}
