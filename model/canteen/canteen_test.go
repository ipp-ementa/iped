package canteen

import (
	"testing"

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
