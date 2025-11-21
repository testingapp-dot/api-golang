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
    fmt.Println("Server running at http://localhost:5000/")
    if err := http.ListenAndServe(":5000", nil); err != nil {
        fmt.Println("error starting server:", err)
    }
}