package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Span struct {
	TraceID   string            `json:"trace_id"`
	SpanID    string            `json:"span_id"`
	ParentID  string            `json:"parent_id"`
	Name      string            `json:"name"`
	StartTime string            `json:"start_time"`
	EndTime   string            `json:"end_time"`
	Tags      map[string]string `json:"tags"`
}

func main() {
	file, _ := os.ReadFile("traces.json")
	lines := strings.Split(string(file), "\n")

	// 1. Group spans by TraceID
	traces := make(map[string][]Span)
	for _, line := range lines {
		if line == "" { continue }
		var s Span
		json.Unmarshal([]byte(line), &s)
		traces[s.TraceID] = append(traces[s.TraceID], s)
	}

	// 2. Print each Trace as a "Waterfall"
	for tid, spans := range traces {
		fmt.Printf("\n--- Trace: %s ---\n", tid)
		
		// Sort by start time
		sort.Slice(spans, func(i, j int) bool {
			return spans[i].StartTime < spans[j].StartTime
		})

		for _, s := range spans {
			indent := ""
			if s.ParentID != "" {
				indent = "  └── " // Indent children
			}
			fmt.Printf("%s%s (%s)\n", indent, s.Name, s.SpanID)
			
			// Print tags if they exist
			for k, v := range s.Tags {
				fmt.Printf("%s    [%s: %s]\n", indent, k, v)
			}
		}
	}
}