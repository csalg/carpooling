package data

import (
	"container/list"
	"errors"
	"github.com/csalg/carpooling/backing"
	"github.com/csalg/carpooling/models"
	"io"
	"time"
)

type JourneyQueue struct {
	backing.HashQueue
}

func NewJourneyQueue()*JourneyQueue {
	jq := new(JourneyQueue)
	jq.ById = make(map[int]*list.Element)
	return jq
}

func (q *JourneyQueue) GetOldestSmallerThan(val int)(*list.Element, *models.Journey, error){
	if val > len(q.BySize) || val < 1 { 
		return nil, nil, errors.New("Value is outside range") 
	}
	
	timestamp := time.Now().UnixNano()
	element   := new(list.Element)
	journey   := new(models.Journey)

	element, journey = nil, nil

	for i := 1; i != val+1; i++ {
		temp := q.BySize[i].Front()
		if temp != nil {
			journeyCandidate := q.BySize[i].Front().Value.(*models.Journey)
			if journeyCandidate.GetTimestamp() < timestamp { 
				element   = temp
				journey   = journeyCandidate
				timestamp = journeyCandidate.GetTimestamp()
			}
		}
	}
	if element == nil {
		return nil, nil, errors.New("No journey found!")
	}
	return element, journey, nil
}

func (q *JourneyQueue) AddFromJsonRequest(b io.ReadCloser) error {
	j, err := models.NewJourneyFromBody(b)
	if err != nil { return err }
	q.Add(j)
	return nil
}

func (jq *JourneyQueue) AssignCar(journey_element *list.Element, journey *models.Journey, cid int){
	journey.AssignCar(cid)
	jq.BySize[journey.GetSize()].Remove(journey_element)
}

func (q *JourneyQueue ) GetById(id int) (*list.Element, *models.Journey, error){
	el, ok := q.ById[id]
	if !ok {
		return nil, nil, errors.New("Not found")
	}
	return el, el.Value.(*models.Journey), nil
}