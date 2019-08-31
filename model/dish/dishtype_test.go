package dish

import "testing"

func TestUnexistingDishTypeFailsValidateFunction(t *testing.T) {
	invalid := Validate(-1)

	if invalid {
		t.Error("Validate(-1) succeeded but should have failed as there is no dish type for the value -1")
	}

	invalid = Validate(4)

	if invalid {
		t.Error("Validate(4) succeeded but should have failed as there is no dish type for the value 4")
	}
}

func TestExistingDishTypeSucceedsValidateFunction(t *testing.T) {

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
