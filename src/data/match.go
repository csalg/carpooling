package data

import (
	"errors"
)

// Match all possible journeys to available cars in 
// journey arrival order
func Match(cq *carQueueType, jq *journeyQueueType) error {

		mostAvailableSeats := cq.MostAvailableSeats()
		if mostAvailableSeats == 0 {
			return errors.New("No more cars available") 
		}

		journeyElement, journey, err := jq.GetOldestSmallerThan(mostAvailableSeats)
		if err != nil { 
			return err
		}

		carElement, car, err := cq.GetCarLargerThan(journey.GetSize())
		if err != nil { 
			return err
			}	

		carId := car.GetId()
		err  = cq.Move(carElement, car.GetSize() - journey.GetSize())
		if err != nil { 
			return err
		}
		
		jq.AssignCar(journeyElement, journey, carId)
	return nil
}

// Dropoff removes a journey from the queue and updates the car's
// capacity accordingly if it was travelling.
func Dropoff(carQ *carQueueType, journeyQueue *journeyQueueType, jid int) error {
	defer journeyQueue.Delete(jid)

	_, journey, err := journeyQueue.GetById(jid)
	if err != nil {
		return err
	}

	if journey.IsTravelling() {
		carElement, car, err := carQ.GetById(journey.Car)
		if err != nil {
			return err
		} else {
			carQ.Move(carElement, car.GetSize()+journey.GetSize())
			defer carQ.Delete(car.GetId())
		}
	}
	return nil
}
