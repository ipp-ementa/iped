package dish

import "testing"

func TestInvalidDishTypeFailsValidateFunction(t *testing.T) {
	invalid := Validate(-1)

	if invalid {
		t.Error("Validate(-1) succeeded but should have failed as there is no dish type for the value -1")
	}

	invalid = Validate(4)

	if invalid {
		t.Error("Validate(4) succeeded but should have failed as there is no dish type for the value 4")
	}
}

func TestValidDishTypeSucceedsValidateFunction(t *testing.T) {

	// First test is for Meat dish type

	valid := Validate(0)

	if !valid {
		t.Error("Validate(0) failed but should have succeeded as 0 is the Meat DishType")
	}

	// Second test is for Fish dish type

	valid = Validate(1)

	if !valid {
		t.Error("Validate(1) failed but should have succeeded as 1 is the Fish DishType")
	}

	// Third test is for Vegetarian dish type

	valid = Validate(2)

	if !valid {
		t.Error("Validate(2) failed but should have succeeded as 2 is the Vegetarian DishType")
	}

	// Fourth test is for Diet dish type

	valid = Validate(3)

	if !valid {
		t.Error("Validate(3) failed but should have succeeded as 3 is the Diet DishType")
	}
}

func TestCallingStringOnInvalidDishTypeReturnsNilString(t *testing.T) {
	invalidDishType := DishType(-1)

	if str := invalidDishType.String(); str != "nil" {
		t.Errorf("DishType(-1) is not valid so calling String() should have returned the string 'nil' but got: %s ", str)
	}

	invalidDishType = DishType(4)

	if str := invalidDishType.String(); str != "nil" {
		t.Errorf("DishType(4) is not valid so calling String() should have returned the string 'nil' but got: %s ", str)
	}
}

func TestCallingStringOnValidDishTypeReturnsProperString(t *testing.T) {
	meatDishType := DishType(0)

	if str := meatDishType.String(); str != "meat" {
		t.Errorf("DishType(0) is valid so calling String() should have returned the string 'meat' but got: %s ", str)
	}

	fishDishType := DishType(1)

	if str := fishDishType.String(); str != "fish" {
		t.Errorf("DishType(1) is valid so calling String() should have returned the string 'fish' but got: %s ", str)
	}

	vegetarianDishType := DishType(2)

	if str := vegetarianDishType.String(); str != "vegetarian" {
		t.Errorf("DishType(2) is valid so calling String() should have returned the string 'vegetarian' but got: %s ", str)
	}

	dietDishType := DishType(3)

	if str := dietDishType.String(); str != "diet" {
		t.Errorf("DishType(3) is valid so calling String() should have returned the string 'diet' but got: %s ", str)
	}

}
