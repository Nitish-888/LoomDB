package exporter

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
	"github.com/Nitish_Thotakura/loomdb/internal/tracing"
)

type BatchExporter struct {
	FilePath    string
	batchSize   int
	buffer      []*tracing.Span
	mu          sync.Mutex
	spanChan    chan *tracing.Span
}

func NewBatchExporter(path string, size int) *BatchExporter {
	be := &BatchExporter{
		FilePath:  path,
		batchSize: size,
		spanChan:  make(chan *tracing.Span, 100),
	}
	go be.start() // Start the background worker
	return be
}

func (be *BatchExporter) Export(span *tracing.Span) {
	be.spanChan <- span // Send span to the "waiting room"
}

func (be *BatchExporter) start() {
	ticker := time.NewTicker(5 * time.Second) // Flush every 5 seconds
	for {
		select {
		case span := <-be.spanChan:
			be.mu.Lock()
			be.buffer = append(be.buffer, span)
			if len(be.buffer) >= be.batchSize {
				be.flush()
			}
			be.mu.Unlock()
		case <-ticker.C:
			be.mu.Lock()
			be.flush()
			be.mu.Unlock()
		}
	}
}

func (be *BatchExporter) flush() {
	if len(be.buffer) == 0 {
		return
	}
	f, _ := os.OpenFile(be.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	for _, s := range be.buffer {
		data, _ := json.Marshal(s)
		f.Write(append(data, '\n'))
	}
	fmt.Printf("🚀 Batched %d traces to %s\n", len(be.buffer), be.FilePath)
	be.buffer = nil // Clear the waiting room
}