package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"LRU-Cache/cache"
	"LRU-Cache/handlers"
)

func TestGetHandler(t *testing.T) {
	// Initialize cache
	c := cache.NewLRUCache(10)
	c.Set("testKey", "testValue")

	// Create a request
	req, err := http.NewRequest("GET", "/get?key=testKey", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler function using GetHandler
	handler := http.HandlerFunc(handlers.GetHandler(c))

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

}

func TestSetHandler(t *testing.T) {
	// Initialize cache
	c := cache.NewLRUCache(10)

	// Create a request body
	requestBody := map[string]interface{}{
		"key":   "testKey",
		"value": "testValue",
	}
	body, _ := json.Marshal(requestBody)

	// Create a request
	req, err := http.NewRequest("POST", "/set", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Set content type
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler function using SetHandler
	handler := http.HandlerFunc(handlers.SetHandler(c))

	// Serve the request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	// Check the cache value
	value, found := c.Get("testKey")
	if !found {
		t.Error("cache value not set")
	}
	if value != "testValue" {
		t.Errorf("cache value is incorrect: got %v want %v",
			value, "testValue")
	}
}
