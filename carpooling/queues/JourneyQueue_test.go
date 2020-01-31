package queues

import (
	"testing"
	"fmt"
)

func TestNewJourney(t *testing.T){

	for i := -10; i !=20; i++ {
		_, err := NewJourney(i,i)
		if (i <= 6 && i >= 1) {
			if err != nil  { t.Errorf(err.Error()) }
		} else if err == nil {
			t.Errorf("Journey with %d people was created!", i)
		}	
	}
}

func TestJourneyQueueAdd(t *testing.T){
	jq := NewJourneyQueue()

	j, err := NewJourney(1,7)
	err = jq.Add(j)
	if err == nil {
		t.Errorf("Queue inserted a journey outside range!")
	}

	j, err = NewJourney(1,2)
	err = jq.Add(j)
	if err != nil {
		t.Errorf("Failed to insert a valid journey.")
	}

	j, err = NewJourney(1,3)
	err = jq.Add(j)
	if err == nil {
		t.Errorf("Inserted a journey with a known duplicate key.")
	}

	if jq.Add(nil) == nil {
		t.Errorf("Queue inserted a null pointer!")
	}
}

func TestJourneyQueueDelete(t *testing.T){
	jq := NewJourneyQueue()

	for i := 1; i != 500; i++ {
		j, err := NewJourney(i,i%5+1)
		err = jq.Add(j)
		if err != nil {
			t.Errorf(fmt.Sprintf("Error adding journey. i=%d", i))
		}
	}

	for i := 1; i != 500; i++ {
		jq.Delete(i)
		_, exists := jq.ById[i]
		if exists { t.Errorf("Error deleting key from JourneyQueue")}
	}

	for i :=0; i!=7; i++{
		if jq.BySize[i].Front() != nil {
			t.Errorf("There are still some values left in ByPeople")
		}
	}
}


func TestGetOldestSmallerThan(t *testing.T){
	jq := NewJourneyQueue()
	j1, err1 := NewJourney(1,5)
	j2, err2 := NewJourney(2,1)
	if err1 != nil || err2 != nil {
		t.Errorf("Error constructing journeys.")
	}

	if 	jq.Add(j1) != nil || jq.Add(j2) != nil {
		t.Errorf("Error adding journeys to queue.")
	}

	for i:= 3; i != 100; i++ {
		j,_ := NewJourney(i, i%5+2)
		jq.Add(j)
	}

	old1, err4 := jq.GetOldestSmallerThan(6)
	if old1.Value != j1 || err4 != nil{
		t.Errorf("Error retrieving oldest journey.")
	}

	old2, err5 := jq.GetOldestSmallerThan(1)
	if old2.Value != j2 || err5 != nil{
		t.Errorf("Error retrieving journey smaller than or equal to 1.")
	}

	jq.Delete(2)

	old3, err6 := jq.GetOldestSmallerThan(1)
	if old3 != nil || err6 != nil{
		t.Errorf("Error retrieving a null pointer when there are no matching journeys.")
	}
}