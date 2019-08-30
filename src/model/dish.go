package model

import "strings"

// Dish is a model for what a person can choose to eat in canteen
// A UML overview of this model can be found at https://github.com/ipp-ementa/iped-documentation/wiki/Architecture
type Dish struct {
	Type        DishType
	Description string
}

// New initializes a Dish model using a dish type and a description
// A FieldError is returned either if the dish description is empty or the dish type isn't valid
func New(Type int, Description string) (Dish, error) {

	dish := Dish{DishType(Type), Description}

	err := grantValidDishType(Type)

	if err != nil {
		return dish, err
	}

	err = grantValidDescription(Description)

	return dish, err
}

// This function grants that a dish description is valid, and if not returns an error
// A dish description is invalid if it is empty
func grantValidDescription(description string) error {

	var err error

	if len(strings.TrimSpace(description)) == 0 {
		err = &FieldError{"description", "dish"}
	}

	return err
}

// This function grants that a dish type is valid, and if not returns an error
// See [DishType.Validate] for validation logic
func grantValidDishType(dishtype int) error {

	var err error

	if !Validate(dishtype) {
		err = &FieldError{"dishtype", "dish"}
	}

	return err
}
