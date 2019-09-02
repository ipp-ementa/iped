package school

import (
	"net/http"
	"strconv"

	model "github.com/ipp-ementa/iped/model/school"
	view "github.com/ipp-ementa/iped/view/school"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// AvailableSchools handles GET /schools functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/schools.md#available-schools
func AvailableSchools(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB)

	if !ok {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	schools := []model.School{}

	// Finds all available schools

	err := db.Find(&schools).Error

	if err != nil || len(schools) == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	modelview := view.ToGetAvailableSchoolsModelView(schools)

	return c.JSON(http.StatusOK, modelview)

}

// DetailedSchoolInformation handles GET /schools/:id functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/schools.md#available-schools
func DetailedSchoolInformation(c echo.Context) error {

	db, ok := c.Get("db").(*gorm.DB)

	if !ok {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	id, perr := strconv.Atoi(c.Param("id"))

	if perr != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	var school model.School

	// Finds school by ID

	err := db.Find(&school, id).Error

	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	modelview := view.ToGetDetailedSchoolInformationModelView(school)

	return c.JSON(http.StatusOK, modelview)

}
