package canteen

import (
	"strings"
	"time"

	"github.com/ipp-ementa/iped/model/customerror"
	"github.com/ipp-ementa/iped/model/menu"
)

// Canteen is a model that has the responsibility to inform the user which menus are available at the time
// A canteen has a unique name and is offered by a school
// A UML overview of this model can be found at https://github.com/ipp-ementa/iped-documentation/wiki/Architecture#models-structure
type Canteen struct {
	Name  string
	Menus map[time.Time][]menu.Menu
}

// New initializes a Canteen model using its name
// A FieldError is returned if the canteen name is invalid
func New(Name string) (Canteen, error) {

	canteen := Canteen{Name, map[time.Time][]menu.Menu{}}

	err := grantCanteenNameIsNotEmpty(Name)

	if err != nil {
		return canteen, err
	}

	return canteen, err
}

// This function grants that a canteen name is not empty, and if empty returns an error
func grantCanteenNameIsNotEmpty(name string) error {

	var err error

	if len(strings.TrimSpace(name)) == 0 {
		err = &customerror.FieldError{Field: "name", Model: "canteen"}
	}

	return err

}
