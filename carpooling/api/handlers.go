package main

import (
	"net/http"
	"github.com/unrolled/render"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "reflect"
	"strconv"
	"github.com/csalg/carpooling/data"
	// "github.com/csalg/carpooling/models"
)

// TO DO:
// * Move a bunch of logic further downstream to the queues and models

var cq = data.NewCarQueue()
var jq = data.NewJourneyQueue()

//var q := data.NewCarAndJourneysQueue()

// CarsHandler loads the list of available cars in the service 
// and removes all previous data (existing journeys and cars).
// This method may be called more than once during the life cycle 
// of the service.
func CarsHandler (formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":

			// err := models.ValidateJsonCars(r.Body)
			// if err...
			// q.ResetCars()
			// q.AddCarsFromJsonRequest()
			// q.ResetJourneys()

			err := cq.MakeFromJsonRequest(r.Body)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			jq = data.NewJourneyQueue()
			formatter.JSON(w,http.StatusOK,"Cars updated successfully")
			return

		default:
			http.Error(w, "Wrong method", 400)
			return
		}
	}
}

// JourneyHandler registers individual groups of people 
// looking for rides on the system
func JourneyHandler (formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{

		case "POST":

			// err := models.ValidateJsonJourney(r.Body)
			// if err...
			//jq.AddFromJsonRequest(r.Body)


			err := jq.AddFromJsonRequest(r.Body)

			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			data.Match(cq,jq)
			formatter.JSON(w,200,"Successfully posted")
			return
		default:
			http.Error(w, "Not implemented!", 400)
		return
	}
}
}

// DropoffHandler deletes journeys from the system. Specs: 
// **Body** _required_ A form with the group ID, such that `ID=X`
// **Content Type** `application/x-www-form-urlencoded`
// Responses:
// * **200 OK** or **204 No Content** When the group is unregistered correctly.
// * **404 Not Found** When the group is not to be found.
// * **400 Bad Request** When there is a failure in the request format or the
//   payload can't be unmarshalled.
func DropoffHandler (formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{

		case "POST":

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

			if !jq.Has(id){
				http.Error(w,"Not found", 404)
				return
			}

			err = data.Dropoff(cq, jq, id)
			//q.Dropoff(id)
			if err != nil { http.Error(w,err.Error(), 400) }
			return
			
		default:
			http.Error(w, "Not implemented", 400)
			return
		}
	}
}

// LocateHandler returns the car the group is traveling
// with, or no car if they are still waiting to be served.
func LocateHandler (formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		http.Error(w, "Not implemented!", 400)
		return
	}
}