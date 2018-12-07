package list

import "errors"

type emptyList struct{}

var (
	Empty        PersistentList = &emptyList{}
	ErrEmptyList                = errors.New("Empty List")
)

func (e *emptyList) Head() (interface{}, bool) {
	return nil, false
}

func (e *emptyList) Tail() (PersistentList, bool) {
	return nil, false
}

func (e *emptyList) IsEmpty() bool {
	return true
}

func (e *emptyList) Length() uint {
	return 0
}

func (e *emptyList) Add(head interface{}) PersistentList {
	return &list{head, e}
}

func (e *emptyList) Insert(val interface{}, pos uint) (PersistentList, error) {
	if pos == 0 {
		return e.Add(val), nil
	}
	return nil, ErrEmptyList
}

func (e *emptyList) Get(pos uint) (interface{}, bool) {
	return nil, false
}

func (e *emptyList) Remove(pos uint) (PersistentList, error) {
	return nil, ErrEmptyList
}

func (e *emptyList) Find(func(interface{}) bool) (interface{}, bool) {
	return nil, false
}

func (e *emptyList) FindIndex(func(interface{}) bool) int {
	return -1
}

func (e *emptyList) Map(func(interface{}) interface{}) []interface{} {
	return nil
}
