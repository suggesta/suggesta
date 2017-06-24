package middleware

import "github.com/labstack/echo"

// defaultSkipper returns false which processes the middleware.
func defaultSkipper(c echo.Context) bool {
	return false
}
