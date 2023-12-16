package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddCarHandler(t *testing.T) {
	router := setupRouter()

	car := Car{Make: "Toyota", Model: "Corolla"}
	payload, _ := json.Marshal(car)

	
	req, err := http.NewRequest("POST", "/addCar", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	
	req.Header.Set("Content-Type", "application/json")

	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addCarHandler)

	
	handler.ServeHTTP(rr, req)

	
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}
}

func TestGetCarHandler(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/getCar?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getCarHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}


