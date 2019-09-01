package canteen

import (
	"testing"

	"github.com/ipp-ementa/iped/model/dish"
	"github.com/ipp-ementa/iped/model/menu"

	"github.com/ipp-ementa/iped/model/customerror"
)

func TestEmptyCanteenNameReturnError(t *testing.T) {
	_, err := New("")

	if err == nil {
		t.Error("Canteen initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "name" {
		t.Error("Even though that canteen initialization returned an error, the error should have been caused by the field name")
	}
}

func TestCanteenNameWithOnlySpacesReturnError(t *testing.T) {
	_, err := New(" ")

	if err == nil {
		t.Error("Canteen initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "name" {
		t.Error("Even though that canteen initialization returned an error, the error should have been caused by the field name")
	}
}

func TestNotEmptyCanteenNameDoesNotReturnError(t *testing.T) {
	_, err := New("Cantina do H")

	if err != nil {
		t.Errorf("Canteen initilization should have been successful but got %s", err)
	}
}

func TestAvailableMenusMethodReturnsEmptySliceIfNoMenusAreAvailable(t *testing.T) {
	_canteen, _ := New("Cantina do H")

	availableMenus := _canteen.AvailableMenus()

	if lenab := len(availableMenus); lenab != 0 {
		t.Errorf("The length of availableMenus slice should be 0 but got: %d", lenab)
	}
}

func TestAvailableMenusMethodReturnsUnmodifiableSlice(t *testing.T) {
	_canteen, _ := New("Cantina do H")

	availableMenus := _canteen.AvailableMenus()

	// This verification is to grant that the returned available menus slice length is 0

	if lenab := len(availableMenus); lenab != 0 {
		t.Errorf("The length of availableMenus slice should be 0 but got: %d", lenab)
	}

	// If we add a new menu to the the returned slice,
	// it should not modify the slice pointed on the canteen struct

	_dish, _ := dish.New(0, "Fried Noodles")

	_menu, _ := menu.New(0, []dish.Dish{_dish})

	availableMenus = append(availableMenus, _menu)

	if lenam := len(availableMenus); lenam != 1 {
		t.Errorf("The length of availableMenus slice should now be 1 but got: %d", lenam)
	}

	availableMenusAfterModification := _canteen.AvailableMenus()

	if lenaam := len(availableMenusAfterModification); lenaam != 0 {
		t.Errorf("The length of availableMenusAfterModification slice should be 0 but got: %d", lenaam)
	}
}
