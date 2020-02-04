// This is the specialized version of a HashQueue for cars, with additional 
// logic wrapping the Car json deserializer, and a matching procedure between
// a CarQueue and a JourneyQueue.
package queues

import (
	"errors"
	"container/list"
	"github.com/csalg/carpooling/models"
	"io"
	// "fmt"
)

type carQueue struct {
	HashQueue
}

// NewCarQueue is the constructor for carQueue, which is kept private
// to prevent the client from not initializing the map and getting nil
// pointer errors
func NewCarQueue()*carQueue {
	jq := new(carQueue)
	jq.ById = make(map[int]*list.Element)
	return jq
}

// This is just sugar for changeSize.
func (q *carQueue) Move(el *list.Element, new_seats_available int) error {
	el, err := q.changeSize(el,new_seats_available)
	if err != nil {return err}
	err = el.Value.(*models.Car).SetSeatsAvailable(new_seats_available)
	return err
}


// MakeFromJsonRequest first calls the BodyToCars decoder and
// if that succeeds overwrites the CarQueue with the new ones.
func (q *carQueue) MakeFromJsonRequest(b io.ReadCloser)error{

	cars, err := models.BodyToCars(b)
	if err != nil { return err }
	q = NewCarQueue()

	for _, car := range *cars {
		q.Add(&car)
	}

	return nil
}

// GetCarLargerThan returns a car larger than or equal to val
func (q *carQueue) GetCarLargerThan(val int) (*list.Element, *models.Car, error) {
	for i := val; i <= 6; i++ {
		if q.BySize[i].Front() != nil { 
			c, ok := q.BySize[i].Front().Value.(*models.Car)
			if ok { return q.BySize[i].Front(), c, nil }
		}
	}
	return nil, nil, errors.New("Car not found!")
}

// MaxAvailable finds the car with the most seats available 
// and returns the amount.
func (q *carQueue) MostAvailableSeats() int {
	for i := len(q.BySize)-1; i != -1; i-- {
		if q.BySize[i].Front() != nil { 
			return i
		}
	}
	return 0
}

func (q *carQueue ) GetById(id int) (*list.Element, *models.Car, error){
	el, ok := q.ById[id]
	if !ok {
		return nil, nil, errors.New("Not found")
	}
	return el, el.Value.(*models.Car), nil
}