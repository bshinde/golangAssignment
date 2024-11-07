package service

import (
	validator "golangAssignment/helper"
	"golangAssignment/model"
	"testing"
)

func setupService() *Service {
	validate := validator.NewValidator()
	return NewService(validate)
}

func TestValidatePayload_ValidRequest(t *testing.T) {
	service := setupService()

	// Create a valid request payload
	payload := model.RequestPayload{
		Name:   "Test_Name",
		Pan:    "ABCDE1234F",
		Mobile: "9876543210",
		Email:  "Test_Name@example.com",
	}

	// Validate the payload
	err := service.ValidatePayload(payload)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestValidatePayload_InvalidPAN(t *testing.T) {
	service := setupService()

	// Create an invalid PAN request payload
	payload := model.RequestPayload{
		Name:   "Test_Name",
		Pan:    "ABCDE12345", // Invalid PAN format
		Mobile: "9876543210",
		Email:  "Test_Name@example.com",
	}

	// Validate the payload
	err := service.ValidatePayload(payload)
	if err == nil {
		t.Errorf("Expected validation error for PAN, got no error")
	}
}

func TestValidatePayload_InvalidMobile(t *testing.T) {
	service := setupService()

	// Create an invalid mobile request payload
	payload := model.RequestPayload{
		Name:   "Test_Name",
		Pan:    "ABCDE1234F",
		Mobile: "98765432", // Invalid mobile (only 8 digits)
		Email:  "Test_Name@example.com",
	}

	// Validate the payload
	err := service.ValidatePayload(payload)
	if err == nil {
		t.Errorf("Expected validation error for Mobile, got no error")
	}
}
