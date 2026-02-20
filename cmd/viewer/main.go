package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"github.com/Nitish_Thotakura/loomdb/internal/tracing"
)

func main() {
	file, err := os.Open("traces.json")
	if err != nil {
		fmt.Println("No traces found. Run the server and client first!")
		return
	}
	defer file.Close()

	fmt.Println("--- LoomDB Trace History ---")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var s tracing.Span
		if err := json.Unmarshal(scanner.Bytes(), &s); err == nil {
			duration := s.EndTime.Sub(s.StartTime)
			fmt.Printf("[%s] %-20s | ID: %s | Duration: %v\n", 
				s.TraceID[:8], s.Name, s.SpanID, duration)
		}
	}
}