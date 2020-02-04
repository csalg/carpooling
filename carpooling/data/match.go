package queues

import (
	"errors"
)

// Match all possible journeys to available cars in 
// journey arrival order
func Match(cq *carQueue, jq *JourneyQueue) error {

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

		cid := car.GetId()
		err  = cq.Move(carElement, car.GetSize() - journey.GetSize())
		if err != nil { 
			return err
		}
		
		jq.AssignCar(journeyElement, journey, cid)
	return nil
}

// Dropoff removes a journey from the queue and updates the car's
// capacity accordingly if it was travelling.
func Dropoff(cq *carQueue, jq *JourneyQueue, jid int) error {
	defer jq.Delete(jid)

	_, journey, err := jq.GetById(jid)
	if err != nil {
		return err
	}

	if journey.IsTravelling() {
		carElement, car, err := cq.GetById(journey.Car)
		if err != nil {
			return err
		} else {
			cq.Move(carElement, car.GetSize()+journey.GetSize())
			defer cq.Delete(car.GetId())
		}
	}
	return nil
}
