package api

import (
	"bytes"
	"encoding/json"
	"github.com/unrolled/render"
	"strconv"

	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

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

	//Let's add 500 users
	for i := 1; i != 501; i++ {
		j, _ := models.NewJourney(i, i%5+1)
		b, err := json.Marshal(j)
		if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
		dispatchJSON(t, JourneyHandler, "POST", 200, b)
	}

	//Let's first test that the locate handler returns that they are waiting to be assigned, then
	//not found after dropoff
	for i := 1; i != 401; i++ {
		data = url.Values{}
		i_str := strconv.Itoa(i)
		data.Set("ID", i_str)
		dispatchForm(t, LocateHandler, 204, data)
		dispatchForm(t, DropoffHandler, 204, data)
		dispatchForm(t, LocateHandler, 404, data)

	}
	//
	////Now let's add a bunch of cars so that everyone is assigned and confirm that they are assigned
	//for i := 1; i != 100; i++ {
	//	models.NewCar(i,6)
	//
	//}
	//
	////All of those previous passengers should have been assigned.
	//for i := 400; i != 501; i++ {
	//	data = url.Values{}
	//	i_str := strconv.Itoa(i)
	//	data.Set("ID", i_str)
	//	dispatchForm(t, LocateHandler, 200, data) // Good request
	//}
		return
}