package queues

import (
	"errors"
	"container/list"
)


// TODO:
// * Implement that cq.match function
// * Move out the models to their own subpackage so it's cleaner
// * Create methods for marshalling and unmarshalling json
// * Hook things up with the API handlers


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


func (c *Car) SetSeatsAvailable (val int) error{
	if val < 0 || val > 6 {
		return errors.New("Cars are required to have 0 to 6 available seats.")
	}	
	c.seatsAvailable = val
	return nil
}





type CarQueue struct {
	ByAvailableSeats [7]list.List // Cars can have 0-6 seats available
}


/*
A constructor is not necessary; simply create a new instance 
by either `q := new(CarQueue)` or `q := CarQueue{}`
*/


func (q *CarQueue) Add(c *Car) error {
	if c == nil {
		return errors.New("Cannot insert a null pointer")
	}

	q.ByAvailableSeats[c.seatsAvailable].PushFront(c)
	return nil
}


func (q *CarQueue) Move(c *Car, seatsAvailable int) error {
	if c == nil {
		return errors.New("Cannot move a null pointer")
	}

	if q.ByAvailableSeats[c.seatsAvailable].Front().Value != c {
		return errors.New("Car not found in head of linked list.")
	}

	prev := c.seatsAvailable
	err := c.SetSeatsAvailable(seatsAvailable)
	if err != nil {
		return err
	}
	q.ByAvailableSeats[prev].Remove(q.ByAvailableSeats[prev].Front())
	q.ByAvailableSeats[seatsAvailable].PushFront(c)
	return nil
}


func (q *CarQueue) GetCarLargerThan(val int) *Car {
	for i := val; i <= 6; i++ {
		if q.ByAvailableSeats[i].Front() != nil { 
			c, ok := q.ByAvailableSeats[i].Front().Value.(*Car)
			if ok { return c }
		}
	}
	return nil
}


func (q *CarQueue) AssignCar(c *Car, j *Journey) error {
	// if c.seatsAvailable < j.People { 
	// 	return errors.New("Cannot assign car with less seats than people in the journey") 
	// }
	// return q.Move(c, c.seatsAvailable - j.People)
	return nil
}


// // func (cq *CarQueue) Match(jq *JourneyQueue{
// // 	// Matches all possible journeys to available cars in 
// // 	// journey arrival order

// // 	maxAvailable := 6
// // 	for maxAvailable < 0 {
// // 	// For starters, we need to know what the largest car capacity is so that we can efficiently
// // 	// filter the journeys queue.
// // 	for cq.ByAvailableSeats[maxAvailable] == nil && maxAvailable > 0 {
// // 		maxAvailable--
// // 	}

// // 	oldest_journey := jq.GetOldestSmallerThan(maxAvailable) 
// // 	if oldest_journey == nil {
// // 		// In this case all journeys in the queue are of more people than the
// // 		// largest car available so we exit the loop.
// // 		break
// // 	}

// // 	smallest_car := cq.GetCarLargerThan(oldest_journey.People)

// // 	cq.AssignCar(smallest_car, oldest_journey)
// // 	jq.SetInTransit(oldest_journey)


// // 	// timestamp := time.Now().Unix()
// // 	// oldest_journey := new(Journey);

// // 	// fmt.Println(timestamp)
// // 	// for i := 0; i != maxAvailable-1; i++ { // up to maxAvailable -1 because indexing is different
// // 	// 										// Will probably move this implementation detail to a method
// // 	// 										// getOldestSmallerThan()
// // 	// 	if jq.ByPeople[i]

// // 	// }

// // 	// All of this can be done in O(1), so it scales well
// // 	}

// // }