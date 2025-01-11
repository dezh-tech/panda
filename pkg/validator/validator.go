package validator

import (
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validateInstance *Validator
	validatorLock    = &sync.Mutex{}
	validate         = validator.New()

	localeEN      = en.New()
	uni           = ut.New(localeEN, localeEN)
	translator, _ = uni.GetTranslator("en")
)

type Validator struct{}

// Validator initializes and returns a singleton instance of the validator.
func NewValidator() *Validator {
	if validateInstance == nil {
		validatorLock.Lock()
		defer validatorLock.Unlock()

		if validateInstance == nil {
			validateInstance = &Validator{}
			registerTranslations()
		}
	}

	return validateInstance
}

// Validate validates a given struct and returns custom validation errors.
func (*Validator) Validate(s interface{}) []*ValidationError {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		return formatValidationErrors(validationErrors)
	}

	// Return a generic error if the type assertion fails
	return []*ValidationError{
		{
			Field:   "unknown",
			Message: err.Error(),
		},
	}
}

// registerTranslations adds translations for validation error messages.
func registerTranslations() {
	en_translations.RegisterDefaultTranslations(validate, translator)
}

// formatValidationErrors formats the validation errors for API responses.
func formatValidationErrors(errs validator.ValidationErrors) []*ValidationError {
	var errors []*ValidationError
	for _, err := range errs {
		errors = append(errors, &ValidationError{
			Field:   err.Field(),
			Message: err.Translate(translator),
		})
	}
	return errors
}

// ValidationError represents a single validation error.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}