package model

// FieldError is a custom error for invalid fields found during
// the initialization of a model
type FieldError struct {
	// this string should be the name of the field which is invalid
	Field string
	// and this string should be the name of the model which the error was created
	Model string
}
