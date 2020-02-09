package data

import (
	"container/list"
	"errors"
	"github.com/csalg/carpooling/src/backing"
	"github.com/csalg/carpooling/src/models"
	"io"
	"time"
)

type journeyQueueType struct {
	backing.HashQueue
}

func NewJourneyQueue()*journeyQueueType {
	journeyQueue := new(journeyQueueType)
	journeyQueue.ById = make(map[int]*list.Element)
	return journeyQueue
}

func (journeyQueue *journeyQueueType) GetOldestSmallerThan(val int)(*list.Element, *models.Journey, error){
	if val > len(journeyQueue.BySize) || val < 1 {
		return nil, nil, errors.New("Value is outside range") 
	}
	
	timestamp := time.Now().UnixNano()
	element   := new(list.Element)
	journey   := new(models.Journey)

	element, journey = nil, nil

	for i := 1; i != val+1; i++ {
		maybeElement := journeyQueue.BySize[i].Front()
		if maybeElement != nil {
			maybeCandidate := journeyQueue.BySize[i].Front().Value.(*models.Journey)
			if maybeCandidate.GetTimestamp() < timestamp {
				element   = maybeElement
				journey   = maybeCandidate
				timestamp = maybeCandidate.GetTimestamp()
			}
		}
	}
	if element == nil {
		return nil, nil, errors.New("No journey found!")
	}
	return element, journey, nil
}

func (journeyQueue *journeyQueueType) AddFromJsonRequest(body io.ReadCloser) error {
	journey, err := models.NewJourneyFromBody(body)
	if err != nil { return err }
	journeyQueue.Add(journey)
	return nil
}

func (journeyQueue *journeyQueueType) AssignCar(journeyElement *list.Element, journey *models.Journey, carId int){
	journey.AssignCar(carId)
	journeyQueue.BySize[journey.GetSize()].Remove(journeyElement)
}

func (journeyQueue *journeyQueueType) GetById(id int) (*list.Element, *models.Journey, error){
	element, ok := journeyQueue.ById[id]
	if !ok {
		return nil, nil, errors.New("Not found")
	}
	return element, element.Value.(*models.Journey), nil
}