package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	// Create a simple HTTP server for testing
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer testServer.Close()

	// Create a context with timeout for testing
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Use the testing HTTP server URL and the testing context
	url := testServer.URL
	done := make(chan bool, 1)

	// Call the fetch function with the testing parameters
	go fetch(ctx, url, done)

	// Wait for the done channel or timeout
	select {
	case <-done:
		// The fetch function completed successfully
	case <-ctx.Done():
		t.Error("Test failed: Timeout reached")
	}
}
