// A queue that allows CRUD operations by ID in constant time
// thanks to a hashmap. It is a similar idea to a LRU cache but with more queues
// (one per possible size).
package persistence

import (
	"github.com/csalg/carpooling/src/domain/interfaces"
	"container/list"
	// "time"
	// "fmt"
	"errors"
	"reflect"

)

// WithSizeAndId is the interface that must be implemented by
// anything that is compatible with a HashQueue.

type HashQueue struct {
	ById map[int]*list.Element
	BySize [7]list.List
}

// Has is true if an id corresponds to an element in the structure and false otherwise.
func (q *HashQueue) Has(id int) bool {
	_, ok := q.ById[id]
	return ok
}


// Add adds an element into the structure.
func (q *HashQueue) Add(e interfaces.WithSizeAndId) error {
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


// Delete eliminates an element from both the map and the queue. Returns an error if not found.
func (q *HashQueue) Delete (id int) error {
	el, ok := q.ById[id]
	if !ok { return errors.New("Id not found")}

	q.BySize[el.Value.(interfaces.WithSizeAndId).GetSize()].Remove(el)
	delete(q.ById, id)
	return nil
}


// ChangeSize moves elements from one queue to another.
func (q *HashQueue) ChangeSize(el *list.Element, newSize int) (*list.Element, error) {
	if newSize < 0 || newSize > len(q.BySize)-1 {
		return nil, errors.New("Invalid new size.")
	}

	val := el.Value.(interfaces.WithSizeAndId)
	id := val.GetId()
	previousSize := val.GetSize()
	err := val.SetSize(newSize)
	if err != nil {
		return nil, err
	}

	el_ := q.BySize[previousSize].Remove(el)
	q.BySize[newSize].PushBack(el_)
	q.ById[id] = q.BySize[newSize].Back()
	return el, nil
}
