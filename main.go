package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Server running on http://localhost:5000")
    http.ListenAndServe(":5000", nil)
}
