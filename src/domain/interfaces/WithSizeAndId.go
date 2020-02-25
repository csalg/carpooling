package interfaces

type WithSizeAndId interface {
	GetId() int
	GetSize() int
	SetSize(id int) error

}