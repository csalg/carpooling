package main

import (
	"net/http"
	"github.com/unrolled/render"
	"encoding/json"
	// "fmt"
	// "io/ioutil"
	// "reflect"
	"strconv"
	"github.com/csalg/carpooling/queues"
)

// TO DO:
// * Move a bunch of logic further downstream to the queues and models

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
				return
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
	// For now I'm just going to do things the naive way.

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{

		case "POST":

			if r.Body == nil {
				http.Error(w, "Body cannot be empty", 400)
				return
			}

			journey_temp := Journey{}

			err := json.NewDecoder(r.Body).Decode(&journey_temp)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			if journey_temp.People < 1 || journey_temp.People > 6 {
				http.Error(w, "People must be between 1 and 6", 400)
				return
			}

			journeys = append(journeys, journey_temp)
			formatter.JSON(w,200,cars)
			return
		default:
			http.Error(w, "Not implemented!", 400)
		return
	}
}
}


func dropoffHandler (formatter *render.Render) http.HandlerFunc {
	// A group of people requests to be dropped off. Whether they traveled or not.

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{

		case "POST":

			if r.Body == nil {
				http.Error(w, "Body cannot be empty", 400)
				return
			}

			r.ParseForm()
			if len(r.Form["ID"]) == 0 {
				http.Error(w, "Error parsing ID", 400)
				return
			}

			id, err := strconv.Atoi(r.Form["ID"][0])
			if err != nil {
				http.Error(w, "Not an integer: " + strconv.Itoa(id),  400)
				return
			}

			for i := 0; i != len(journeys); i++ {
				if journeys[i].Id == id {
					return
				}
			}

			http.Error(w,"Not found", 404)
			return
			
		default:
			http.Error(w, "Not implemented", 400)
			return
		}
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