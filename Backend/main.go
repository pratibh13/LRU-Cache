package main

import (
	"LRU-Cache/cache"
	"LRU-Cache/handlers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	size := 1024
	c := cache.NewLRUCache(size)

	// Initialize CORS middleware with custom options
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Register handlers directly with http.HandleFunc
	http.HandleFunc("/get", handlers.GetHandler(c))
	http.HandleFunc("/set", handlers.SetHandler(c))

	// Start server with CORS handler
	log.Fatal(http.ListenAndServe(":8080", corsHandler.Handler(http.DefaultServeMux)))

}
