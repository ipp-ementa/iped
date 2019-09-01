package school

// This file contains model views representation for GET functionalities of schools collection

// GetAvailableSchoolsModelView is the model view representation
// for the available schools functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/schools.md#available-schools
type GetAvailableSchoolsModelView []struct {
	ID      int    `json:"id"`
	Acronym string `json:"acronym"`
	Name    string `json:"name"`
}

// GetDetailedSchoolInformationModelView is the model view representation
// for the detailed school information functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/schools.md#detailed-school-information
type GetDetailedSchoolInformationModelView struct {
	ID       int    `json:"id"`
	Acronym  string `json:"acronym"`
	Name     string `json:"name"`
	Canteens []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"canteens"`
}
