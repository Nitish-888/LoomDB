# LoomDB - Distributed Tracing System (Go)

LoomDB is a lightweight, distributed tracing library built in Go. It allows developers to track requests as they flow through multiple microservices, providing visibility into latency and system behavior.

## ✨ Features
* **Distributed Context Propagation:** Seamlessly carries Trace IDs across HTTP boundaries.
* **Middleware Integration:** Automated span creation for HTTP servers.
* **Modular Exporters:** Support for Console and JSON File logging.
* **Thread-Safe:** Implements Mutex locking for concurrent span updates.

## 🚀 Getting Started
1. **Start the Server:** `go run cmd/server/main.go`
2. **Run the Client:** `go run cmd/client/main.go`
3. **View Traces:** `go run cmd/viewer/main.go`
