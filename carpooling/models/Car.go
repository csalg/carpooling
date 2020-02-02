package models

import (
	"errors"
	"io"
	// "fmt"
	"encoding/json"
)

type Car struct {
	Id int `json:"id"`
	Seats int `json:"seats"`
	seatsAvailable int
}

func NewCar(id int, seats int) (*Car, error) {
	if !(seats == 4 || seats == 6){
		return nil, errors.New("Cars are required to have 4 to 6 seats")
	}
	c := Car{Id:id,Seats:seats}
	c.SetSeatsAvailable(seats)
	return &c, nil
}


func (c Car)  GetId() int	 { return c.Id }


func (c Car) GetSize() int { return c.seatsAvailable }


func (c Car) SetSize(val int) error { 
	return 	c.SetSeatsAvailable(val)
}

func (c *Car) SetSeatsAvailable (val int) error{
	if val < 0 || val > 6 {
		return errors.New("Cars are required to have 0 to 6 available seats.")
	}	
	c.seatsAvailable = val
	return nil
}

// Serialization

type carJSON struct {
	Id int `json:"id"`
	Seats int `json:"seats"`
}

// BodyToCars deserializes from a json request to []Car
func BodyToCars(b io.ReadCloser) (*[]Car, error) {

	var cars_temp []carJSON;
	var cars []Car
	err := json.NewDecoder(b).Decode(&cars_temp)
	if err != nil { return nil, err }

	for _, c := range cars_temp {
		car, err := NewCar(c.Id, c.Seats)
		if err != nil { return nil, err }
		cars = append(cars, *car)
	}

	return &cars, nil
}
