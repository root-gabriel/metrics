package main

import (
    "log"
    "net/http"
    "time"
    "github.com/root-gabriel/metrics/internal/handlers"
)

func sendTestMetrics() {
    client := &http.Client{}
    for {
        req, _ := http.NewRequest("POST", "http://localhost:8080/update/counter/testCounter/1", nil)
        _, err := client.Do(req)
        if err != nil {
            log.Println("Error sending metric:", err)
        }
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go sendTestMetrics() // Запускаем генерацию метрик в отдельной горутине

    http.HandleFunc("/update/counter/", handlers.UpdateCounter)
    http.HandleFunc("/update/gauge/", handlers.UpdateGauge)
    http.HandleFunc("/value/counter/", handlers.GetCounterValue)
    http.HandleFunc("/value/gauge/", handlers.GetGaugeValue)
    http.HandleFunc("/update/", handlers.UpdateUnknown)
    http.HandleFunc("/", handlers.NotFound)

    log.Println("Agent is starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

