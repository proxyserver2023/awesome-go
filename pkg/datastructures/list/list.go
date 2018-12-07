package list

type list struct {
	head interface{}
	tail PersistentList
}

func (l *list) Head() (interface{}, bool) {
	return l.head, true
}

func (l *list) Tail() (PersistentList, bool) {
	return l.tail, true
}

func (l *list) IsEmpty() bool {
	return false
}

func (l *list) Length() uint {
	curr := l
	length := uint(0)

	for {
		length++
		tail, _ := curr.Tail()

		if tail.IsEmpty() {
			return length
		}

		curr = tail.(*list)
	}
}

func (l *list) Add(head interface{}) PersistentList {
	return &list{head, l}
}

func (l *list) Insert(val interface{}, pos uint) (PersistentList, error) {
	if pos == 0 {
		return l.Add(val), nil
	}

	nl, err := l.tail.Insert(val, pos-1)
	if err != nil {
		return nil, err
	}

	return nl.Add(l.head), nil

}

func (l *list) Get(pos uint) (interface{}, bool) {
	if pos == 0 {
		return l.head, true
	}
	return l.tail.Get(pos - 1)
}

func (l *list) Remove(pos uint) (PersistentList, error) {
	if pos == 0 {
		nl, _ := l.Tail()
		return nl, nil
	}

	nl, err := l.tail.Remove(pos - 1)

	if err != nil {
		return nil, err
	}

	return &list{l.head, nl}, nil
}

func (l *list) Find(pred func(interface{}) bool) (interface{}, bool) {
	if pred(l.head) {
		return l.head, true
	}

	return l.tail.Find(pred)
}

func (l *list) FindIndex(pred func(interface{}) bool) int {
	curr := l
	idx := 0
	for {
		if pred(l.head) {
			return idx
		}

		tail, _ := curr.Tail()

		if tail.IsEmpty() {
			return -1
		}

		curr = tail.(*list)
		idx++
	}
}

func (l *list) Map(f func(interface{}) interface{}) []interface{} {
	return append(l.tail.Map(f), f(l.head))
}
