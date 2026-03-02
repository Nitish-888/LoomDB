package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Span struct {
	TraceID   string            `json:"trace_id"`
	Name      string            `json:"name"`
	StartTime string            `json:"start_time"`
	EndTime   string            `json:"end_time"`
	Tags      map[string]string `json:"tags"`
}

func main() {
	// 1. Read the JSON file
	file, _ := os.ReadFile("traces.json")
	lines := strings.Split(string(file), "\n")

	var htmlContent strings.Builder
	htmlContent.WriteString("<html><head><title>LoomDB Dashboard</title>")
	htmlContent.WriteString("<style>body{font-family:sans-serif; background:#f4f4f9; padding:20px;}")
	htmlContent.WriteString(".trace-card{background:white; border-radius:8px; padding:15px; margin-bottom:10px; box-shadow:0 2px 5px rgba(0,0,0,0.1);}")
	htmlContent.WriteString(".bar{background:#4a90e2; color:white; padding:5px; border-radius:4px; margin:5px 0; font-size:12px;}</style>")
	htmlContent.WriteString("</head><body><h1>🚀 LoomDB Trace Dashboard</h1>")

	// 2. Loop through spans and create "Bars"
	for _, line := range lines {
		if line == "" { continue }
		var s Span
		json.Unmarshal([]byte(line), &s)

		htmlContent.WriteString(fmt.Sprintf(`
			<div class="trace-card">
				<strong>Trace ID: %s</strong>
				<div class="bar" style="width: 80%%;">%s</div>
				<small>Tags: %v</small>
			</div>`, s.TraceID, s.Name, s.Tags))
	}

	htmlContent.WriteString("</body></html>")

	// 3. Save as HTML file
	os.WriteFile("dashboard.html", []byte(htmlContent.String()), 0644)
	fmt.Println("✨ Dashboard generated! Open 'dashboard.html' in your browser.")
}