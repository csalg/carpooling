package interfaces

import "container/list"

type IBasicDataAccessOperations interface {
	Has(id int) bool
	Add(e WithSizeAndId) error
	Delete (id int) error
	ChangeSize(el *list.Element, newSize int) (*list.Element, error)
}
