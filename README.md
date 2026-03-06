# 🧶 LoomDB

**LoomDB** is a lightweight, distributed tracing engine written in Go. It provides deep visibility into microservices without the overhead of massive external dependencies. Designed for high-concurrency environments, LoomDB implements professional-grade patterns for data collection, performance optimization, and search.

---

## 🛠 Technical Architecture

LoomDB is built on three pillars of high-performance observability:

### 1. Context-Aware Propagation
LoomDB leverages the Go `context` package to pass Trace IDs across function boundaries and network hops. This ensures that every child span (DB queries, external API calls) is perfectly linked to its parent request.

### 2. Probabilistic Root Sampler
To prevent storage bloat, LoomDB implements a **Probabilistic Sampling** algorithm. 
- **The Logic:** Sampling decisions are made at the *Root Span* and propagated down the entire tree.
- **Efficiency:** Captures a representative 10-20% of traffic, ensuring high-volume systems remain performant while still providing statistical visibility.

### 3. Non-Blocking Batch Exporter
LoomDB utilizes a **Worker-Pool Pattern** to flush traces to disk. 
- **The Benefit:** Instead of writing to a file on every request (which would slow down the server), spans are collected in a thread-safe buffer and flushed in batches, minimizing Disk I/O contention.

---

## 🚀 Key Features

* **Thread-Safe Spans:** Uses `sync.Mutex` to handle concurrent tag updates and event logging.
* **Built-in CLI Search:** A dedicated filtering engine to find traces by `TraceID` or filter for only `Error` states.
* **Interactive Dashboard:** A dynamic HTML generator that turns raw JSON traces into a visual waterfall timeline.
* **Error Instrumentation:** Automatic capture of Go errors into span events for rapid debugging.

---

## 📂 Project Structure

```text
.
├── cmd/
│   ├── server/    # Main API Server with tracing middleware
│   ├── client/    # Traffic generator for testing propagation
│   ├── dashboard/ # Generates the HTML visualization
│   ├── search/    # CLI tool for filtering JSON traces
│   └── stress/    # Concurrency stress-tester
├── internal/
│   └── tracing/   # Core logic (Span, Tracer, Sampler, Batcher)
└── traces.json    # Persistent storage of captured spans
