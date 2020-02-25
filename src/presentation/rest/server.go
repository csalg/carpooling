package rest

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a negroni handler.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: false,
		DisableCharset: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

// initRoutes is where all the url routing happens (maps url strings to handler functions)
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/status", Status(formatter))
	mx.HandleFunc("/cars", Cars(formatter))
	mx.HandleFunc("/journey", Journey(formatter))
	mx.HandleFunc("/dropoff", Dropoff(formatter))
	mx.HandleFunc("/locate", Locate(formatter))
}
