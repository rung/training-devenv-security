package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/", handler)

	// Cloud Run injects the PORT environment variable into the container.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// detected function is detected by linter
func detected() {
	_, err := os.Open("test")
	_, err = os.Open("test2")
	if err != nil {
		return
	}
	return
}
