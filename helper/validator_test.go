package validator

import (
	"golangAssignment/model"
	"testing"
)

func TestPANValidator_Valid(t *testing.T) {
	// Initialize the custom validator
	validator := NewValidator()

	// Valid PAN number, along with other required fields
	payload := model.RequestPayload{
		Name:   "Test Name",
		Pan:    "ABCDE1234F", // Valid PAN
		Mobile: "9876543210",
		Email:  "testname@example.com",
	}

	// Validate the payload
	err := validator.Validate(payload)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
func TestPANValidator_Invalid(t *testing.T) {
	// Initialize the custom validator
	validator := NewValidator()

	// Invalid PAN number
	payload := model.RequestPayload{
		Pan: "ABCDE12345", // Invalid PAN (should have 1 letter at the end)
	}

	// Validate the payload
	err := validator.Validate(payload)
	if err == nil {
		t.Errorf("Expected validation error for PAN, but got no error")
	}
}

func TestPANValidator_InvalidFormat(t *testing.T) {
	// Initialize the custom validator
	validator := NewValidator()

	// Invalid PAN number with non-alphabetic characters
	payload := model.RequestPayload{
		Pan: "ABCD@1234F", // Invalid format
	}

	// Validate the payload
	err := validator.Validate(payload)
	if err == nil {
		t.Errorf("Expected validation error for PAN, but got no error")
	}
}

func TestMobileValidator_Valid(t *testing.T) {
	// Initialize the custom validator
	validator := NewValidator()

	// Valid mobile number (10 digits)
	payload := model.RequestPayload{
		Name:   "Test Name",
		Pan:    "ABCDE1234F",
		Mobile: "9876543210", // Valid mobile
		Email:  "test@example.com",
	}

	// Validate the payload
	err := validator.Validate(payload)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestMobileValidator_Invalid(t *testing.T) {
	// Initialize the custom validator
	validator := NewValidator()

	// Invalid mobile number (less than 10 digits)
	payload := model.RequestPayload{
		Mobile: "98765432", // Invalid mobile (only 8 digits)
	}

	// Validate the payload
	err := validator.Validate(payload)
	if err == nil {
		t.Errorf("Expected validation error for Mobile, but got no error")
	}
}

func TestEmailValidator_Valid(t *testing.T) {
	// Initialize the custom validator
	validator := NewValidator()

	// Valid email
	payload := model.RequestPayload{
		Name:   "Test Name",
		Pan:    "ABCDE1234F",
		Mobile: "9876543210",
		Email:  "test@example.com",
	}

	// Validate the payload
	err := validator.Validate(payload)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestEmailValidator_Invalid(t *testing.T) {
	// Initialize the custom validator
	validator := NewValidator()

	// Invalid email format
	payload := model.RequestPayload{
		Email: "test@com", // Invalid email format
	}

	// Validate the payload
	err := validator.Validate(payload)
	if err == nil {
		t.Errorf("Expected validation error for Email, but got no error")
	}
}

func TestCompleteRequestPayload_Valid(t *testing.T) {
	// Initialize the custom validator
	validator := NewValidator()

	// Create a valid payload
	payload := model.RequestPayload{
		Name:   "Test Name",
		Pan:    "ABCDE1234F",
		Mobile: "9876543210",
		Email:  "test@example.com",
	}

	// Validate the payload
	err := validator.Validate(payload)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCompleteRequestPayload_Invalid(t *testing.T) {
	// Initialize the custom validator
	validator := NewValidator()

	// Create an invalid payload (invalid PAN)
	payload := model.RequestPayload{
		Name:   "Test Name",
		Pan:    "ABCDE12345", // Invalid PAN format
		Mobile: "9876543210",
		Email:  "test@example.com",
	}

	// Validate the payload
	err := validator.Validate(payload)
	if err == nil {
		t.Errorf("Expected validation error, but got no error")
	}
}
