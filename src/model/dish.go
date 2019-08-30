package model

import "strings"

// Dish is a model for what a person can choose to eat in canteen
// A UML overview of this model can be found at https://github.com/ipp-ementa/iped-documentation/wiki/Architecture
type Dish struct {
	Type        DishType
	Description string
}

// New initializes a Dish model using a dish type and a description
// An error is returned if the dish description is empty
func New(Type DishType, Description string) (Dish, error) {

	err := grantValidDescription(Description)
	dish := Dish{Type, Description}

	return dish, err
}

// This function grants that a dish description is valid, and if not returns an error
// A dish description is invalid if it is empty
func grantValidDescription(description string) error {

	var err *FieldError

	if len(strings.TrimSpace(description)) == 0 {
		err = &FieldError{"description", "dish"}
	}

	return err
}
