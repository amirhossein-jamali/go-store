package model

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Product represents a product in the database.
type Product struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255);not null" validate:"min=3,max=50"`
}

// Validate validates the Product struct and provides custom error messages.
func (p *Product) Validate() error {
	// Initialize the validator
	validate := validator.New()

	// Validate the struct
	err := validate.Struct(p)
	if err == nil {
		return nil // Validation passed
	}

	// Handle validation errors
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		// Build custom error messages
		var errorMessages []string
		for _, fieldErr := range validationErrors {
			message := generateErrorMessage(fieldErr)
			errorMessages = append(errorMessages, message)
		}
		return errors.New(fmt.Sprintf("Validation errors: %s", errorMessages))
	}

	return err // Unexpected error
}

// generateErrorMessage creates custom error messages for validation errors.
func generateErrorMessage(fieldErr validator.FieldError) string {
	switch fieldErr.Tag() {
	case "min":
		return fmt.Sprintf("The field '%s' must be at least %s characters long.", fieldErr.Field(), fieldErr.Param())
	case "max":
		return fmt.Sprintf("The field '%s' must be at most %s characters long.", fieldErr.Field(), fieldErr.Param())
	default:
		return fmt.Sprintf("The field '%s' is invalid.", fieldErr.Field())
	}
}
