package server

import (
	"os"
	"os/exec"
	"reflect"
	"testing"

	myMW "github.com/go-shosa/shosa/middleware"
	"github.com/go-shosa/shosa/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	dummyRoutes = []router.Route{
		{"GET", "/", dummyHandler},
		{"POST", "/", dummyHandler},
		{"PUT", "/", dummyHandler},
		{"PATCH", "/", dummyHandler},
		{"DELETE", "/", dummyHandler},
		{"OPTIONS", "/", dummyHandler},
		{"HEAD", "/", dummyHandler},
		{"CONNECT", "/", dummyHandler},
	}
)

func TestNewConfigWithoutRouting(t *testing.T) {
	conf := NewConfig("", nil)
	if !reflect.DeepEqual(conf, defaultServer) {
		t.Errorf("conf should be equal to defaultServer.\nactual: %#v\nexpected: %#v", conf, defaultServer)
	}
}

func TestNewConfig(t *testing.T) {
	// actual
	conf := NewConfig("", dummyRoutes)
	// expected
	conf2 := NewConfig("", nil)
	conf2.Routing = dummyRoutes

	if !reflect.DeepEqual(conf, conf2) {
		t.Errorf("conf should be equal to conf2.\nactual: %#v\nexpected: %#v", conf, conf2)
	}
}

func TestNewInstance(t *testing.T) {
	ms := []echo.MiddlewareFunc{
		middleware.Recover(),
		myMW.RequestID(),
		myMW.Logger(),
	}
	newInstance(dummyRoutes, ms)
}

func TestRun(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("This test is not expected panic.")
		}
	}()

	go Run("", dummyRoutes)
}

func TestRunWithoutRouting(t *testing.T) {
	if os.Getenv("BE_TEST_RUN_WITHOUT_ROUTING") == "1" {
		Run("", nil)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestRunWithoutRouting")
	cmd.Env = append(os.Environ(), "BE_TEST_RUN_WITHOUT_ROUTING=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func dummyHandler(c echo.Context) (err error) {
	return c.String(200, "hello")
}
