package handlers

import (
    "net/http"
    "strconv"
    "strings"

    "github.com/root-gabriel/metrics/pkg/storage"
)

// UpdateCounter обрабатывает запросы на обновление счетчиков
func UpdateCounter(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/counter/"), "/")
    if len(parts) != 2 {
        http.Error(w, "Invalid request", http.StatusNotFound)
        return
    }
    metric, valueStr := parts[0], parts[1]
    value, err := strconv.Atoi(valueStr)
    if err != nil {
        http.Error(w, "Invalid value", http.StatusBadRequest)
        return
    }
    storage.UpdateCounter(metric, value)
    w.WriteHeader(http.StatusOK)
}

// UpdateGauge обрабатывает запросы на обновление показателей
func UpdateGauge(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/gauge/"), "/")
    if len(parts) != 2 {
        http.Error(w, "Invalid request", http.StatusNotFound)
        return
    }
    metric, valueStr := parts[0], parts[1]
    value, err := strconv.ParseFloat(valueStr, 64)
    if err != nil {
        http.Error(w, "Invalid value", http.StatusBadRequest)
        return
    }
    storage.UpdateGauge(metric, value)
    w.WriteHeader(http.StatusOK)
}

// UpdateUnknown обрабатывает запросы с неизвестными типами метрик
func UpdateUnknown(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/"), "/")
    if len(parts) != 3 {
        http.Error(w, "Invalid request", http.StatusNotFound)
        return
    }
    http.Error(w, "Invalid metric type", http.StatusNotImplemented)
}

// NotFound обрабатывает неизвестные маршруты
func NotFound(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Not found", http.StatusNotFound)
}

