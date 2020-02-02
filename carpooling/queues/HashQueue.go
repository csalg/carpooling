/*
Both the CarQueue and JourneyQueue have some common elements. Following SOLID
principles, it makes sense to abstract these common algorithmic patterns into
a structure which is completely decoupled from the models so that they can be
tested independently and reused.

This data structure is designed to index objects with an ID and a size.
The following operations are implemented in constant time & space:
- Retrieval by id.
- Retrieval by size of the oldest element in a queue (FIFO).
- Deletion by id.
- Changing size by id.
The tradeoff is that iterating through its elements would be painfully slow
(no locality of reference in linked lists).
*/

package queues

import (
	"container/list"
	// "time"
	// "fmt"
	"errors"
	"reflect"
)

// WithSizeAndId is the interface that must be implemented by
// anything that is compatible with a HashQueue.
type WithSizeAndId interface {
	GetId() int
	GetSize() int
	SetSize(id int) error

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

func (q *HashQueue) Has(id int) bool {
	_, ok := q.ById[id]
	return ok
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


func (q *HashQueue) changeSize (el *list.Element, newSize int) (*list.Element, error) {
	if newSize < 0 || newSize > len(q.BySize)-1 {
		return nil, errors.New("Invalid new size.")
	}

	val := el.Value.(WithSizeAndId)
	id := val.GetId()
	previousSize := val.GetSize()
	err := val.SetSize(newSize)
	if err != nil {
		return nil, err
	}

	//fmt.Println("Before")
	//fmt.Print(q.BySize[previousSize].Front() == nil, " / ")
	//fmt.Print(q.BySize[newSize].Front() == nil, " / ")
	//fmt.Print(q.BySize[previousSize].Len(), " / ")
	//fmt.Println(q.BySize[newSize].Len())
	el_ := q.BySize[previousSize].Remove(el)
	q.BySize[newSize].PushBack(el_)
	q.ById[id] = q.BySize[newSize].Back()
	//fmt.Println("After")
	//fmt.Printf("Removed from %d, now in %d \n", previousSize, newSize)
	//fmt.Print(q.BySize[previousSize].Front() == nil, " / ")
	//fmt.Print(q.BySize[newSize].Front() == nil, " / ")
	//fmt.Print(q.BySize[previousSize].Len(), " / ")
	//fmt.Println(q.BySize[newSize].Len())
	return el, nil
}