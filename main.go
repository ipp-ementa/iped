// This gofile is the entrypoint for iped
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ipp-ementa/iped/controller/canteen"
	"github.com/ipp-ementa/iped/controller/db"
	"github.com/ipp-ementa/iped/controller/dish"
	"github.com/ipp-ementa/iped/controller/menu"
	"github.com/ipp-ementa/iped/controller/middleware"
	"github.com/ipp-ementa/iped/controller/school"
	"github.com/labstack/echo"
)

func main() {

	echo.NotFoundHandler = middleware.NotFoundHandler()

	ech := echo.New()

	ech.Use(middleware.DbAccessMiddleware())

	ech.Use(middleware.ResourceIdentifierValidationMiddleware())

	// schools collection functionalities

	ech.GET("/api/schools", school.AvailableSchools)

	ech.GET("/api/schools/:id", school.DetailedSchoolInformation)

	ech.POST("/api/schools", school.CreateNewSchool)

	// canteens collection functionalities

	ech.GET("/api/schools/:id/canteens", canteen.AvailableCanteens)

	ech.GET("/api/schools/:id/canteens/:id2", canteen.DetailedCanteenInformation)

	ech.POST("/api/schools/:id/canteens", canteen.CreateNewCanteen)

	// menus collection functionalities

	ech.GET("/api/schools/:id/canteens/:id2/menus", menu.AvailableMenus)

	ech.GET("/api/schools/:id/canteens/:id2/menus/:id3", menu.DetailedMenuInformation)

	ech.POST("/api/schools/:id/canteens/:id2/menus", menu.CreateNewMenu)

	// dishes collection functionalities

	ech.GET("/api/schools/:id/canteens/:id2/menus/:id3/dishes", dish.AvailableDishes)

	ech.GET("/api/schools/:id/canteens/:id2/menus/:id3/dishes/:id4", dish.DetailedDishInformation)

	port, perr := strconv.Atoi(os.Getenv("PORT"))

	if perr != nil {
		panic(fmt.Sprint("Server couldn't be open as the specified port is not valid"))
	}

	ech.Start(fmt.Sprintf(":%d", port))

	defer db.Db.Close()
}
