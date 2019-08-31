package dish

import (
	"testing"

	"github.com/ipp-ementa/iped/model/customerror"
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

func TestDishesWithDifferentTypesAreNotEqual(t *testing.T) {
	dishone, _ := New(0, "Fried Noodles")

	dishtwo, _ := New(1, "Fried Noodles")

	equality := dishone.Equals(dishtwo)

	if equality {
		t.Errorf("dishone has dish type: %d and dishtwo has dish type: %d, which are different but were proved to be equal", dishone.Type, dishtwo.Type)
	}
}

func TestDishesWithDifferentDescriptionsAreNotEqual(t *testing.T) {
	dishone, _ := New(0, "Fried Noodles")

	dishtwo, _ := New(0, "Fries with beef")

	equality := dishone.Equals(dishtwo)

	if equality {
		t.Errorf("dishone has description: %s and dishtwo has description: %s, which are different but were proved to be equal", dishone.Description, dishtwo.Description)
	}
}

func TestDishesWithDifferentTypesAndDescriptionsAreNotEqual(t *testing.T) {
	dishone, _ := New(0, "Fried Noodles")

	dishtwo, _ := New(1, "Fries with beef")

	equality := dishone.Equals(dishtwo)

	if equality {
		t.Error("dishone and dishtwo both have different dish types and descriptions but were proved to be equal")
	}
}

func TestDishesWithEqualTypesAndDescriptionsAreEqual(t *testing.T) {
	dishone, _ := New(0, "Fried Noodles")

	dishtwo, _ := New(0, "Fried Noodles")

	equality := dishone.Equals(dishtwo)

	if !equality {
		t.Errorf("dishone and dishtwo both have equal dish types and descriptions but were proved to be different")
	}
}
