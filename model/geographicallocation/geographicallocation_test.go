package geographicallocation

import (
	"testing"
)

func TestLatitudeWithAValueThatIsLowerThanMinimumBoundaryReturnError(t *testing.T) {

	latitude := float32(-180.01)

	longitude := float32(45.0)

	_, err := New(latitude, longitude)

	if err == nil {
		t.Error("Geographical Location struct initilization should have returned an error but got nil")
	}

	if err.Field != "latitude" {
		t.Error("Even though that geographical location struct initialization returned an error, the error should have been caused by the field latitude")
	}
}

func TestLatitudeWithAValueThatIsHigherThanMaximumBoundaryReturnError(t *testing.T) {

	latitude := float32(180.01)

	longitude := float32(45.0)

	_, err := New(latitude, longitude)

	if err == nil {
		t.Error("Geographical Location struct initilization should have returned an error but got nil")
	}

	if err.Field != "latitude" {
		t.Error("Even though that geographical location struct initialization returned an error, the error should have been caused by the field latitude")
	}
}

func TestLongitudeWithAValueThatIsLowerThanMinimumBoundaryReturnError(t *testing.T) {

	latitude := float32(45.0)

	longitude := float32(-90.01)

	_, err := New(latitude, longitude)

	if err == nil {
		t.Error("Geographical Location struct initilization should have returned an error but got nil")
	}

	if err.Field != "longitude" {
		t.Error("Even though that geographical location struct initialization returned an error, the error should have been caused by the field longitude")
	}
}

func TestLongitudeWithAValueThatIsHigherThanMaximumBoundaryReturnError(t *testing.T) {

	latitude := float32(45)

	longitude := float32(180.01)

	_, err := New(latitude, longitude)

	if err == nil {
		t.Error("Geographical Location struct initilization should have returned an error but got nil")
	}

	if err.Field != "longitude" {
		t.Error("Even though that geographical location struct initialization returned an error, the error should have been caused by the field longitude")
	}
}
