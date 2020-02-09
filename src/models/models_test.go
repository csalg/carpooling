package models

import (
	"testing"
)


func TestNewJourney(t *testing.T){
	for i := -10; i !=20; i++ {
		_, err := NewJourney(i,i)
		if i <= 6 && i >= 1 {
			if err != nil  { t.Errorf(err.Error()) }
		} else if err == nil {
			t.Errorf("Journey with %d people was created!", i)
		}	
	}
}


func TestNewCar(t *testing.T){
	for i := -10; i !=20; i++ {
		_, err := NewCar(1,i)
		if i == 4 || i ==6 {
			if err != nil  { t.Errorf(err.Error()) }
		} else if err == nil {
			t.Errorf("Car with %d seats was created!", i)
		}	
	}
}
