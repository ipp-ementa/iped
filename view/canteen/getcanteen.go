package canteen

// This file contains model views representation for GET functionalities of canteens collection

// GetAvailableCanteensModelView is the model view representation
// for the available canteens functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/canteens.md#available-canteens
type GetAvailableCanteensModelView []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetDetailedCanteenInformationModelView is the model view representation
// for the detailed canteen information functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/canteens.md#detailed-canteen-information
type GetDetailedCanteenInformationModelView struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
