package main

import (
    "fmt"
    "net/http"
    "log"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go server!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "OK")
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/health", healthHandler)

    fmt.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}