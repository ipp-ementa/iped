package menu

import (
	"testing"

	"github.com/ipp-ementa/iped/model/dish"

	"github.com/ipp-ementa/iped/model/customerror"
)

func TestUnexistingMenuTypeReturnError(t *testing.T) {

	_dish := dish.Dish{Type: 0, Description: "Fried Noodles"}

	dishes := []dish.Dish{_dish}

	_, err := New(-1, dishes)

	if err == nil {
		t.Error("Menu initialization should have returned an error as there is no menu type for the value -1")
	}

	if err.(*customerror.FieldError).Field != "menutype" {
		t.Error("Even though that menu initialization returned an error, the error should have been caused by the field menutype")
	}
}

func TestExistingMenuTypeDoesNotReturnError(t *testing.T) {
	_dish := dish.Dish{Type: 0, Description: "Fried Noodles"}

	dishes := []dish.Dish{_dish}

	_, err := New(0, dishes)

	if err != nil {
		t.Errorf("Menu initilization should have been successful but got %s", err)
	}
}

func TestNilDishListReturnError(t *testing.T) {

	_, err := New(0, nil)

	if err == nil {
		t.Error("Menu initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "dishes" {
		t.Error("Even though that menu initialization returned an error, the error should have been caused by the field dishes")
	}
}

func TestEmptyDishListReturnError(t *testing.T) {
	_, err := New(0, []dish.Dish{})

	if err == nil {
		t.Error("Menu initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "dishes" {
		t.Error("Even though that menu initialization returned an error, the error should have been caused by the field dishes")
	}
}

func TestNotEmptyDishDescriptionDoesNotReturnError(t *testing.T) {
	_dish := dish.Dish{Type: 0, Description: "Fried Noodles"}

	dishes := []dish.Dish{_dish}

	_, err := New(0, dishes)

	if err != nil {
		t.Errorf("Menu initilization should have been successful but got %s", err)
	}
}
