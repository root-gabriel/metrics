package storage

import "sync"

var (
    counters = make(map[string]int)
    gauges   = make(map[string]float64)
    mutex    = &sync.Mutex{}
)

// UpdateCounter обновляет значение счетчика
func UpdateCounter(metric string, value int) {
    mutex.Lock()
    defer mutex.Unlock()
    counters[metric] += value
}

// UpdateGauge обновляет значение показателя
func UpdateGauge(metric string, value float64) {
    mutex.Lock()
    defer mutex.Unlock()
    gauges[metric] = value
}

