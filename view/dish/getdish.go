package dish

// This file contains model views representation for GET functionalities of dishes collection

// GetAvailableDishesModelView is the model view representation
// for the available dishes functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/dishes.md#available-dishes
type GetAvailableDishesModelView []struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// GetDetailedDishInformationModelView is the model view representation
// for the detailed dish information functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/dishes.md#detailed-dish-information
type GetDetailedDishInformationModelView struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
