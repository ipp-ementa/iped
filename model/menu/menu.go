package menu

import (
	"github.com/ipp-ementa/iped/model/customerror"
	"github.com/ipp-ementa/iped/model/dish"
)

// Menu is a model that contains a set of dishes available at either lunch or dinner
// A UML overview of this model can be found at https://github.com/ipp-ementa/iped-documentation/wiki/Architecture#models-structure
type Menu struct {
	Type   MenuType
	dishes []dish.Dish
}

// New initializes a Menu model using a menu type and a set of dishes
// A FieldError is returned either if the menu type is invalid, no dishes were provided or the dishes are not unique
func New(Type int, Dishes []dish.Dish) (Menu, error) {

	menu := Menu{MenuType(Type), Dishes}

	err := grantValidMenuType(Type)

	if err != nil {
		return menu, err
	}

	err = grantThatAtLeastOneDishWasProvided(Dishes)

	if err != nil {
		return menu, err
	}

	err = grantNoDuplicatedDishesExist(Dishes)

	return menu, err
}

// Dishes returns the available dishes on a menu as a slice
// The returned slice has different reference of the one in Menu struct
// In order to prevent modifications
func (menu Menu) Dishes() []dish.Dish {
	availableDishes := make([]dish.Dish, len(menu.dishes))

	copy(availableDishes, menu.dishes)

	return availableDishes
}

// This function grants that a menu type is valid, and if not returns an error
// See [MenuType.Validate] for validation logic
func grantValidMenuType(menutype int) error {

	var err error

	if !Validate(menutype) {
		err = &customerror.FieldError{Field: "menutype", Model: "menu"}
	}

	return err
}

// This function grants that at least one dish is provided in given dish slice
// If the given dish slice is nil or empty an error is returned
func grantThatAtLeastOneDishWasProvided(dishes []dish.Dish) error {

	var err error

	if dishes == nil || len(dishes) == 0 {
		err = &customerror.FieldError{Field: "dishes", Model: "menu"}
	}

	return err
}

// This function grants that all dishes given in a slice are unique
// If a dish proves equality to any other dish in the slice, an error is returned
func grantNoDuplicatedDishesExist(dishes []dish.Dish) error {

	var err error

	unique := true
	dishesLength := len(dishes)
	i := 0

	for i < dishesLength {
		j := i + 1
		for j < dishesLength {
			unique = !dishes[i].Equals(dishes[j])
			j++
			if !unique {
				i = dishesLength
				j = i
			}
		}
		i++
	}

	if !unique {
		err = &customerror.FieldError{Field: "dishes", Model: "menu"}
	}

	return err
}
