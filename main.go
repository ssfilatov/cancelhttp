package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func fetch(ctx context.Context, url string, done chan bool) {
	defer func() {
		done <- true
	}()

	client := http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request for %s: %v\n", url, err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making GET request to %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("URL: %s, Status Code: %d\n", url, resp.StatusCode)
}

func main() {
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://facebook.com",
		"https://twitter.com",
		"https://github.com",
	}

	done := make(chan bool, len(urls))

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	for _, url := range urls {
		go fetch(ctx, url, done)
	}

	// Wait for all goroutines to finish or for timeout
	for {
		select {
		case <-done:
			// One request completed
		case <-ctx.Done():
			fmt.Println("Timeout reached. Cancelling remaining requests.")
			return
		}
	}
}
