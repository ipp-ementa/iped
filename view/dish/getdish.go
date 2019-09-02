package dish

import "github.com/ipp-ementa/iped/model/dish"

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

// DishesToGetAvailableDishesModelView creates a GetAvailableDishesModelView using a slice of dishes
func DishesToGetAvailableDishesModelView(dishes []dish.Dish) GetAvailableDishesModelView {
	modelview := make(GetAvailableDishesModelView, len(dishes))

	for index, dish := range dishes {
		element := modelview[index]
		element.ID = int(dish.ID)
		element.Description = dish.Description
		element.Type = dish.Type.String()
	}

	return modelview
}
