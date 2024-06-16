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

// GetCounter возвращает значение счетчика
func GetCounter(metric string) (int, bool) {
    mutex.Lock()
    defer mutex.Unlock()
    value, found := counters[metric]
    return value, found
}

// GetGauge возвращает значение показателя
func GetGauge(metric string) (float64, bool) {
    mutex.Lock()
    defer mutex.Unlock()
    value, found := gauges[metric]
    return value, found
}

