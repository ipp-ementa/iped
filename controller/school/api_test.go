package school

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/ipp-ementa/iped/model/canteen"
	"github.com/ipp-ementa/iped/model/school"

	"github.com/ipp-ementa/iped/controller/db"

	"github.com/labstack/echo"
)

var ech *echo.Echo

func init() {
	os.Setenv("IPEW_CONNECTION_STRING", ":memory:")

	ech = echo.New()

}

func TestAvailableSchoolsReturnsInternalServerErrorIfDatabaseIsNotAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")

	// GET /schools
	AvailableSchools(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("No middleware was setup to provide the database connection object to the request, so the response status code should be 500 but was: %d", rec.Code)
	}
}

func TestAvailableSchoolsReturnsNotFoundIfNoSchoolsExist(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")
	ctx.Set("db", db.Db)

	// GET /schools
	AvailableSchools(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("As there are no schools available in the database the response status code should be 404 but was: %d", rec.Code)
	}

}

func TestAvailableSchoolsReturnsOKIfSchoolsExist(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do H")

	_school, _ := school.New("ISEP", "Instituto Superior de Engenharia do Porto", []canteen.Canteen{_canteen})

	// Inserts a school in the database

	db.Db.Create(&_school)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools")
	ctx.Set("db", db.Db)

	// GET /schools/:id
	AvailableSchools(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("There is one school available in the database so the response status code should be 200 but was: %d", rec.Code)
	}

}

func TestDetailedSchoolInformationReturnsInternalServerErrorIfDatabaseIsNotAvailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	// GET /schools
	DetailedSchoolInformation(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("No middleware was setup to provide the database connection object to the request, so the response status code should be 500 but was: %d", rec.Code)
	}
}

func TestDetailedSchoolInformationReturnsNotFoundIfResourceIdIsNotAnIntegerParsable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("non integer")
	ctx.Set("db", db.Db)

	// GET /schools/:id
	DetailedSchoolInformation(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("The resource id is not a parsable intenger so the response status code should be 404 but was: %d", rec.Code)
	}
}

func TestDetailedSchoolInformationReturnsNotFoundIfResourceWasNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("0")
	ctx.Set("db", db.Db)

	// GET /schools/:id
	DetailedSchoolInformation(ctx)

	if rec.Code != http.StatusNotFound {
		t.Errorf("There is no school with the resource id '0' so the response status code should be 404 but was: %d", rec.Code)
	}
}

func TestDetailedSchoolInformationReturnsOKIfResourceWasFound(t *testing.T) {

	_canteen, _ := canteen.New("Cantina do ESMAE")

	_school, _ := school.New("ESMAE", "Escola Superior de Música e Artes do Espetáculo", []canteen.Canteen{_canteen})

	db.Db.Create(&_school)

	id := _school.ID

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := ech.NewContext(req, rec)
	ctx.SetPath("/schools/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.FormatUint(uint64(id), 10))
	ctx.Set("db", db.Db)

	// GET /schools/:id
	DetailedSchoolInformation(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("There is a school with the resource id %d so the response status code should be 200 but was: %d", id, rec.Code)
	}
}
