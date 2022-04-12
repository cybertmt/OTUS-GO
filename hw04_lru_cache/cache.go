package main

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

// cacheItemMap дополнительный словарь [указатель на эелемент]{ключ, значение элемента}.
var cacheItemMap = make(map[*ListItem]cacheItem)

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Clear полностью обнуляет кэш.
func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

// Set добавляет элемент в кэш.
func (l *lruCache) Set(key Key, value interface{}) bool {
	// Блокируем кэш для изменения/чтения.
	l.mu.Lock()
	defer l.mu.Unlock()
	// Проверяем
	i, ok := l.items[key]
	// Если ключ уже присутвует в кэше, обновляем значение элемента в списке и словаре cacheItemMap .
	// Переносим элемент в начало списка. Возвращаем true.
	if ok {
		l.items[key].Value = value
		cacheItemMap[l.items[key]] = cacheItem{key, value}
		l.queue.MoveToFront(i)
		return true
	}
	if l.queue.Len() == l.capacity {
		keyToRemove := cacheItemMap[l.queue.Back()].key
		delete(l.items, keyToRemove)
		l.queue.Remove(l.queue.Back())
	}
	// Если ключ отсутсвует в кэше, добавляем элемент в начало списка.
	// Добавляем элемент в словарь cacheItemMap. Возвращаем false.
	l.items[key] = l.queue.PushFront(value)
	cacheItemMap[l.items[key]] = cacheItem{key, value}
	return false
}

// Get получает значение элемента из кэша.
func (l *lruCache) Get(key Key) (interface{}, bool) {
	// Блокируем кэш для изменения/чтения.
	l.mu.Lock()
	defer l.mu.Unlock()
	// Если элемент присутвует в кэше, передвигаем элемент в начало списка.
	// Возвращаем значение элемента и true.
	i, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(i)
		return i.Value, true
	}
	// Если элемент отсутвует в кэше, возвращаем nil и false.
	return nil, false
}
