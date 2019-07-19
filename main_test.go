package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandle(t *testing.T) {
	endpoints := []string{"/", "/health", "/ready"}

	for _, e := range endpoints {
		req, err := http.NewRequest("GET", e, nil)
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(EchoHandler)

		handler.ServeHTTP(rec, req)

		if status := rec.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `{"OK": true}`
		if rec.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rec.Body.String(), expected)
		}
	}
}
