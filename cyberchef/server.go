package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Get directory to serve
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	// Get port
	port := "8000"
	if len(os.Args) > 2 {
		port = os.Args[2]
	}

	// Get absolute path
	absDir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}

	// Check if directory exists
	if _, err := os.Stat(absDir); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", absDir)
	}

	// Create file server with CORS support
	fs := http.FileServer(http.Dir(absDir))
	handler := corsMiddleware(fs)

	// Set up routes
	http.Handle("/", handler)

	fmt.Printf("🚀 Serving CyberChef from: %s\n", absDir)
	fmt.Printf("🌐 Access at: http://localhost:%s\n", port)
	fmt.Printf("⏹️  Press Ctrl+C to stop\n\n")

	// Start server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
