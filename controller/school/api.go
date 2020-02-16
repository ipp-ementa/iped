package school

import (
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ipp-ementa/iped/model/geographicallocation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	customerrormodel "github.com/ipp-ementa/iped/model/customerror"
	customerrorview "github.com/ipp-ementa/iped/view/customerror"

	"github.com/ipp-ementa/iped/model/canteen"

	model "github.com/ipp-ementa/iped/model/school"
	view "github.com/ipp-ementa/iped/view/school"

	"github.com/labstack/echo"
)

// AvailableSchools handles GET /schools functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/schools.md#available-schools
func AvailableSchools(c echo.Context) error {

	// db, ok := c.Get("db").(*mongo.Database)

	// if !ok {
	// 	return c.NoContent(http.StatusInternalServerError)
	// }

	// schools := []model.School{}

	// // Finds all available schools

	// err := db.Find(&schools).Error

	// if err != nil || len(schools) == 0 {
	// 	return c.NoContent(http.StatusNotFound)
	// }

	// modelview := view.ToGetAvailableSchoolsModelView(schools)

	// return c.JSON(http.StatusOK, modelview)

	return c.NoContent(200)

}

// DetailedSchoolInformation handles GET /schools/:id functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/schools.md#detailed-school-information
func DetailedSchoolInformation(c echo.Context) error {

	// db, ok := c.Get("db").(*mongo.Database)

	// if !ok {
	// 	return c.NoContent(http.StatusInternalServerError)
	// }

	// id, _ := strconv.Atoi(c.Param("id"))

	// var school model.School

	// // Finds school by ID

	// err := db.Find(&school, id).Error

	// if err != nil {
	// 	return c.NoContent(http.StatusNotFound)
	// }

	// // Find school canteens

	// db.Model(&school).Related(&school.CanteensSlice)

	// modelview := view.ToGetDetailedSchoolInformationModelView(school)

	// return c.JSON(http.StatusOK, modelview)

	return c.NoContent(200)

}

// CreateNewSchool handles POST /schools functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/schools.md#create-a-new-school
func CreateNewSchool(c echo.Context) error {

	db, ok := c.Get("db").(*mongo.Database)

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	var modelview view.CreateNewSchoolModelView

	c.Bind(&modelview)

	canteens := make([]canteen.Canteen, len(modelview.Canteens))

	for index := range modelview.Canteens {

		location, lerr := geographicallocation.New(modelview.Canteens[index].Location.Latitude, modelview.Canteens[index].Location.Longitude)

		if lerr != nil {

			modelview := customerrorview.UsingFieldErrorToErrorMessageModelView(*lerr)

			return c.JSON(http.StatusBadRequest, modelview)

		}

		canteen, cerr := canteen.New(modelview.Canteens[index].Name, location)
		if cerr != nil {

			modelview := customerrorview.UsingFieldErrorToErrorMessageModelView(*cerr)

			return c.JSON(http.StatusBadRequest, modelview)
		}
		canteens[index] = canteen
	}

	school, serr := model.New(modelview.Acronym, modelview.Name, canteens)

	if serr != nil {

		modelview := customerrorview.UsingFieldErrorToErrorMessageModelView(*serr)

		return c.JSON(http.StatusBadRequest, modelview)
	}

	// Finds if school with same acronym already exists

	collection := db.Collection("schools")

	filter := bson.M{"acronym": school.Acronym}

	ctx, cancelFunction := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelFunction()

	err := collection.FindOne(ctx, filter).Err()

	if err != mongo.ErrNoDocuments {

		cerr := customerrormodel.FieldError{Field: "acronym", Model: "school", Explanation: "a school with the same acronym already exists"}

		modelview := customerrorview.UsingFieldErrorToErrorMessageModelView(cerr)

		return c.JSON(http.StatusBadRequest, modelview)

	}

	document, err := bson.Marshal(school)

	if err != nil {

		return c.NoContent(http.StatusInternalServerError)

	}

	res, err := collection.InsertOne(ctx, document)

	if err != nil {

		return c.NoContent(http.StatusInternalServerError)

	}

	documentID := res.InsertedID.(primitive.ObjectID)

	id := documentID.Hex()

	school.ID = id

	modelviewres := view.ToGetDetailedSchoolInformationModelView(school)

	return c.JSON(http.StatusCreated, modelviewres)

}
