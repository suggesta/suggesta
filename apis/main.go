package main

import (
	"log"
	"os"
	"time"

	myMW "github.com/go-shosa/shosa/middleware"
	"github.com/go-shosa/shosa/server"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// If panic is happen, run this function.
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("Recover!:%v", err)
		}
	}()

	// url := os.Getenv("IP") + ":" + os.Getenv("PORT")
	url := ":" + os.Getenv("PORT")
	if url == "" {
		log.Fatal("Not set environment variable IP and PORT.")
	}

	// Get PID.
	log.Printf("start pid %d", os.Getpid())

	// Server runs
	server.RunWithConfig(server.Server{
		Middlewares: []echo.MiddlewareFunc{
			middleware.Recover(),
			middleware.CORS(),
			myMW.RequestID(),
			myMW.Logger(),
		},
		ShutdownTimeout: 15 * time.Second,
		URI:             url,
		Routing:         Routes,
	})
	// server.Run(url, Routes)
}
