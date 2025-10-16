package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type PageData struct {
	BedtimesValue string
}

func main() {
	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Parse template once at startup
	tmpl := template.Must(template.ParseFiles("index.html"))

	// Serve the index.html file at the root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only serve index.html at root path
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Parse query parameters
		bedtimesParam := r.URL.Query().Get("bedtimes")

		// Convert comma-separated to newline-separated
		bedtimesValue := ""
		if bedtimesParam != "" {
			bedtimes := strings.Split(bedtimesParam, ",")
			var validBedtimes []string
			for _, bt := range bedtimes {
				trimmed := strings.TrimSpace(bt)
				if trimmed != "" {
					validBedtimes = append(validBedtimes, trimmed)
				}
			}
			bedtimesValue = strings.Join(validBedtimes, "\n")
		}

		// Render template with data
		data := PageData{
			BedtimesValue: bedtimesValue,
		}

		if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
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
