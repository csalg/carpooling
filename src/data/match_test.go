package data

import (
	"testing"
	"github.com/csalg/carpooling/models"
	// "fmt"
)

func TestMatch(t *testing.T){
	jq := NewJourneyQueue()
	cq := NewCarQueue()

	c1, _ := models.NewCar(1,6)
	cq.Add(c1)
	j1, _ := models.NewJourney(1,4)
	j2, _ := models.NewJourney(2,2)
	jq.Add(j1)
	jq.Add(j2)

	err := Match(cq,jq)
	if err != nil {t.Errorf(err.Error())}

	err = Match(cq,jq)
	if err != nil {t.Errorf(err.Error())}

	mostAvailableSeats := cq.MostAvailableSeats()
	if mostAvailableSeats != 0 { t.Errorf("Problem with the value of most available seats. Expected: %d. Got: %d", 0, mostAvailableSeats)}

	for i := 1; i != len(cq.BySize); i++ {
		if cq.BySize[i].Front() != nil {
			t.Errorf("All queues greater than 0 should be empty, but i=%d isn't", i)
		}
	}

	// Sanity check: generate some cars and lots of users,
	// match should assign all cars and leave users waiting
	seats := 4
	for i:=2; i!=100; i++ {
		c,_ := models.NewCar(i,seats)
		cq.Add(c)
		if seats == 4 {
			seats = 6
		} else {
			seats = 4
		}
	}
	for i:=3; i!=2000; i++ {
		j,_ := models.NewJourney(i,i%6+1)
		jq.Add(j)
		Match(cq,jq)
	}
	for i :=1; i!=7; i++ {
		if cq.BySize[i].Front() != nil {
			t.Errorf("All cars should have been matched but some cars remain unassigned. i=%d", i)
		}
		if jq.BySize[i].Front() == nil {
			t.Errorf("There should be users waiting for rides. i=%d", i)
		}
	}
}