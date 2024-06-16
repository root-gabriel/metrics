package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestUpdateCounter(t *testing.T) {
    req, err := http.NewRequest("POST", "/update/counter/testCounter/10", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(UpdateCounter)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}

func TestUpdateGauge(t *testing.T) {
    req, err := http.NewRequest("POST", "/update/gauge/testGauge/10.5", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(UpdateGauge)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}

