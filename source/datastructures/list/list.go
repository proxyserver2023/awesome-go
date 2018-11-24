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
