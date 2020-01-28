package main

import (
	"bytes"
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"strings"
	// "reflect"

	"github.com/unrolled/render"
)

// ------------------------------ Variables and factory functions --------------------------------------

var formatter = render.New(render.Options{
	IndentJSON: true,
})

func dispatch(t *testing.T,
	client *http.Client,
	req *http.Request,
	handler func  (formatter *render.Render) http.HandlerFunc, 
	method string,
	expected_status int) {
	// Sends requests, asserts payload and status code are correct

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

func dispatchJSON (t *testing.T,
	handler func  (formatter *render.Render) http.HandlerFunc, 
	method string,
	expected_status int,
	body []byte) {
	// Factory for json requests

	client := &http.Client{}
	server := httptest.NewServer(
		http.HandlerFunc(handler(formatter)))
	defer server.Close()

	req,err := http.NewRequest(method, server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error creating %s request: %v", method, err)
	}

	req.Header.Add("Content-Type", "application/json")

	dispatch(t, client, req, handler, method, expected_status)
	return
}

func dispatchForm(t *testing.T,
	handler func  (formatter *render.Render) http.HandlerFunc, 
	expected_status int,
	data url.Values) {
	// Factory for form requests

	method := "POST"
	client := &http.Client{}
	server := httptest.NewServer(
		http.HandlerFunc(dropoffHandler(formatter)))
	defer server.Close()

	b := strings.NewReader(data.Encode())

	req,err := http.NewRequest(method, server.URL, b)
	if err != nil {  t.Errorf("Error creating %s request: %v", "POST", err)  }

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	dispatch(t, client, req, dropoffHandler, method, expected_status)
	return 
}

// ------------------------------ Tests --------------------------------------


func TestCars(t *testing.T){

	dispatchJSON(t, carsHandler, "PUT", 400, []byte{})

	cars := []Car{
		{Id:1,
		Seats:5},
	}
	b, err := json.Marshal(cars)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatchJSON(t, carsHandler, "PUT", 200, b)

	cars = []Car{
		{Id:1,
		Seats:10},
		{Id:2,
		Seats:2},
	}
	b, err = json.Marshal(cars)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatchJSON(t, carsHandler, "PUT", 400, b)
}


func TestJourney(t *testing.T){

	dispatchJSON(t, journeyHandler, "POST", 400, []byte{}) // **Body** _required_ A form with the group ID, such that `ID=X`

	journey := Journey{ Id:1, People:5 }
	b, err := json.Marshal(journey)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatchJSON(t, journeyHandler, "POST", 200, b)

	journey = Journey{ Id:2, People:50 }
	b, err = json.Marshal(journeys)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatchJSON(t, journeyHandler, "POST", 400, b)

}

func TestDropoff(t *testing.T){
	// **Body** _required_ A form with the group ID, such that `ID=X`
	// **Content Type** `application/x-www-form-urlencoded`
	// Responses:
	// * **200 OK** or **204 No Content** When the group is unregistered correctly.
	// * **404 Not Found** When the group is not to be found.
	// * **400 Bad Request** When there is a failure in the request format or the
	//   payload can't be unmarshalled.

	data := url.Values{}
	data.Set("ID", "10")
	dispatchForm(t, dropoffHandler, 404, data) // Not found

	data = url.Values{}
	data.Set("foo", "bar")
	dispatchForm(t, dropoffHandler, 400, data) // Bad request

	data = url.Values{}
	dispatchForm(t, dropoffHandler, 400, data) // Bad request (empty body)


	journey := Journey{ Id:50, People:5 }
	b, err := json.Marshal(journey)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatchJSON(t, journeyHandler, "POST", 200, b)
	data = url.Values{}
	data.Set("ID", "50")
	dispatchForm(t, dropoffHandler, 200, data) // Good request

	return 




	
	// res,err := client.Do(req)
	// if err != nil {
	// 	t.Errorf("Error sending %s request: %v", method, err)
	// }

	// defer res.Body.Close()

	// payload, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	t.Errorf("Error reading response body: %v, %s", err, payload)
	// }

	// if res.StatusCode != expected_status {
	// 	t.Errorf("Expected response status %d, received %s", expected_status, res.Status)
	// }

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