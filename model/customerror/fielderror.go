package customerror

import "fmt"

// FieldError is a custom error for invalid fields found during
// the initialization of a model
type FieldError struct {
	// this string should be the name of the field which is invalid
	Field string
	// this string should be the name of the model which holds the invalid field
	Model string
	// and this string should be the explanation of the error
	Explanation string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("Field: %s on Model: %s is invalid due to: %s", e.Field, e.Model, e.Explanation)
}
