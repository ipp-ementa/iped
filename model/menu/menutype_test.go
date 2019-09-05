package menu

import "testing"

func TestInvalidMenuTypeFailsValidateFunction(t *testing.T) {
	invalid := Validate(-1)

	if invalid {
		t.Error("Validate(-1) succeeded but should have failed as there is no menu type for the value -1")
	}

	invalid = Validate(2)

	if invalid {
		t.Error("Validate(2) succeeded but should have failed as there is no menu type for the value 2")
	}
}

func TestValidMenuTypeSucceedsValidateFunction(t *testing.T) {

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

func TestCallingStringOnInvalidMenuTypeReturnsNilString(t *testing.T) {
	invalidMenuType := MenuType(-1)

	if str := invalidMenuType.String(); str != "nil" {
		t.Errorf("MenuType(-1) is not valid so calling String() should have returned the string 'nil' but got: %s ", str)
	}

	invalidMenuType = MenuType(2)

	if str := invalidMenuType.String(); str != "nil" {
		t.Errorf("MenuType(2) is not valid so calling String() should have returned the string 'nil' but got: %s ", str)
	}
}

func TestCallingStringOnValidMenuTypeReturnsProperString(t *testing.T) {
	lunchMenuType := MenuType(0)

	if str := lunchMenuType.String(); str != "lunch" {
		t.Errorf("MenuType(0) is valid so calling String() should have returned the string 'lunch' but got: %s ", str)
	}

	dinnerMenuType := MenuType(1)

	if str := dinnerMenuType.String(); str != "dinner" {
		t.Errorf("MenuType(1) is valid so calling String() should have returned the string 'dinner' but got: %s ", str)
	}

}

func TestParseInvalidMenuTypeStringReturnsInvalidMenuType(t *testing.T) {
	invalidMenuType := Parse("menu")

	if invalidMenuType != -1 {
		t.Errorf("'menu' is not a valid menu type so Parse should have returned -1 but got: %d ", invalidMenuType)
	}
}

func TestParseValidMenuTypeStringReturnsValidMenuType(t *testing.T) {
	validMenuType := Parse("lunch")

	if validMenuType != 0 {
		t.Errorf("'lunch' is a valid menu type so Parse should have returned 0 but got: %d ", validMenuType)
	}

	validMenuType = Parse("dinner")

	if validMenuType != 1 {
		t.Errorf("'dinner' is a valid menu type so Parse should have returned 1 but got: %d ", validMenuType)
	}
}
