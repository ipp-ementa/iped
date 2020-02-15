package canteen

import "github.com/ipp-ementa/iped/model/customerror"

// GeographicalLocation identifies a location which is positioned on a geographical map
// using its latitude and longitude
type GeographicalLocation struct {
	Latitude  float32
	Longitude float32
}

// New initializes a struct that is identified by the location latitude and longitude
// If the latitude received in parameters does not belong in the range of [-180, 180] an error is returned
// If the longitude received in parameters does not belong in the range of [-90, 90] an error is returned
func New(Latitude float32, Longitude float32) (GeographicalLocation, *customerror.FieldError) {

	geoLocation := GeographicalLocation{Latitude, Longitude}

	err := grantLatitudeIsNotLowerThanMinimumBoundary(Latitude)

	if err != nil {
		return geoLocation, err
	}

	err = grantLatitudeIsNotHigherThanMaximumBoundary(Latitude)

	if err != nil {
		return geoLocation, err
	}

	err = grantLongitudeIsNotLowerThanMinimumBoundary(Longitude)

	if err != nil {
		return geoLocation, err
	}

	err = grantLongitudeIsNotHigherThanMaximumBoundary(Longitude)

	return geoLocation, err
}

func grantLatitudeIsNotLowerThanMinimumBoundary(latitude float32) *customerror.FieldError {

	var err *customerror.FieldError

	if latitude < -180 {
		err = &customerror.FieldError{Field: "latitude", Model: "geographical location", Explanation: "geographical location latitude cannot be lower than -180"}
	}

	return err

}

func grantLatitudeIsNotHigherThanMaximumBoundary(latitude float32) *customerror.FieldError {

	var err *customerror.FieldError

	if latitude > 180 {
		err = &customerror.FieldError{Field: "latitude", Model: "geographical location", Explanation: "geographical location latitude cannot be higher than 180"}
	}

	return err

}

func grantLongitudeIsNotLowerThanMinimumBoundary(longitude float32) *customerror.FieldError {

	var err *customerror.FieldError

	if longitude < -90 {
		err = &customerror.FieldError{Field: "longitude", Model: "geographical location", Explanation: "geographical location longitude cannot be lower than -90"}
	}

	return err

}

func grantLongitudeIsNotHigherThanMaximumBoundary(longitude float32) *customerror.FieldError {

	var err *customerror.FieldError

	if longitude > 90 {
		err = &customerror.FieldError{Field: "longitude", Model: "geographical location", Explanation: "geographical location longitude cannot be higher than 90"}
	}

	return err

}
