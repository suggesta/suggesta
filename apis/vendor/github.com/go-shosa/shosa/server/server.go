package server

import (
	"time"

	"github.com/go-shosa/shosa/log"
	myMW "github.com/go-shosa/shosa/middleware"
	"github.com/go-shosa/shosa/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server defines the configuration for serving.
type Server struct {
	// labstack/echo instance
	Routing []router.Route

	// Middlewares labstack/echo middlewares
	Middlewares []echo.MiddlewareFunc

	// ShutdownTimeout is the duration before we begin force closing connections.
	ShutdownTimeout time.Duration

	// URI is server base URI
	URI string
}

var (
	// defaultServer returns default server configuration.
	defaultServer = Server{
		Middlewares: []echo.MiddlewareFunc{
			middleware.Recover(),
			myMW.RequestID(),
			myMW.Logger(),
		},
		ShutdownTimeout: 15 * time.Second,
		URI:             ":8080",
	}
)

// Run runs server with default config.
func Run(url string, routing []router.Route) {
	config := NewConfig(url, routing)
	RunWithConfig(config)
}

// RunWithConfig runs server with costom config.
func RunWithConfig(config Server) {
	if config.Middlewares == nil {
		config.Middlewares = defaultServer.Middlewares
	}
	if config.URI == "" {
		config.URI = defaultServer.URI
	}
	if config.ShutdownTimeout == 0 {
		config.ShutdownTimeout = defaultServer.ShutdownTimeout
	}
	if config.Routing == nil {
		log.Fatal("Routing parameter is missing. Server.Routing is required.")
	}

	i := newInstance(config.Routing, config.Middlewares)
	i.ShutdownTimeout = config.ShutdownTimeout
	i.Logger.Fatal(i.Start(config.URI))
}

// NewConfig creates a server configuration.
func NewConfig(url string, routing []router.Route) Server {
	conf := defaultServer
	conf.Routing = routing
	if url != "" {
		conf.URI = url
	}
	return conf
}

// newInstance creates a echo instance.
func newInstance(routing []router.Route, middlewares []echo.MiddlewareFunc) *echo.Echo {
	e := echo.New()
	for _, m := range middlewares {
		e.Use(m)
	}
	router.SetRoutes(e, routing)
	return e
}
