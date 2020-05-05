package models

// Model defines base interface
type Model interface {
	All() interface{}
	Create([]byte) (interface{}, error)
	Find(int) interface{}
	Update(int, []byte) (interface{}, error)
	Delete(uint)
}
