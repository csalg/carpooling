package use_cases

import (
	"errors"
	"github.com/csalg/carpooling/src/domain/interfaces"
)

// Match all possible journeys to available cars in 
// journey arrival order
func Match(carRepository interfaces.ICarRepository, journeyRepository interfaces.IJourneyRepository) error {

	mostAvailableSeats := carRepository.MostAvailableSeats()
	if mostAvailableSeats == 0 {
		return errors.New("No more cars available")
	}

	journeyElement, journey, err := journeyRepository.GetOldestSmallerThan(mostAvailableSeats)
	if err != nil {
		return err
	}

	carElement, car, err := carRepository.GetCarLargerThan(journey.GetSize())
	if err != nil {
		return err
		}

	carId := car.GetId()
	err  = carRepository.Move(carElement, car.GetSize() - journey.GetSize())
	if err != nil {
		return err
	}

	journeyRepository.AssignCar(journeyElement, journey, carId)
	return nil
}