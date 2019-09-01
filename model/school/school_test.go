package school

import (
	"testing"

	"github.com/ipp-ementa/iped/model/canteen"

	"github.com/ipp-ementa/iped/model/customerror"
)

func TestEmptySchoolAcronymReturnError(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("", "Instituto Superior de Engenharia do Porto", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "acronym" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field acronym")
	}
}

func TestSchoolAcronymWithOnlySpacesReturnError(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New(" ", "Instituto Superior de Engenharia do Porto", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "acronym" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field acronym")
	}
}

func TestSchoolAcronymWithSpacesBetweenLettersReturnError(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("IS EP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "acronym" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field acronym")
	}
}

func TestEmptySchoolNameReturnError(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("ISEP", "", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "name" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field name")
	}
}

func TestSchoolNameWithOnlySpacesReturnError(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("ISEP", " ", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "name" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field name")
	}
}

func TestIfNoSchoolCanteensAreProvidedAnErrorIsReturned(t *testing.T) {

	_canteens := []canteen.Canteen{}

	_, err := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "canteens" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field canteens")
	}
}

func TestIfDuplicatedSchoolCanteensAreProvidedAnErrorIsReturned(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")
	_canteens := []canteen.Canteen{_canteen, _canteen}

	_, err := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err == nil {
		t.Error("School initilization should have returned an error but got nil")
	}

	if err.(*customerror.FieldError).Field != "canteens" {
		t.Error("Even though that school initialization returned an error, the error should have been caused by the field canteens")
	}
}

func TestNotEmptyAndNoSpacesBetweenLettersSchoolAcronymDoesNotReturnError(t *testing.T) {
	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err != nil {
		t.Errorf("School initilization should have been successful but got %s", err)
	}
}

func TestNotEmptySchoolNameDoesNotReturnError(t *testing.T) {
	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err != nil {
		t.Errorf("School initilization should have been successful but got %s", err)
	}
}

func TestNotEmptyOrDuplicatedSchoolCanteensDoesNotReturnError(t *testing.T) {
	_canteen, _ := canteen.New("Cantina do H")

	_canteens := []canteen.Canteen{_canteen}

	_, err := New("ISEP", "Instituto Superior de Engenharia do Porto", _canteens)

	if err != nil {
		t.Errorf("School initilization should have been successful but got %s", err)
	}
}
