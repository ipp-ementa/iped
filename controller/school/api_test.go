package school

import (
	"net/http"
	"net/http/httptest"
	"os"
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

	// GET /schools
	AvailableSchools(ctx)

	if rec.Code != http.StatusOK {
		t.Errorf("There is one school available in the database so the response status code should be 200 but was: %d", rec.Code)
	}

}
