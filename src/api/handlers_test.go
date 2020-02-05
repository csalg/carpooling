package api

import (
	"bytes"
	"encoding/json"
	"github.com/unrolled/render"
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	// "reflect"
	"strconv"

	"github.com/csalg/carpooling/src/models"
)

// ------------------------------ Variables and factory functions --------------------------------------

var formatter = render.New(render.Options{
	IndentJSON: true,
})

func dispatch(t *testing.T,
	client *http.Client,
	req *http.Request,
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
		t.Errorf("Expected response status %d, received %s. Request body: %s", expected_status, res.Status, payload)
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

	dispatch(t, client, req, method, expected_status)
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
		http.HandlerFunc(handler(formatter)))
	defer server.Close()

	b := strings.NewReader(data.Encode())

	req,err := http.NewRequest(method, server.URL, b)
	if err != nil {  t.Errorf("Error creating %s request: %v", "POST", err)  }

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	dispatch(t, client, req, method, expected_status)
	return 
}

// ------------------------------ Tests --------------------------------------


func TestCars(t *testing.T){

	c1, err1 := models.NewCar(1,6)
	c2, _ := models.NewCar(2,4)

	if err1 != nil {
		t.Errorf(err1.Error())
	}

	cars := []models.Car{*c1,*c2}

	b, err := json.Marshal(cars)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatchJSON(t, CarsHandler, "PUT", 200, b)

	dispatchJSON(t, CarsHandler, "PUT", 400, []byte{})
}


func TestJourney(t *testing.T){

	dispatchJSON(t, JourneyHandler, "POST", 400, []byte{}) // **Body** _required_ A form with the group ID, such that `ID=X`

	journey, _ := models.NewJourney(1,5)
	b, err := json.Marshal(journey)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatchJSON(t, JourneyHandler, "POST", 200, b)

	j := models.Journey{Id: 2, Size: 50}
	b, err = json.Marshal(j)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	if err == nil { dispatchJSON(t, JourneyHandler, "POST", 400, b) }

}

func TestDropoffAndLocate(t *testing.T){
	data := url.Values{}
	data.Set("ID", "10")
	dispatchForm(t, LocateHandler, 404, data) // Not found
	dispatchForm(t, DropoffHandler, 404, data) // Not found

	data = url.Values{}
	data.Set("foo", "bar")
	dispatchForm(t, DropoffHandler, 400, data) // Bad request

	data = url.Values{}
	dispatchForm(t, DropoffHandler, 400, data) // Bad request (empty body)

	for i := 1; i != 500; i++ {
		j, _ := models.NewJourney(i, i%5+1)
		b, err := json.Marshal(j)
		if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
		dispatchJSON(t, JourneyHandler, "POST", 200, b)
	}

	for i := 1; i != 500; i++ {
		data = url.Values{}
		i_str := strconv.Itoa(i)
		data.Set("ID", i_str)
		dispatchForm(t, DropoffHandler, 200, data) // Good request
	}
	return 
}

// func TestLocate(t *testing.T){
// 	// **Body** _required_ A url encoded form with the group ID such that `ID=X`
// 	// **Content Type** `application/x-www-form-urlencoded`
// 	// **Accept** `application/json`
// 	// Responses:
// 	// * **200 OK** With the car as the payload when the group is assigned to a car.
// 	// * **204 No Content** When the group is waiting to be assigned to a car.
// 	// * **404 Not Found** When the group is not to be found.
// 	// * **400 Bad Request** When there is a failure in the request format or the
// 	// payload can't be unmarshalled.

// }