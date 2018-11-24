package list

import "errors"

type emptyList struct{}

var (
	Empty PersistentList = &emptyList{}

	ErrEmptyList = errors.New("Empty List")
)

// PersistentList What is this? :TODO:
type PersistentList interface {
	Head() (interface{}, bool)
	Tail() (PersistentList, bool)
	IsEmpty() bool
	Length() uint
	Add(head interface{}) PersistentList
	Insert(val interface{}, pos uint) (PersistentList, error)
	Get(pos uint) (interface{}, error)
	Remove(pos uint) (PersistentList, error)
	Find(func(interface{}) bool) (interface{}, bool)
	FindIndex(func(interface{}) bool) int
	Map(func(interface{}) interface{}) []interface{}
}
