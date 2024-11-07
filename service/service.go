package service

import (
	validator "golangAssignment/helper"
	"golangAssignment/model"
)

// Service handles the business logic
type Service struct {
	Validator *validator.CustomValidator
}

// NewService creates a new instance of the service
func NewService(validator *validator.CustomValidator) *Service {
	return &Service{Validator: validator}
}

// ValidatePayload validates the incoming request payload
func (s *Service) ValidatePayload(payload model.RequestPayload) error {
	return s.Validator.Validate(payload)
}
