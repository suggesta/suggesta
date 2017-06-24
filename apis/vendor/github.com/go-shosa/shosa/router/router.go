package router

import (
	"fmt"
	"strings"

	"github.com/labstack/echo"
)

// Route defines the route for method and path with matching handler.
type Route struct {
	Method  string
	Routing string
	Func    echo.HandlerFunc
}

// SetRoutes registers routes for methods and paths with matching handlers.
func SetRoutes(e *echo.Echo, routes []Route) (err error) {
	for _, rt := range routes {
		switch strings.ToUpper(rt.Method) {
		case "GET":
			e.GET(rt.Routing, rt.Func)
		case "POST":
			e.POST(rt.Routing, rt.Func)
		case "PUT":
			e.PUT(rt.Routing, rt.Func)
		case "PATCH":
			e.PATCH(rt.Routing, rt.Func)
		case "DELETE":
			e.DELETE(rt.Routing, rt.Func)
		case "OPTIONS":
			e.OPTIONS(rt.Routing, rt.Func)
		case "HEAD":
			e.HEAD(rt.Routing, rt.Func)
		case "CONNECT":
			e.CONNECT(rt.Routing, rt.Func)
		default:
			return fmt.Errorf("invalid HTTP method. method: %s. routing: %s", rt.Method, rt.Routing)
		}
	}
	return nil
}
