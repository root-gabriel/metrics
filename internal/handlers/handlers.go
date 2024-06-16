package handlers

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
    "github.com/root-gabriel/metrics/pkg/storage"
)

// UpdateCounter обрабатывает запросы на обновление счетчиков
func UpdateCounter(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        log.Println("UpdateCounter: Method not allowed")
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/counter/"), "/")
    if len(parts) != 2 {
        log.Println("UpdateCounter: Invalid request")
        http.Error(w, "Invalid request", http.StatusNotFound)
        return
    }
    metric, valueStr := parts[0], parts[1]
    value, err := strconv.Atoi(valueStr)
    if err != nil {
        log.Println("UpdateCounter: Invalid value")
        http.Error(w, "Invalid value", http.StatusBadRequest)
        return
    }
    log.Printf("Updating counter %s by %d\n", metric, value)
    storage.UpdateCounter(metric, value)
    w.WriteHeader(http.StatusOK)
}

// UpdateGauge обрабатывает запросы на обновление показателей
func UpdateGauge(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        log.Println("UpdateGauge: Method not allowed")
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/gauge/"), "/")
    if len(parts) != 2 {
        log.Println("UpdateGauge: Invalid request")
        http.Error(w, "Invalid request", http.StatusNotFound)
        return
    }
    metric, valueStr := parts[0], parts[1]
    value, err := strconv.ParseFloat(valueStr, 64)
    if err != nil {
        log.Println("UpdateGauge: Invalid value")
        http.Error(w, "Invalid value", http.StatusBadRequest)
        return
    }
    log.Printf("Updating gauge %s to %f\n", metric, value)
    storage.UpdateGauge(metric, value)
    w.WriteHeader(http.StatusOK)
}

// GetCounterValue возвращает значение счетчика
func GetCounterValue(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        log.Println("GetCounterValue: Method not allowed")
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    metric := strings.TrimPrefix(r.URL.Path, "/value/counter/")
    value, found := storage.GetCounter(metric)
    if !found {
        log.Println("GetCounterValue: Metric not found")
        http.Error(w, "Metric not found", http.StatusNotFound)
        return
    }
    log.Printf("Returning counter value for %s: %d\n", metric, value)
    fmt.Fprintf(w, "%d", value)
}

// GetGaugeValue возвращает значение показателя
func GetGaugeValue(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        log.Println("GetGaugeValue: Method not allowed")
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    metric := strings.TrimPrefix(r.URL.Path, "/value/gauge/")
    value, found := storage.GetGauge(metric)
    if !found {
        log.Println("GetGaugeValue: Metric not found")
        http.Error(w, "Metric not found", http.StatusNotFound)
        return
    }
    log.Printf("Returning gauge value for %s: %f\n", metric, value)
    fmt.Fprintf(w, "%f", value)
}

// UpdateUnknown обрабатывает запросы с неизвестными типами метрик
func UpdateUnknown(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        log.Println("UpdateUnknown: Method not allowed")
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/"), "/")
    if len(parts) != 3 {
        log.Println("UpdateUnknown: Invalid request")
        http.Error(w, "Invalid request", http.StatusNotFound)
        return
    }
    log.Println("UpdateUnknown: Invalid metric type")
    http.Error(w, "Invalid metric type", http.StatusNotImplemented)
}

// NotFound обрабатывает неизвестные маршруты
func NotFound(w http.ResponseWriter, r *http.Request) {
    log.Println("NotFound: Unknown route")
    http.Error(w, "Not found", http.StatusNotFound)
}

