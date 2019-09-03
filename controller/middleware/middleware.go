package middleware

import (
	"github.com/ipp-ementa/iped/controller/db"
	"github.com/labstack/echo"
)

// DbAccessMiddleware is a middleware that sets the database connection object using "db" as key
func DbAccessMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db.Db)
			next(c)
			return nil
		}
	}
}
