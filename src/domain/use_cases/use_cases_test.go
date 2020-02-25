package use_cases

import (
	"github.com/csalg/carpooling/src/domain/entities"
	"github.com/csalg/carpooling/src/persistence"
	"testing"
)

func TestMatch(t *testing.T){
	journeyQueue := persistence.NewJourneyRepository()
	carQueue := persistence.NewCarRepository()

	car1, _ := entities.NewCar(1,6)
	carQueue.Add(car1)
	journey1, _ := entities.NewJourney(1,4)
	journey2, _ := entities.NewJourney(2,2)
	journeyQueue.Add(journey1)
	journeyQueue.Add(journey2)

	err := Match(carQueue, journeyQueue)
	if err != nil {t.Errorf(err.Error())}
	_, journey, err := journeyQueue.GetById(1)
	if journey.Car != 1 {
		t.Errorf("Car was not asigned properly. Expected 1, got %d", journey.Car)
	}

	err = Match(carQueue, journeyQueue)
	if err != nil {t.Errorf(err.Error())}

	mostAvailableSeats := carQueue.MostAvailableSeats()
	if mostAvailableSeats != 0 { t.Errorf("Problem with the value of most available seats. Expected: %d. Got: %d", 0, mostAvailableSeats)}

	for i := 1; i != len(carQueue.BySize); i++ {
		if carQueue.BySize[i].Front() != nil {
			t.Errorf("All queues greater than 0 should be empty, but i=%d isn't", i)
		}
	}

	// Sanity check: generate some cars and lots of users,
	// match should assign all cars and leave users waiting
	seats := 4
	for i:=2; i!=100; i++ {
		car,_ := entities.NewCar(i,seats)
		carQueue.Add(car)
		if seats == 4 {
			seats = 6
		} else {
			seats = 4
		}
	}
	for i:=3; i!=2000; i++ {
		j,_ := entities.NewJourney(i,i%6+1)
		journeyQueue.Add(j)
		Match(carQueue, journeyQueue)
	}
	for i :=1; i!=7; i++ {
		if carQueue.BySize[i].Front() != nil {
			t.Errorf("All cars should have been matched but some cars remain unassigned. i=%d", i)
		}
		if journeyQueue.BySize[i].Front() == nil {
			t.Errorf("There should be users waiting for rides. i=%d", i)
		}
	}
}