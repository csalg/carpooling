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
	j := new(Journey)
	j.Id = id
	j.Size = size
	j.SetTimestamp()
	return j, nil
}

func (j *Journey) SetTimestamp(){
	j.timestamp = time.Now().UnixNano()
}

func (j *Journey) GetTimestamp()int64 {
	return j.timestamp
}

func (j *Journey)  GetId() int	 { return j.Id }

func (j *Journey) GetSize() int { return j.Size }

func (j *Journey) SetSize(val int) error { return nil }

func (j *Journey) IsTravelling() bool { return j.isTravelling }

func (j *Journey) AssignCar(id int) { 
	j.Car = id
	j.isTravelling = true
 }

type journeyJSON struct {
	Id int `json:"id"`
	People int `json:"people"`
}

func NewJourneyFromBody(b io.ReadCloser) (*Journey, error) {
	var jTemp journeyJSON
	var j *Journey

	err := json.NewDecoder(b).Decode(&jTemp)
	if err != nil { return nil, err }

	j, err = NewJourney(jTemp.Id, jTemp.People)
	if err != nil { return nil, err }

	return j, nil
}
