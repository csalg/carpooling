// This is the data access layer, which consists of specialised versions of the HashQueue (carQueue and journeyQueue)
// and the Match and Dropoff functions.
package data

import (
	"container/list"
	"errors"
	"github.com/csalg/carpooling/src/backing"
	"github.com/csalg/carpooling/src/models"
	"io"
)

type carQueueType struct {
	backing.HashQueue
}


// NewCarQueue is the constructor for carQueueType, which is kept private
// to prevent the client from not initializing the map and getting nil
// pointer errors
func NewCarQueue()*carQueueType {
	carQueue := new(carQueueType)
	carQueue.ById = make(map[int]*list.Element)
	return carQueue
}

// This is just sugar for changeSize.
func (carQueue *carQueueType) Move(element *list.Element, newSeatsAvailable int) error {
	element, err := carQueue.ChangeSize(element, newSeatsAvailable)
	if err != nil {return err}
	err = element.Value.(*models.Car).SetSeatsAvailable(newSeatsAvailable)
	return err
}


// MakeFromJsonRequest first calls the BodyToCars decoder and
// if that succeeds overwrites the CarQueue with the new ones.
func (carQueue *carQueueType) MakeFromJsonRequest(b io.ReadCloser)error{

	carsArray, err := models.BodyToCars(b)
	if err != nil { return err }
	*carQueue = *NewCarQueue()

	for _, car := range *carsArray {
		carQueue.Add(&car)
	}

	return nil
}


// GetCarLargerThan returns a car larger than or equal to val
func (carQueue *carQueueType) GetCarLargerThan(val int) (*list.Element, *models.Car, error) {
	for i := val; i <= 6; i++ {
		if carQueue.BySize[i].Front() != nil {
			c, ok := carQueue.BySize[i].Front().Value.(*models.Car)
			if ok { return carQueue.BySize[i].Front(), c, nil }
		}
	}
	return nil, nil, errors.New("Car not found!")
}


// MaxAvailable finds the car with the most seats available 
// and returns the amount.
func (carQueue *carQueueType) MostAvailableSeats() int {
	for i := len(carQueue.BySize)-1; i != -1; i-- {
		if carQueue.BySize[i].Front() != nil {
			return i
		}
	}
	return 0
}


// GetById fetches the list element and the car payload it contains or an error if not found
func (carQueue *carQueueType) GetById(id int) (*list.Element, *models.Car, error){
	element, ok := carQueue.ById[id]
	if !ok {
		return nil, nil, errors.New("Not found")
	}
	return element, element.Value.(*models.Car), nil
}


// GetCarJsonById returns a CarJson object or an error if the ID is not in the structure.
func (carQueue *carQueueType) GetCarJsonById(id int) (*models.CarJson, error){
	_, car, err := carQueue.GetById(id)
	if err != nil {
		return nil, err
	}
	return car.ToCarJson(), nil
}