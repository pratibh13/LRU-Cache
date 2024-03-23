package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"LRU-Cache/cache"
)

// GetHandler returns an HTTP handler function for retrieving a value from the cache
func GetHandler(cache *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// enableCors(&w)
		fmt.Println("Get method")
		key := r.URL.Query().Get("key")
		value, found := cache.Get(key)
		if !found {
			http.NotFound(w, r)
			return
		}
		// Set Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")
		// Encode the value to JSON and write it to the response writer
		json.NewEncoder(w).Encode(value)
	}
}

// SetHandler returns an HTTP handler function for setting a value in the cache
func SetHandler(cache *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// enableCors(&w)
		fmt.Println("Set method")
		var data struct {
			Key   string      `json:"key"`
			Value interface{} `json:"value"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		cache.Set(data.Key, data.Value)
		// Return the data in the response
		jsonResponse := map[string]interface{}{
			"key":   data.Key,
			"value": data.Value,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonResponse)
	}
}
