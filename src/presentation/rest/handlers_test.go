package rest

import (
	"bytes"
	"encoding/json"
	"github.com/csalg/carpooling/src/domain/entities"
	"github.com/unrolled/render"
	"strconv"

	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
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

	c1, err1 := entities.NewCar(1,6)
	c2, _ := entities.NewCar(2,4)

	if err1 != nil {
		t.Errorf(err1.Error())
	}

	cars := []entities.Car{*c1,*c2}

	b, err := json.Marshal(cars)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatchJSON(t, Cars, "PUT", 200, b)

	dispatchJSON(t, Cars, "PUT", 400, []byte{})
}


func TestJourney(t *testing.T){

	dispatchJSON(t, Journey, "POST", 400, []byte{}) // **Body** _required_ A form with the group ID, such that `ID=X`

	journey, _ := entities.NewJourney(1,5)
	b, err := json.Marshal(journey)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	dispatchJSON(t, Journey, "POST", 200, b)

	j := entities.Journey{Id: 2, Size: 50}
	b, err = json.Marshal(j)
	if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
	if err == nil { dispatchJSON(t, Journey, "POST", 400, b) }

}

func TestDropoffAndLocate(t *testing.T){
	data := url.Values{}
	data.Set("ID", "10")
	dispatchForm(t, Locate, 404, data)  // Not found
	dispatchForm(t, Dropoff, 404, data) // Not found

	data = url.Values{}
	data.Set("foo", "bar")
	dispatchForm(t, Dropoff, 400, data) // Bad request

	data = url.Values{}
	dispatchForm(t, Dropoff, 400, data) // Bad request (empty body)

	//Let's add 500 users
	for i := 1; i != 501; i++ {
		j, _ := entities.NewJourney(i, i%5+1)
		b, err := json.Marshal(j)
		if err != nil {  t.Errorf("Error marshalling into json: %v", err) }
		dispatchJSON(t, Journey, "POST", 200, b)
	}

	//Let's first test that the locate handler returns that they are waiting to be assigned, then
	//not found after dropoff
	for i := 1; i != 401; i++ {
		data = url.Values{}
		iStr := strconv.Itoa(i)
		data.Set("ID", iStr)
		dispatchForm(t, Locate, 204, data)
		dispatchForm(t, Dropoff, 204, data)
		dispatchForm(t, Locate, 404, data)

	}
	//
	////Now let's add a bunch of cars so that everyone is assigned and confirm that they are assigned
	//for i := 1; i != 100; i++ {
	//	domain.NewCar(i,6)
	//
	//}
	//
	////All of those previous passengers should have been assigned.
	//for i := 400; i != 501; i++ {
	//	persistence = url.Values{}
	//	i_str := strconv.Itoa(i)
	//	persistence.Set("ID", i_str)
	//	dispatchForm(t, Locate, 200, persistence) // Good request
	//}
		return
}