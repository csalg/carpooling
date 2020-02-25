package use_cases

import (
	"github.com/csalg/carpooling/src/domain/interfaces"
)

// Dropoff removes a journey from the queue and updates the car's
// capacity accordingly if it was travelling.
func Dropoff(carRepository interfaces.ICarRepository, journeyRepository interfaces.IJourneyRepository, jid int) error {
	defer journeyRepository.Delete(jid)

	_, journey, err := journeyRepository.GetById(jid)
	if err != nil {
		return err
	}

	if journey.IsTravelling() {
		carElement, car, err := carRepository.GetById(journey.Car)
		if err != nil {
			return err
		} else {
			carRepository.Move(carElement, car.GetSize()+journey.GetSize())
			defer carRepository.Delete(car.GetId())
		}
	}
	return nil
}
