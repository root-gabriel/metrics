package storage

import (
    "sync"
    "fmt"
)

var metrics = make(map[string]string)
var mutex = &sync.Mutex{}

func UpdateMetric(metric, value string) {
    mutex.Lock()
    defer mutex.Unlock()
    metrics[metric] = value
}

func GetMetric(metric string) (string, error) {
    mutex.Lock()
    defer mutex.Unlock()
    value, exists := metrics[metric]
    if !exists {
        return "", fmt.Errorf("metric not found")
    }
    return value, nil
}
