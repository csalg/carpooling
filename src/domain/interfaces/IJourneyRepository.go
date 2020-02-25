package interfaces

import (
	"container/list"
	"github.com/csalg/carpooling/src/domain/entities"
	"io"
)

type IJourneyRepository interface {
	IBasicDataAccessOperations
	GetOldestSmallerThan(val int)(*list.Element, *entities.Journey, error)
	AddFromJsonRequest(body io.ReadCloser) error
	AssignCar(journeyElement *list.Element, journey *entities.Journey, carId int)
	GetById(id int) (*list.Element, *entities.Journey, error)
}



