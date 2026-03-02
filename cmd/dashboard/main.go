package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Span struct {
	TraceID   string            `json:"trace_id"`
	Name      string            `json:"name"`
	StartTime time.Time         `json:"start_time"`
	EndTime   time.Time         `json:"end_time"`
	Tags      map[string]string `json:"tags"`
}

func main() {
	file, _ := os.ReadFile("traces.json")
	lines := strings.Split(string(file), "\n")

	var htmlContent strings.Builder
	htmlContent.WriteString("<html><head><title>LoomDB Pro Dashboard</title>")
	htmlContent.WriteString("<style>body{font-family:sans-serif; background:#1a1a1a; color:#eee; padding:40px;}")
	htmlContent.WriteString(".trace-card{background:#2d2d2d; border-radius:8px; padding:20px; margin-bottom:15px; border-left: 5px solid #4a90e2;}")
	htmlContent.WriteString(".bar-container{background:#444; width:100%; height:24px; border-radius:12px; margin:10px 0; overflow:hidden;}")
	htmlContent.WriteString(".bar{background:linear-gradient(90deg, #4a90e2, #63b3ed); height:100%; color:white; padding-left:10px; line-height:24px; font-size:12px; font-weight:bold; transition: width 0.5s;}")
	htmlContent.WriteString(".error-bar{background:linear-gradient(90deg, #e53e3e, #fc8181);}")
	htmlContent.WriteString("</style></head><body><h1>📊 LoomDB Performance Dashboard</h1>")

	for _, line := range lines {
		if line == "" { continue }
		var s Span
		json.Unmarshal([]byte(line), &s)

		// 1. CALCULATE DURATION
		duration := s.EndTime.Sub(s.StartTime)
		ms := duration.Milliseconds()
		
		// 2. SCALE THE WIDTH (e.g., 1ms = 2px, max 100%)
		width := ms * 2
		if width > 100 { width = 100 }
		if width < 5 { width = 10 } // Minimum visibility

		// 3. CHECK FOR ERRORS
		barClass := "bar"
		if s.Tags["error"] == "true" {
			barClass = "bar error-bar"
		}

		htmlContent.WriteString(fmt.Sprintf(`
			<div class="trace-card">
				<strong>%s</strong> <small>(%s)</small>
				<div class="bar-container">
					<div class="%s" style="width: %d%%;">%d ms</div>
				</div>
				<small>Method: %s | Target: %s</small>
			</div>`, s.Name, s.TraceID, barClass, width, ms, s.Tags["http.method"], s.Tags["db.system"]))
	}

	htmlContent.WriteString("</body></html>")
	os.WriteFile("dashboard.html", []byte(htmlContent.String()), 0644)
	fmt.Println("🚀 Day 7 Dashboard generated with real durations and error highlighting!")
}