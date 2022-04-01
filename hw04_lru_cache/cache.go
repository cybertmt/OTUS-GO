package main

import "sync"

type Key string

var mu sync.Mutex

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

var cacheItemMap = make(map[*ListItem]cacheItem)

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	mu.Lock()
	defer mu.Unlock()
	i, ok := l.items[key]
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
	l.items[key] = l.queue.PushFront(value)
	cacheItemMap[l.items[key]] = cacheItem{key, value}
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	mu.Lock()
	defer mu.Unlock()
	i, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(i)
		return i.Value, true
	}
	return nil, false
}
