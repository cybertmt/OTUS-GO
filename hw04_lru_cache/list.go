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

// Len возвращает длину списк.
func (l *list) Len() int {
	return l.length
}

// Front возвращает первый элемент списка.
func (l *list) Front() *ListItem {
	return l.front
}

// Back возвращает последний элемент списка.
func (l *list) Back() *ListItem {
	return l.back
}

// PushFront добавляет элемент в начало списка.
func (l *list) PushFront(v interface{}) *ListItem {
	// Создаем первый элемент, если список пустой.
	if l.length == 0 {
		l.front = &ListItem{v, nil, nil}
		// Корректируем длину для объекта списка.
		l.length++
		// Элемент в списке единственный, уравниваем первый и последний элементы.
		l.back = l.front
		return l.front
	}
	// Создаем элемент и привязываем в конец списка и возвращаем элемент.
	l.front.Prev = &ListItem{v, nil, l.front}
	l.length++
	l.front = l.front.Prev
	return l.front
}

// PushBack добавляет элемент в конец списка.
func (l *list) PushBack(v interface{}) *ListItem {
	// Создаем первый элемент, если список пустой и возвращаем элемент.
	if l.length == 0 {
		l.back = &ListItem{v, nil, nil}
		// Корректируем длину для объекта списка.
		l.length++
		l.front = l.back
		return l.back
	}
	// Создаем элемент и привязываем в конец списка и возвращаем элемент.
	l.back.Next = &ListItem{v, l.back, nil}
	// Корректируем длину для объекта списка.
	l.length++
	// Элемент в списке единственный, уравниваем первый и последний элементы.
	l.back = l.back.Next
	return l.back
}

// Remove удаляет элемент.
func (l *list) Remove(i *ListItem) {
	// Если список пустой сразу выходим.
	if i == nil || l.length == 0 {
		return
	}
	// Если в списке всего один элемент, обнуляем длину
	// и первый и последний элементы переводим в nil.
	if l.length == 1 {
		l.front, l.back = nil, nil
		l.length--
		return
	}

	// Если удаляемый элемент первый в списке.
	if i == l.front {
		l.front = i.Next
		l.front.Prev = nil
		l.length--
		return
	}
	// Если удаляемый элемент последний в списке.
	if i == l.back {
		l.back = i.Prev
		l.back.Next = nil
		l.length--
		return
	}
	// Удаляем элемент в середине списка.
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	l.length--
}

// MoveToFront перемещает элемент в начало списка.
func (l *list) MoveToFront(i *ListItem) {
	// Первый элемент двигать вперед нет необходимости, сразу выходим.
	if i == nil || i == l.front {
		return
	}
	// Удаляем перемещаемый элемент из списка.
	l.Remove(i)
	// Перемещаем этот элемент в начало списка.
	l.front.Prev = i
	i.Next = l.front
	l.front = i
	l.front.Prev = nil
	l.length++
}
