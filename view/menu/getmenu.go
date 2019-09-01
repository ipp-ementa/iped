package menu

// This file contains model views representation for GET functionalities of menus collection

// GetAvailableMenusModelView is the model view representation
// for the available menus functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/menus.md#available-menus
type GetAvailableMenusModelView []struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

// GetDetailedMenuInformationModelView is the model view representation
// for the detailed menu information functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/menus.md#detailed-menu-information
type GetDetailedMenuInformationModelView struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Dishes []struct {
		ID          int    `json:"id"`
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"dishes"`
}
