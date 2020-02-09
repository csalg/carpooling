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


func (car Car)  GetId() int { return car.Id }


func (car Car) GetSize() int { return car.seatsAvailable }


func (car Car) SetSize(val int) error {
	return 	car.SetSeatsAvailable(val)
}

func (car *Car) SetSeatsAvailable (val int) error{
	if val < 0 || val > 6 {
		return errors.New("Cars are required to have 0 to 6 available seats.")
	}	
	car.seatsAvailable = val
	return nil
}

// Serialization

type carJSON struct {
	Id int `json:"id"`
	Seats int `json:"seats"`
}

// BodyToCars deserializes from a json request to []Car
func BodyToCars(body io.ReadCloser) (*[]Car, error) {

	var tempCarsArray []carJSON;
	var carsArray []Car
	err := json.NewDecoder(body).Decode(&tempCarsArray)
	if err != nil { return nil, err }

	for _, car := range tempCarsArray {
		car, err := NewCar(car.Id, car.Seats)
		if err != nil { return nil, err }
		carsArray = append(carsArray, *car)
	}

	return &carsArray, nil
}
