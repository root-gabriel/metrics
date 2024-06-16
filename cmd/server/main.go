package main

import (
    "log"
    "net/http"
    "github.com/root-gabriel/metrics/internal/handlers"
)

func main() {
    http.HandleFunc("/update/counter/", handlers.UpdateCounter)
    http.HandleFunc("/update/gauge/", handlers.UpdateGauge)
    http.HandleFunc("/", handlers.NotFound)

    log.Println("Server is starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

