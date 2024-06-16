package handlers

import (
    "fmt"
    "net/http"
    "github.com/root-gabriel/metrics/pkg/storage"
)

func UpdateMetric(w http.ResponseWriter, r *http.Request) {
    metric := r.URL.Query().Get("metric")
    value := r.URL.Query().Get("value")
    storage.UpdateMetric(metric, value)
    fmt.Fprintf(w, "Updated metric %s to %s", metric, value)
}

func GetMetric(w http.ResponseWriter, r *http.Request) {
    metric := r.URL.Query().Get("metric")
    value, err := storage.GetMetric(metric)
    if err != nil {
        http.Error(w, "Metric not found", 404)
        return
    }
    fmt.Fprintf(w, "Value of metric %s is %s", metric, value)
}
