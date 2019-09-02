package canteen

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/ipp-ementa/iped/model/customerror"
	"github.com/ipp-ementa/iped/model/menu"
)

// Canteen is a model that has the responsibility to inform the user which menus are available at the time
// A canteen has a unique name and is offered by a school
// A UML overview of this model can be found at https://github.com/ipp-ementa/iped-documentation/wiki/Architecture#models-structure
type Canteen struct {
	gorm.Model
	Name  string
	menus map[time.Time][]menu.Menu
}

// New initializes a Canteen model using its name
// A FieldError is returned if the canteen name is empty
func New(Name string) (Canteen, error) {

	canteen := Canteen{gorm.Model{}, Name, map[time.Time][]menu.Menu{}}

	err := grantCanteenNameIsNotEmpty(Name)

	if err != nil {
		return canteen, err
	}

	return canteen, err
}

// AddTodayMenu allows the addition of a menu to today available menus
// A FieldError is returned if it was found to exist a menu that has the same type
// as the existing available menus
func (canteen *Canteen) AddTodayMenu(Menu menu.Menu) error {

	var err error

	availableMenus := canteen.AvailableMenus()

	if lena := len(availableMenus); lena != 0 {
		index := 0
		for index < lena {
			if availableMenus[index].Type == Menu.Type {
				index = lena
				err = &customerror.FieldError{Field: "menus", Model: "canteen"}
			} else {
				index++
			}
		}
	}

	if err == nil {
		availableMenus = append(availableMenus, Menu)
		todayDate := todayDateTime()
		canteen.menus[todayDate] = availableMenus
	}

	return err

}

// AvailableMenus returns the menus which the canteen is providing at the time being asked
// as a slice
// If no menus are available an empty slice is returned
// The returned slice has different reference of the one in Canteen struct
// In order to prevent modifications
func (canteen Canteen) AvailableMenus() []menu.Menu {

	todayDate := todayDateTime()

	availableMenus, exists := canteen.menus[todayDate]

	if !exists {
		availableMenus = []menu.Menu{}
	} else {
		availableMenusCopy := make([]menu.Menu, len(canteen.menus))

		copy(availableMenusCopy, availableMenus)
	}

	return availableMenus

}

// Equals compares equality between two canteens
// A canteen proves true equality to other canteen if both canteen names are equal
// Canteen names are case sensitive
func (canteen Canteen) Equals(comparingCanteen Canteen) bool {
	return strings.ToUpper(canteen.Name) == strings.ToUpper(comparingCanteen.Name)
}

// Returns today date as a [time.Time] struct
// The struct returned is formatted to be DD-MM-YYYY 00:00:00
func todayDateTime() time.Time {
	datetime := time.Now()

	datetime = time.Date(datetime.Year(), datetime.Month(), datetime.Day(), int(0), int(0), int(0), int(0), datetime.Location())

	return datetime
}

// This function grants that a canteen name is not empty, and if empty returns an error
func grantCanteenNameIsNotEmpty(name string) error {

	var err error

	if len(strings.TrimSpace(name)) == 0 {
		err = &customerror.FieldError{Field: "name", Model: "canteen"}
	}

	return err

}
