package menu

import "testing"

func TestUnexistingMenuTypeFailsValidateFunction(t *testing.T) {
	invalid := Validate(-1)

	if invalid {
		t.Error("Validate(-1) succeeded but should have failed as there is no menu type for the value -1")
	}

	invalid = Validate(2)

	if invalid {
		t.Error("Validate(2) succeeded but should have failed as there is no menu type for the value 2")
	}
}

func TestExistingMenuTypeSucceedsValidateFunction(t *testing.T) {

	// First test is for Lunch menu type

	valid := Validate(0)

	if !valid {
		t.Error("Validate(0) failed but should have succeeded as 0 is the Lunch MenuType")
	}

	// Second test is for Dinner menu type

	valid = Validate(1)

	if !valid {
		t.Error("Validate(1) failed but should have succeeded as 1 is the Dinner MenuType")
	}
}
