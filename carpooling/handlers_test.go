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

func helper(t *testing.T,
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

	// cars := [Car{id:1,seats:5}]
	cars := []Car{
		{Id:1,
		Seats:5},
	}
	// c
	b, err := json.Marshal(cars)
	if err != nil {
		t.Errorf("Error marshalling into json: %v", err)
	}
	dispatch(t, carsHandler, "PUT", http.StatusOK, b)

	cars = []Car{
		{Id:1,
		Seats:10},
		{Id:2,
		Seats:2},
	}
	b, err = json.Marshal(cars)

	dispatch(t, carsHandler, "PUT", 400, b)
}