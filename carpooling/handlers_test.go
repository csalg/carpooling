package main

import (
	"bytes"
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/unrolled/render"
)

var formatter = render.New(render.Options{
	IndentJSON: true,
})

func dispatch(t *testing.T,
	 handler func  (formatter *render.Render) http.HandlerFunc, 
	 method string,
	 expected_status int,
	 body []byte) {

	client := &http.Client{}
	server := httptest.NewServer(
		http.HandlerFunc(handler(formatter)))
	defer server.Close()

	req,err := http.NewRequest(method, server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error creating %s request: %v", method, err)
	}

	req.Header.Add("Content-Type", "application/json")

	res,err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending %s request: %v", method, err)
	}

	defer res.Body.Close()

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v, %s", err, payload)
	}

	if res.StatusCode != expected_status {
		t.Errorf("Expected response status %d, received %s", expected_status, res.Status)
	}
}

func TestCars(t *testing.T){

	dispatch(t, carsHandler, "PUT", 400, []byte{})

	cars := []Car{
		{Id:1,
		Seats:5},
	}
	b, err := json.Marshal(cars)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatch(t, carsHandler, "PUT", 200, b)

	cars = []Car{
		{Id:1,
		Seats:10},
		{Id:2,
		Seats:2},
	}
	b, err = json.Marshal(cars)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatch(t, carsHandler, "PUT", 400, b)
}


func TestJourney(t *testing.T){

	dispatch(t, journeyHandler, "POST", 400, []byte{})

	journey := Journey{ Id:1, People:5 }
	b, err := json.Marshal(journey)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatch(t, journeyHandler, "POST", 200, b)

	journey = Journey{ Id:2, People:50 }
	b, err = json.Marshal(journeys)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatch(t, journeyHandler, "POST", 400, b)

}

func TestDropoff(t *testing.T){
	// **Body** _required_ A form with the group ID, such that `ID=X`
	// **Content Type** `application/x-www-form-urlencoded`
	// Responses:
	// * **200 OK** or **204 No Content** When the group is unregistered correctly.
	// * **404 Not Found** When the group is not to be found.
	// * **400 Bad Request** When there is a failure in the request format or the
	//   payload can't be unmarshalled.
}

func TestLocate(t *testing.T){
	// **Body** _required_ A url encoded form with the group ID such that `ID=X`
	// **Content Type** `application/x-www-form-urlencoded`
	// **Accept** `application/json`
	// Responses:
	// * **200 OK** With the car as the payload when the group is assigned to a car.
	// * **204 No Content** When the group is waiting to be assigned to a car.
	// * **404 Not Found** When the group is not to be found.
	// * **400 Bad Request** When there is a failure in the request format or the
	// payload can't be unmarshalled.
}