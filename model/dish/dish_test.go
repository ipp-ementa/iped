package dish

import (
	"github.com/ipp-ementa/iped/model/customerror"
	"testing"
)

func TestUnexistingDishTypeReturnError(t *testing.T) {
	_, err := New(-1, "Fried Noodles")

	if err == nil {
		t.Error("Dish initialization should have returned an error as there is no dish type for the value -1")
	}

	if err.(*customerror.FieldError).Field != "dishtype" {
		t.Error("Even though that dish initialization returned an error, the error should have been caused by the field dishtype")
	}
}

func TestExistingDishTypeDoesNotReturnError(t *testing.T) {
	_, err := New(0, "Fried Noodles")

	if err != nil {
		t.Errorf("Dish initilization should have been successful but got %s", err)
	}
}

func TestEmptyDishDescriptionReturnError(t *testing.T) {
	_, err := New(0, "")

	if err == nil {
		t.Error("Dish initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "description" {
		t.Error("Even though that dish initialization returned an error, the error should have been caused by the field description")
	}
}

func TestDescriptionWithOnlySpacesReturnError(t *testing.T) {
	_, err := New(0, " ")

	if err == nil {
		t.Error("Dish initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "description" {
		t.Error("Even though that dish initialization returned an error, the error should have been caused by the field description")
	}
}

func TestNotEmptyDishDescriptionDoesNotReturnError(t *testing.T) {
	_, err := New(0, "Fried Noodles")

	if err != nil {
		t.Errorf("Dish initilization should have been successful but got %s", err)
	}
}
