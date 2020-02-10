package models

import (
	"time"
	"fmt"
	"errors"
	"encoding/json"
	"io"
)
type Journey struct {
	Id           int `json:"id"`
	Size         int `json:"people"`
	isTravelling bool
	Car          int
	timestamp    int64
}


func NewJourney(id int, size int) (*Journey, error) {
	if size > 6 || size < 1 {
		return nil, errors.New(fmt.Sprintf("Number of people must be between 1 and 6. Got %d", size))
	}
	journey := new(Journey)
	journey.Id = id
	journey.Size = size
	journey.SetTimestamp()
	return journey, nil
}


// SetTimestamp marks the current timestamp
func (journey *Journey) SetTimestamp(){
	journey.timestamp = time.Now().UnixNano()
}


func (journey *Journey) GetTimestamp()int64 {
	return journey.timestamp
}


func (journey *Journey)  GetId() int {
	return journey.Id
}


func (journey *Journey) GetSize() int {
	return journey.Size
}


func (journey *Journey) SetSize(val int) error {
	return errors.New("Journey size can only be set during construction.")
}


func (journey *Journey) IsTravelling() bool {
	return journey.isTravelling
}


func (journey *Journey) AssignCar(id int) {
	journey.Car = id
	journey.isTravelling = true
 }


type journeyJSON struct {
	Id int `json:"id"`
	People int `json:"people"`
}


func NewJourneyFromBody(body io.ReadCloser) (*Journey, error) {
	var jsonJourney journeyJSON
	var journey *Journey

	err := json.NewDecoder(body).Decode(&jsonJourney)
	if err != nil { return nil, err }

	journey, err = NewJourney(jsonJourney.Id, jsonJourney.People)
	if err != nil { return nil, err }

	return journey, nil
}
