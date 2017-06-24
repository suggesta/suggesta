package main

import (
	"log"
	"os"

	"github.com/go-shosa/shosa/server"
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
	server.Run(url, Routes)
}
