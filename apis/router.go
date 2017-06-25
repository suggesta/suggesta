package main

import "github.com/go-shosa/shosa/router"

var (
	// Routes defines account service path routing.
	Routes = []router.Route{
		router.Route{
			Method:  "GET",
			Routing: "/calendar",
			Func:    Calendar,
		},
		router.Route{
			Method:  "GET",
			Routing: "/emotion",
			Func:    EmotionIndex,
		},
		router.Route{
			Method:  "GET",
			Routing: "/emotion/summary",
			Func:    EmotionSummary,
		},
		router.Route{
			Method:  "POST",
			Routing: "/emotion/image",
			Func:    Image,
		},
	}
)
