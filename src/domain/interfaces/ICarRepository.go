package interfaces

import (
	"container/list"
	"github.com/csalg/carpooling/src/domain/entities"
	"io"
)

type ICarRepository interface {
	IBasicDataAccessOperations
	Move(element *list.Element, newSeatsAvailable int) error
	MakeFromJsonRequest(b io.ReadCloser)error
	GetCarLargerThan(val int) (*list.Element, *entities.Car, error)
	MostAvailableSeats() int
	GetById(id int) (*list.Element, *entities.Car, error)
	GetCarJsonById(id int) (*entities.CarJson, error)
}
