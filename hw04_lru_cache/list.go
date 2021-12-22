package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	size  int
}

func (l list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{v, nil, nil}

	if l.first == nil {
		l.first = newItem
	} else {
		newItem.Next = l.first
		l.first.Prev = newItem
		l.first = newItem
	}
	l.size++
	if l.last == nil {
		l.last = l.first
	}
	return l.first
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{v, nil, nil}

	if l.last == nil {
		l.last = newItem
	} else {
		newItem.Prev = l.last
		l.last.Next = newItem
		l.last = newItem
	}
	l.size++
	if l.first == nil {
		l.first = l.last
	}
	return l.last
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.first = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.last = i.Prev
	}
	i.Next = nil
	i.Prev = nil
	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
