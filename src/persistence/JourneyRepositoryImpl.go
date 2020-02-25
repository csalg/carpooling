package persistence

import (
	"container/list"
	"errors"
	"github.com/csalg/carpooling/src/domain/entities"
	"io"
	"time"
)

type journeyQueueType struct {
	HashQueue
}

func NewJourneyRepository()*journeyQueueType {
	journeyQueue := new(journeyQueueType)
	journeyQueue.ById = make(map[int]*list.Element)
	return journeyQueue
}

func (journeyQueue *journeyQueueType) GetOldestSmallerThan(val int)(*list.Element, *entities.Journey, error){
	if val > len(journeyQueue.BySize) || val < 1 {
		return nil, nil, errors.New("Value is outside range") 
	}
	
	timestamp := time.Now().UnixNano()
	element   := new(list.Element)
	journey   := new(entities.Journey)

	element, journey = nil, nil

	for i := 1; i != val+1; i++ {
		maybeElement := journeyQueue.BySize[i].Front()
		if maybeElement != nil {
			maybeCandidate := journeyQueue.BySize[i].Front().Value.(*entities.Journey)
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
	journey, err := entities.NewJourneyFromBody(body)
	if err != nil { return err }
	journeyQueue.Add(journey)
	return nil
}

func (journeyQueue *journeyQueueType) AssignCar(journeyElement *list.Element, journey *entities.Journey, carId int){
	journey.AssignCar(carId)
	journeyQueue.BySize[journey.GetSize()].Remove(journeyElement)
}

func (journeyQueue *journeyQueueType) GetById(id int) (*list.Element, *entities.Journey, error){
	element, ok := journeyQueue.ById[id]
	if !ok {
		return nil, nil, errors.New("Not found")
	}
	return element, element.Value.(*entities.Journey), nil
}