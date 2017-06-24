package router

import (
	"testing"

	"github.com/labstack/echo"
)

var (
	dummyRoutesUpperCase = []Route{
		{"GET", "/", dummyHandler},
		{"POST", "/", dummyHandler},
		{"PUT", "/", dummyHandler},
		{"PATCH", "/", dummyHandler},
		{"DELETE", "/", dummyHandler},
		{"OPTIONS", "/", dummyHandler},
		{"HEAD", "/", dummyHandler},
		{"CONNECT", "/", dummyHandler},
	}
	dummyRoutesLowerCase = []Route{
		{"get", "/", dummyHandler},
		{"post", "/", dummyHandler},
		{"put", "/", dummyHandler},
		{"patch", "/", dummyHandler},
		{"delete", "/", dummyHandler},
		{"options", "/", dummyHandler},
		{"head", "/", dummyHandler},
		{"connect", "/", dummyHandler},
	}
	invalidRoutes = []Route{
		{"GET", "/", dummyHandler},
		{"POST", "/", dummyHandler},
		{"PUT", "/", dummyHandler},
		{"PATCH", "/", dummyHandler},
		{"DELETE", "/", dummyHandler},
		{"OPTIONS", "/", dummyHandler},
		{"FOOBAR", "/", dummyHandler},
		{"HEAD", "/", dummyHandler},
	}
)

func TestSetRouterUpperCase(t *testing.T) {
	e := echo.New()
	if err := SetRoutes(e, dummyRoutesUpperCase); err != nil {
		t.Error(err.Error())
	}
}

func TestSetRouterLowerCase(t *testing.T) {
	e := echo.New()
	if err := SetRoutes(e, dummyRoutesLowerCase); err != nil {
		t.Error(err.Error())
	}
}

func TestSetRouterInvalid(t *testing.T) {
	e := echo.New()
	if err := SetRoutes(e, invalidRoutes); err == nil {
		t.Error("FOOBAR should return error.")
	}
}

func dummyHandler(c echo.Context) (err error) {
	return c.String(200, "hello")
}
