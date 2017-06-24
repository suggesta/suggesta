package main

import "github.com/go-shosa/shosa/router"

var (
	// Routes defines account service path routing.
	Routes = []router.Route{
		router.Route{
			Method:  "GET",
			Routing: "/emotion",
			Func:    EmotionLatest,
		},
		router.Route{
			Method:  "POST",
			Routing: "/emotion/image",
			Func:    Image,
		},
	}
)
