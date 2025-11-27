package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function for the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!") // Write "Hello, World!" to the response
	})

	// Start the HTTP server and listen on port 8080
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil) // nil uses the default ServeMux
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}