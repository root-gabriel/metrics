package handlers

import (
    "net/http"
    "strconv"
    "strings"

    "github.com/root-gabriel/metrics/pkg/storage"
)

// UpdateCounter обрабатывает запросы на обновление счетчиков
func UpdateCounter(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/counter/"), "/")
    if len(parts) != 2 {
        http.Error(w, "Invalid request", 400)
        return
    }
    metric, valueStr := parts[0], parts[1]
    value, err := strconv.Atoi(valueStr)
    if err != nil {
        http.Error(w, "Invalid value", 400)
        return
    }
    storage.UpdateCounter(metric, value)
    w.WriteHeader(200)
}

// UpdateGauge обрабатывает запросы на обновление показателей
func UpdateGauge(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/gauge/"), "/")
    if len(parts) != 2 {
        http.Error(w, "Invalid request", 400)
        return
    }
    metric, valueStr := parts[0], parts[1]
    value, err := strconv.ParseFloat(valueStr, 64)
    if err != nil {
        http.Error(w, "Invalid value", 400)
        return
    }
    storage.UpdateGauge(metric, value)
    w.WriteHeader(200)
}

// NotFound обрабатывает неизвестные маршруты
func NotFound(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Not found", 404)
}

