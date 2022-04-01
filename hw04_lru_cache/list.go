package main

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
	Prev  *ListItem
	Next  *ListItem
}

type list struct {
	front  *ListItem
	back   *ListItem
	length int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.length == 0 {
		l.front = &ListItem{v, nil, nil}
		l.length++
		l.back = l.front
		return l.front
	}
	l.front.Prev = &ListItem{v, nil, l.front}
	l.length++
	l.front = l.front.Prev
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.length == 0 {
		l.back = &ListItem{v, nil, nil}
		l.length++
		l.front = l.back
		return l.back
	}
	l.back.Next = &ListItem{v, l.back, nil}
	l.length++
	l.back = l.back.Next
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i == nil || l.length == 0 {
		return
	}
	if l.length == 1 {
		l.front, l.back = nil, nil
		l.length--
		return
	}
	if i == l.front {
		l.front = i.Next
		l.front.Prev = nil
		l.length--
		return
	}
	if i == l.back {
		l.back = i.Prev
		l.back.Next = nil
		l.length--
		return
	}
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || i == l.front {
		return
	}
	l.Remove(i)
	l.front.Prev = i
	i.Next = l.front
	l.front = i
	l.front.Prev = nil
	l.length++
}
