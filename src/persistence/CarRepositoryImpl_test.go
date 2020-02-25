package persistence

import (
	"fmt"
	"github.com/csalg/carpooling/src/domain/entities"
	"testing"
)

func TestCarQueueAdd(t *testing.T){

	cq := NewCarRepository()

	c, err := entities.NewCar(1,7)
	err = cq.Add(c)
	if err == nil {
		t.Errorf("Queue inserted a car outside range!")
	}

	c, err = entities.NewCar(2,6)
	err = cq.Add(c)
	if err != nil {
		t.Errorf("Queue didn't insert a valid car!")
	}

	if cq.Add(nil) == nil {
		t.Errorf("Queue inserted a null pointer!")
	}
}


func TestCarQueueMove(t *testing.T){
	
	c1, err1 := entities.NewCar(1,6)
	c2, err2 := entities.NewCar(2,6)
	if err1 != nil || err2 != nil{
		t.Errorf("Error constructing a valid car")
	}
	cq := NewCarRepository()
	err1 = cq.Add(c1)
	err2 = cq.Add(c2)
	if err1 != nil || err2 != nil{
		t.Errorf("Error adding valid cars")
	}

	el1 := cq.ById[1]
	el2 := cq.ById[2]
	err1 = cq.Move(el1, 8)
	if err1 == nil {
		t.Errorf("Moved car to invalid spot")
	}

	err2 = cq.Move(el2, 4)
	if err2 != nil {
		t.Errorf(err2.Error())
	}

	if cq.BySize[6].Back().Value == c2 {
		t.Errorf("Car was not dequeued")
	}

	if cq.BySize[4].Back().Value != c2 {
		fmt.Println(c2)
		fmt.Println(cq.BySize[4].Back().Value )
		fmt.Println(cq.BySize[4].Back()==el2 )
		t.Errorf("Car was not re-queued")
	}
}

func TestGetCarLargerThan(t *testing.T){
	c1, err1 := entities.NewCar(1,6)
	c2, err2 := entities.NewCar(2,6)
	if err1 != nil || err2 != nil{
		t.Errorf("Error constructing a valid car")
	}
	cq := NewCarRepository()
	err1 = cq.Add(c1)
	err2 = cq.Add(c2)
	if err1 != nil || err2 != nil{
		t.Errorf("Error adding valid cars")
	}

	el2 := cq.ById[2]
	err2 = cq.Move(el2, 2)
	if err2 != nil {
		t.Errorf(err1.Error())
	}

	for i := 0; i != 3; i++ {
		_, car, _ := cq.GetCarLargerThan(i)
		if car != c2 {
			t.Errorf("Unexpected result retrieving car larger than %d.", i)
		}
	}

	for i := 3; i < 6; i++ {
		_, car, _ := cq.GetCarLargerThan(i)
		if car != c1 {
			t.Errorf("Unexpected result retrieving car larger than %d.", i)
		}
	}
}


