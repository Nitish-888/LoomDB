package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	requests := 100 // Let's send 100 requests at once
	url := "http://localhost:8080/"

	fmt.Printf("🔥 Stress Test: Sending %d requests to LoomDB...\n", requests)

	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("❌ Request %d failed: %v\n", id, err)
				return
			}
			resp.Body.Close()
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ Stress Test Complete! Check your server logs and traces.json.")
}