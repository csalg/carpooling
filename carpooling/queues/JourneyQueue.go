package queues

import (
	"time"
	"fmt"
	"errors"
	"container/list"
)


type Journey struct {
	Id int
	Size int
	Travelling bool
	InCar int
	timestamp int64
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

func (j Journey)  GetId() int	 { return j.Id }

func (j Journey) GetSize() int { return j.Size }





type JourneyQueue struct {
	HashQueue
	// ById map[int]*list.Element
	// ByPeople [6]list.List  // Journeys can be 1-6 people.
}


func NewJourneyQueue()*JourneyQueue {
	jq := new(JourneyQueue)
	jq.ById = make(map[int]*list.Element)
	return jq
}


// func (q *JourneyQueue) Add(j *Journey)error {

// 	if j == nil {
// 		return errors.New("Cannot add a null pointer.")
// 	}

// 	_, exists := q.ById[j.Id]
// 	if exists {
// 		return errors.New("Key already exists")
// 	}

// 	q.ByPeople[j.People-1].PushBack(*j)
// 	q.ById[j.Id] = 	q.ByPeople[j.People-1].Back()

// 	return nil
// }

// func (q *JourneyQueue) Delete (id int) error {
// 	el, ok := q.ById[id]
// 	if !ok { return errors.New("Id not found")}
// 	q.ByPeople[el.Value.(Journey).People-1].Remove(el)
// 	delete(q.ById, id)
// 	return nil

// }


func (q *JourneyQueue) GetOldestSmallerThan(val int)(*list.Element, error){
	if val > 6 || val < 1 { return nil, errors.New("Value is outside range") }
	timestamp := time.Now().UnixNano()
	el := new(list.Element)
	el = nil
	for i := 1; i != val+1; i++ {
		temp := q.BySize[i].Front()
		if temp != nil {
			fmt.Println(temp.Value.(*Journey))
			if temp.Value.(*Journey).timestamp < timestamp { 
				fmt.Println(temp.Value.(*Journey))
				timestamp = temp.Value.(*Journey).timestamp
				el = temp
			}
		}
	}
	return el, nil

}