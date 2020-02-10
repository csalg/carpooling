package api

import (
	"github.com/csalg/carpooling/src/data"
	"github.com/unrolled/render"
	//"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
)

var carQueue = data.NewCarQueue()
var journeyQueue = data.NewJourneyQueue()

var requestCounter = 0


// printRequest is a helper function which I wrote after the strange behaviour of those acceptance tests
// for debugging purposes. It uses the default logger, so it can also output to file by calling log.SetOutput(stream)
// somewhere.
func printRequest(request *http.Request){
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		log.Println(err)
	  }
	log.Println("\n---------------------------- # " + strconv.Itoa(requestCounter) + " # ----------------------------\n" )
	log.Println(string(requestDump))
	requestCounter++
}

// StatusHandler responds with a 200 OK when it handles a GET request
func StatusHandler (formatter *render.Render) http.HandlerFunc {

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		printRequest(request)
		switch request.Method{

		case "GET":
			formatter.JSON(responseWriter, 200, "Service running successfully")
			return
		default:
			http.Error(responseWriter, "Not implemented!", 400)
		return
	}
}
}


// CarsHandler loads the list of available cars in the service 
// and removes all previous data (existing journeys and cars).
// This method may be called more than once during the life cycle 
// of the service.
func CarsHandler (formatter *render.Render) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		printRequest(request)

		if 	request.Header.Get("Content-type") == "application/json" &&
			request.Method == "PUT" {
			err := carQueue.MakeFromJsonRequest(request.Body)
			if err != nil {
				http.Error(responseWriter, err.Error(), 400)
				return
			}

			journeyQueue = data.NewJourneyQueue()
			formatter.JSON(responseWriter,http.StatusOK,"Cars updated successfully")
			return
		} else {
			http.Error(responseWriter, "Wrong request format", 400)
			return
		}
	}
}

// JourneyHandler registers individual groups of people 
// looking for rides on the system
func JourneyHandler (formatter *render.Render) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		printRequest(request)

		if 	request.Header.Get("Content-type") == "application/json" &&
			request.Method == "POST" {

			err := journeyQueue.AddFromJsonRequest(request.Body)

			if err != nil {
				http.Error(responseWriter, err.Error(), 400)
				return
			}

			data.Match(carQueue, journeyQueue)
			responseWriter.WriteHeader(http.StatusOK)
			return
		} else {
			http.Error(responseWriter, "Not implemented!", 400)
		}
		return
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
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		printRequest(request)

		if  request.Header.Get("Content-type") == "application/x-www-form-urlencoded" &&
			request.Method == "POST" {
			request.ParseForm()
			if len(request.Form["ID"]) == 0 {
				http.Error(responseWriter, "Error parsing ID", 400)
				return
			}

			id, err := strconv.Atoi(request.Form["ID"][0])
			if err != nil {
				http.Error(responseWriter, "Not an integer: " + strconv.Itoa(id),  400)
				return
			}

			if !journeyQueue.Has(id){
				http.Error(responseWriter,"Not found", 404)
				return
			}
			err = data.Dropoff(carQueue, journeyQueue, id)
			if err != nil {
				http.Error(responseWriter,err.Error(), 400)
				return
			}
			formatter.JSON(responseWriter,204,"")
			return

		} else {
			http.Error(responseWriter, "Not implemented", 400)
		}
		return
	}
}

// LocateHandler returns the car the group is traveling
// with, or no car if they are still waiting to be served.
func LocateHandler (formatter *render.Render) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		printRequest(request)

		if 	request.Header.Get("Content-type") == "application/x-www-form-urlencoded" &&
			request.Method == "POST" {

			request.ParseForm()
			if len(request.Form["ID"]) == 0 {
				http.Error(responseWriter, "Error parsing ID", 400)
				return
			}

			id, err := strconv.Atoi(request.Form["ID"][0])
			if err != nil {
				http.Error(responseWriter, "Not an integer: " + strconv.Itoa(id),  400)
				return
			}

			if !journeyQueue.Has(id){
				responseWriter.WriteHeader(http.StatusNotFound)
				//http.Error(responseWriter,"Not found!", 404) // Acceptance test wants an empty body.
				return
			}

			_, journey, err := journeyQueue.GetById(id)
			if err != nil {
				http.Error(responseWriter,err.Error(), 400)
				return
			}

			if !journey.IsTravelling(){
				formatter.JSON(responseWriter,204,"")
				return
			} else {
				carJson, err := carQueue.GetCarJsonById(journey.Car)
				if err != nil {
					http.Error(responseWriter, "Error retrieving car", 400)
					return
				}
				formatter.JSON(responseWriter,200,carJson)
				return
			}

		} else {
			http.Error(responseWriter, "Not implemented!", 400)
		}
			return
	}
}