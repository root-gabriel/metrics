package main

import (
    "log"
    "net/http"
    "github.com/root-gabriel/metrics/internal/handlers"
)

func main() {
    http.HandleFunc("/update", handlers.UpdateMetric)
    http.HandleFunc("/value", handlers.GetMetric)

    log.Println("Server is starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
