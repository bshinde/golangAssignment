package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	validator "golangAssignment/helper"
	"golangAssignment/model"
	"golangAssignment/service"
)

func TestHandleRequest(t *testing.T) {
	validate := validator.NewValidator()
	svc := service.NewService(validate)
	h := NewHandler(svc)

	// Valid payload
	payload := model.RequestPayload{
		Name:   "Test Name",
		Pan:    "ABCDE1234F",
		Mobile: "1234567890",
		Email:  "test@example.com",
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleRequest)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, rr.Code)
	}

	// Invalid JSON
	req, err = http.NewRequest("POST", "/validate", bytes.NewBuffer([]byte("{invalid-json}")))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %v for invalid JSON, got %v", http.StatusBadRequest, rr.Code)
	}

	// Missing Content-Type
	req, err = http.NewRequest("POST", "/validate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnsupportedMediaType {
		t.Errorf("Expected status code %v for missing Content-Type, got %v", http.StatusUnsupportedMediaType, rr.Code)
	}
}
