package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/Nitish-888/loomdb/internal/tracing" // Ensure this matches your module name
)

func main() {
	// 1. Define Search Flags
	traceID := flag.String("id", "", "`Search by specific Trace ID")
	showErrors := flag.Bool("errors", false, "Show only traces with errors")
	flag.Parse()

	file, err := os.ReadFile("traces.json")
	if err != nil {
		fmt.Println("❌ Could not read traces.json")
		return
	}

	lines := strings.Split(string(file), "\n")
	foundCount := 0

	fmt.Printf("🔎 Searching for: ID='%s', ErrorsOnly=%v\n", *traceID, *showErrors)
	fmt.Println(strings.Repeat("-", 50))

	for _, line := range lines {
		if line == "" { continue }

		var s tracing.Span
		json.Unmarshal([]byte(line), &s)

		// 2. APPLY FILTER LOGIC
		match := true

		// Filter by ID
		if *traceID != "" && s.TraceID != *traceID {
			match = false
		}

		// Filter by Errors
		if *showErrors && s.Tags["error"] != "true" {
			match = false
		}

		// 3. PRINT MATCHES
		if match {
			foundCount++
			status := "✅ OK"
			if s.Tags["error"] == "true" {
				status = "❌ ERROR"
			}
			fmt.Printf("[%s] %-20s | ID: %s | Time: %vms\n", 
				status, s.Name, s.TraceID, s.EndTime.Sub(s.StartTime).Milliseconds())
		}
	}

	fmt.Printf("\n✨ Found %d matching spans.\n", foundCount)
}