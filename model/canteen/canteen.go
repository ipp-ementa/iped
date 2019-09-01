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
	menus map[time.Time][]menu.Menu
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

// AvailableMenus returns the menus which the canteen is providing at the time being asked
// If no menus are available an empty slice is returned
// The returned slice is unmodifiable in order to prevent modifications
func (canteen Canteen) AvailableMenus() []menu.Menu {

	todayDate := time.Now()

	todayDate = time.Date(todayDate.Year(), todayDate.Month(), todayDate.Day(), int(0), int(0), int(0), int(0), todayDate.Location())

	availableMenus, exists := canteen.menus[todayDate]

	if !exists {
		availableMenus = []menu.Menu{}
	}

	return availableMenus

}

// This function grants that a canteen name is not empty, and if empty returns an error
func grantCanteenNameIsNotEmpty(name string) error {

	var err error

	if len(strings.TrimSpace(name)) == 0 {
		err = &customerror.FieldError{Field: "name", Model: "canteen"}
	}

	return err

}
