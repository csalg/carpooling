package main

import (
	"github.com/codegangsta/negroni"
	"github.com/csalg/carpooling/src/api"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/cars", api.CarsHandler(formatter))
	mx.HandleFunc("/journey", api.JourneyHandler(formatter))
	// mx.HandleFunc("/dropoff", DropoffHandler(formatter))
	// mx.HandleFunc("/locate", LocateHandler(formatter))
}
