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

	db, ok := c.Get("db").(*mongo.Database)

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	schools := []model.School{}

	collection := db.Collection("schools")

	ctx, cancelFunction := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelFunction()

	filter := bson.M{}

	schoolsCursor, err := collection.Find(ctx, filter)

	if err == mongo.ErrNoDocuments {
		return c.NoContent(http.StatusNotFound)
	}

	for schoolsCursor.Next(ctx) {

		school := model.School{}

		schoolsCursor.Decode(&school)

		schools = append(schools, school)

	}

	modelview := view.ToGetAvailableSchoolsModelView(schools)

	return c.JSON(http.StatusOK, modelview)

}

// DetailedSchoolInformation handles GET /schools/:id functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/schools.md#detailed-school-information
func DetailedSchoolInformation(c echo.Context) error {

	db, ok := c.Get("db").(*mongo.Database)

	if !ok {
		return c.NoContent(http.StatusInternalServerError)
	}

	id := c.Param("id")

	school := model.School{}

	// Finds school by ID

	collection := db.Collection("schools")

	oid, perr := primitive.ObjectIDFromHex(id)

	if perr != nil {

		return c.NoContent(http.StatusNotFound)

	}

	filter := bson.M{"_id": oid}

	ctx, cancelFunction := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelFunction()

	schoolResult := collection.FindOne(ctx, filter)

	err := schoolResult.Err()

	if err == mongo.ErrNoDocuments {

		return c.NoContent(http.StatusNotFound)

	}

	schoolResult.Decode(&school)

	modelview := view.ToGetDetailedSchoolInformationModelView(school)

	return c.JSON(http.StatusOK, modelview)

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

	_, err = collection.InsertOne(ctx, document)

	if err != nil {

		return c.NoContent(http.StatusInternalServerError)

	}

	modelviewres := view.ToGetDetailedSchoolInformationModelView(school)

	return c.JSON(http.StatusCreated, modelviewres)

}
