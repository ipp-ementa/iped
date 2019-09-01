package school

import (
	"strings"

	"github.com/ipp-ementa/iped/model/canteen"
	"github.com/ipp-ementa/iped/model/customerror"
)

// School is a model that provides canteens
// A school has a unique acronym, a descriptive name and has to provide at least one canteen
// A UML overview of this model can be found at https://github.com/ipp-ementa/iped-documentation/wiki/Architecture#models-structure
type School struct {
	Acronym  string
	Name     string
	Canteens []canteen.Canteen
}

// New initializes a school model using its acronym, name and canteens
// A FieldError is returned if the canteen acronym is empty or has spaces between letters,
// name is empty, no canteens were provided or if it was found a duplicated canteen
func New(Acronym string, Name string, Canteens []canteen.Canteen) (School, error) {

	school := School{Acronym, Name, Canteens}

	err := grantSchoolAcronymIsNotEmpty(Acronym)

	if err != nil {
		return school, err
	}

	err = grantSchoolAcronymDoesNotHaveSpacesBetweenLetters(Acronym)

	if err != nil {
		return school, err
	}

	err = grantSchoolNameIsNotEmpty(Name)

	if err != nil {
		return school, err
	}

	err = grantAtLeastOneCanteenIsProvided(Canteens)

	if err != nil {
		return school, err
	}

	err = grantNoDuplicatedCanteensExist(Canteens)

	if err != nil {
		return school, err
	}

	return school, err
}

// This function grants that a school acronym is not empty, and if empty returns an error
func grantSchoolAcronymIsNotEmpty(acronym string) error {

	var err error

	if len(strings.TrimSpace(acronym)) == 0 {
		err = &customerror.FieldError{Field: "acronym", Model: "school"}
	}

	return err

}

// This function grants that a school acronym does not have spaces between letters, or else returns an error
func grantSchoolAcronymDoesNotHaveSpacesBetweenLetters(acronym string) error {

	var err error

	if len(acronym) != len(strings.TrimSpace(acronym)) {
		err = &customerror.FieldError{Field: "acronym", Model: "school"}
	}

	return err

}

// This function grants that a school name is not empty, and if empty returns an error
func grantSchoolNameIsNotEmpty(name string) error {

	var err error

	if len(strings.TrimSpace(name)) == 0 {
		err = &customerror.FieldError{Field: "name", Model: "school"}
	}

	return err

}

// This function grants that at least one canteen is provided in given canteen slice
// If the given canteen slice is nil or empty an error is returned
func grantAtLeastOneCanteenIsProvided(canteens []canteen.Canteen) error {

	var err error

	if canteens == nil || len(canteens) == 0 {
		err = &customerror.FieldError{Field: "canteens", Model: "school"}
	}

	return err

}

// This function grants that all canteen given in a slice are unique
// If a canteen proves equality to any other canteen in the slice, an error is returned
func grantNoDuplicatedCanteensExist(canteens []canteen.Canteen) error {

	var err error

	unique := true
	canteensLength := len(canteens)
	i := 0

	for i < canteensLength {
		j := i + 1
		for j < canteensLength {
			unique = !canteens[i].Equals(canteens[j])
			j++
			if !unique {
				i = canteensLength
				j = i
			}
		}
		i++
	}

	if !unique {
		err = &customerror.FieldError{Field: "canteens", Model: "school"}
	}

	return err
}
