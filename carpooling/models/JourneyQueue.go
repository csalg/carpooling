package models

import (
	// "time"
	// "fmt"
	"errors"
	"container/list"
)

// TODO:
// * Implement other methods necessary for that function to work
// * Write unit tests

type Journey struct {
	Id int `json:"id"`
	People int `json:"people"`
	InTransit bool
	InCar int
	timestamp int
}

func NewJourney(id int, people int){

}

type JourneyQueue struct {
	ById map[int]*Journey
	ByPeople [6]list.List  // Journeys can be 1-6 people.
}

// func (q *JourneyQueue) Insert (j *Journey) {
// 	// Create a pointer from id j
// 	// Stick at the tail of the queue by following the prev pointer to the head
// 	// Change both *tail.next and *head.prev to point to it
// 	// Insert a pointer to it in ByID
// }

// func (q *JourneyQueue) Delete () {
// 	// Follow the 
// }

// func (q *JourneyQueue) SetInTransit(j *Journey){

// }

// func (q *JourneyQueue) GetOldestSmallerThan(val int){

// }



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