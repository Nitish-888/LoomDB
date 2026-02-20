package exporter

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"github.com/Nitish_Thotakura/loomdb/internal/tracing"
)

// SpanExporter is an interface for sending spans to a backend.
type SpanExporter interface {
	Export(span *tracing.Span)
}

// FileExporter saves spans to a local JSON file.
type FileExporter struct {
	FilePath string
	mu       sync.Mutex
}

// NewFileExporter creates a new instance of FileExporter.
// This matches the call in your main.go!
func NewFileExporter(path string) *FileExporter {
	return &FileExporter{FilePath: path}
}

func (fe *FileExporter) Export(span *tracing.Span) {
	fe.mu.Lock()
	defer fe.mu.Unlock()

	// Open file in Append mode, create if it doesn't exist.
	f, err := os.OpenFile(fe.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening trace file: %v\n", err)
		return
	}
	defer f.Close()

	// Convert span to JSON and write a new line.
	data, _ := json.Marshal(span)
	f.Write(append(data, '\n'))
}