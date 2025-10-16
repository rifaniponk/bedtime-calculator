package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Serve the index.html file at the root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only serve index.html at root path
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		http.ServeFile(w, r, "index.html")
	})

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	addr := fmt.Sprintf(":%s", port)
	log.Printf("🛏️  Bedtime Calculator server starting on http://localhost%s", addr)
	log.Printf("Press Ctrl+C to stop the server")
	
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

