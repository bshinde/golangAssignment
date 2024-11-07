package validator

import (
	"golangAssignment/model"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// CustomValidator extends the standard validator to include custom rules
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator creates a new instance of the custom validator
func NewValidator() *CustomValidator {
	validate := validator.New()
	validate.RegisterValidation("pan", panValidator)
	return &CustomValidator{validator: validate}
}

// panValidator checks if the PAN number is in the correct format
func panValidator(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[A-Z]{5}[0-9]{4}[A-Z]{1}$`)
	return re.MatchString(fl.Field().String())
}

// Validate validates the request payload using the custom validator
func (v *CustomValidator) Validate(payload model.RequestPayload) error {
	return v.validator.Struct(payload)
}
