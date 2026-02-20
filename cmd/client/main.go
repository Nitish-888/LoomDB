package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"github.com/Nitish_Thotakura/loomdb/internal/tracing"
	"github.com/Nitish_Thotakura/loomdb/pkg/exporter"
)

func main() {
	tracing.GlobalExporter = &exporter.ConsoleExporter{}

	// 1. Start a span for the client's work
	ctx, span := tracing.StartSpan(context.Background(), "client-operation")
	defer span.End()

	// 2. Prepare the HTTP request to the server
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/", nil)
	
	// 3. Inject the trace!
	tracing.InjectTrace(ctx, req)

	fmt.Println("Client sending request with TraceID:", span.TraceID)
	
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error calling server:", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Println("Server responded with status:", resp.Status)
}