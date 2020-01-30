package models

import (
	"testing"
	// "fmt"
)

func TestNewCar(t *testing.T){

	for i := -10; i !=20; i++ {
		_, err := NewCar(1,i)
		if (i == 4 || i ==6) {
			if err != nil  { t.Errorf(err.Error()) }
		} else if err == nil {
			t.Errorf("Car with %d seats was created!", i)
		}	
	}
}

func TestCarQueueAdd(t *testing.T){
	cq := new(CarQueue)

	c, err := NewCar(1,7)
	err = cq.Add(c)
	if err == nil {
		t.Errorf("Queue inserted a car outside range!")
	}

	c, err = NewCar(2,6)
	err = cq.Add(c)
	if err != nil {
		t.Errorf("Queue didn't insert a perfectly valid car!")
	}

	if cq.Add(nil) == nil {
		t.Errorf("Queue inserted a null pointer!")
	}
}

func TestCarQueueMove(t *testing.T){
	c1, err1 := NewCar(1,6)
	c2, err2 := NewCar(2,6)
	if err1 != nil || err2 != nil{
		t.Errorf("Error constructing a valid car")
	}
	cq := new(CarQueue)
	err1 = cq.Add(c1)
	err2 = cq.Add(c2)
	if err1 != nil || err2 != nil{
		t.Errorf("Error adding valid cars")
	}

	err1 = cq.Move(c2, 8)
	if err1 == nil {
		t.Errorf("Moved car to invalid spot")
	}

	err1 = cq.Move(c2, 4)
	if err1 != nil {
		t.Errorf(err1.Error())
	}

	if cq.ByAvailableSeats[4] != c2 {
		t.Errorf("Car was not moved to new head")
	}

	if cq.ByAvailableSeats[6] != c1 {
		t.Errorf("Old head was not successfully updated")
	}
}

func TestGetCarLargerThan(t *testing.T){
	c1, err1 := NewCar(1,6)
	c2, err2 := NewCar(2,6)
	if err1 != nil || err2 != nil{
		t.Errorf("Error constructing a valid car")
	}
	cq := new(CarQueue)
	err1 = cq.Add(c1)
	err2 = cq.Add(c2)
	if err1 != nil || err2 != nil{
		t.Errorf("Error adding valid cars")
	}

	err1 = cq.Move(c2, 2)
	if err1 != nil {
		t.Errorf(err1.Error())
	}

	for i := 0; i != 3; i++ {
		if cq.GetCarLargerThan(i) != c2 {
			t.Errorf("Unexpected result retrieving car larger than %d.", i)
		}
	}

	for i := 3; i < 6; i++ {
		if cq.GetCarLargerThan(i) != c1 {
			t.Errorf("Unexpected result retrieving car larger than %d.", i)
		}
	}
}
