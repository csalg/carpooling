package queues

import (
	// "time"
	// "fmt"
	"errors"
	"container/list"
	"reflect"
)

type WithSizeAndId interface {
	GetId() int
	GetSize() int
}

type HashQueue struct {
	ById map[int]*list.Element
	BySize [7]list.List
}

func NewHashQueue()*HashQueue {
	q := new(HashQueue)
	q.ById = make(map[int]*list.Element)
	return q
}

func (q *HashQueue) Add(e WithSizeAndId) error {

	if e == nil || reflect.ValueOf(e).IsNil() { 
		return errors.New("Cannot add a null pointer.")
	}

	_, exists := q.ById[e.GetId()]
	if exists {
		return errors.New("Key already exists")
	}

	q.BySize[e.GetSize()].PushBack(e)
	q.ById[e.GetId()] = q.BySize[e.GetSize()].Back()

	return nil
}

func (q *HashQueue) Delete (id int) error {
	el, ok := q.ById[id]
	if !ok { return errors.New("Id not found")}
	q.BySize[el.Value.(WithSizeAndId).GetSize()].Remove(el)
	delete(q.ById, id)
	return nil
}