package main

import (
	"net/http"
	"github.com/unrolled/render"
	"encoding/json"
	// "fmt"
	// "io/ioutil"
)

var cars []Car
var journeys []Journey

func carsHandler (formatter *render.Render) http.HandlerFunc {
	// Load the list of available cars in the service and remove all previous data
	// (existing journeys and cars). This method may be called more than once during 
	// the life cycle of the service.

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			if r.Body == nil {
				http.Error(w, "Body cannot be empty!", 400)
			}
			var cars_temp [] Car;
			err := json.NewDecoder(r.Body).Decode(&cars_temp)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			for i:=0; i!=len(cars_temp); i++ {
				if (cars_temp[i].Seats > 6 || cars_temp[i].Seats < 4) {
					http.Error(w, "Input validation error: cars can only have between 4 and 6 seats!", 400)
					return
				}
				cars = cars_temp
				formatter.JSON(w,http.StatusOK,cars)
				return
			}
		default:
			http.Error(w, "Wrong method", 400)
			return
		}
	}
}

func journeyHandler (formatter *render.Render) http.HandlerFunc {
	// A group of people requests to perform a journey.

	return func(w http.ResponseWriter, r *http.Request) {

		http.Error(w, "Not implemented!", 400)
		return
	}
}

func dropoffHandler (formatter *render.Render) http.HandlerFunc {
	// A group of people requests to be dropped off. Whether they traveled or not.

	return func(w http.ResponseWriter, r *http.Request) {

		http.Error(w, "Not implemented!", 400)
		return
	}
}

func locateHandler (formatter *render.Render) http.HandlerFunc {
	// Given a group ID such that `ID=X`, return the car the group is traveling
	// with, or no car if they are still waiting to be served.

	return func(w http.ResponseWriter, r *http.Request) {

		http.Error(w, "Not implemented!", 400)
		return
	}
}