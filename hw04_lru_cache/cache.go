package hw04lrucache

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
	mtx      sync.Mutex
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mtx.Lock()
	valueInMap := l.items[key]
	if valueInMap != nil {
		if valueInMap.Value == value {
			l.queue.MoveToFront(valueInMap)
		} else {
			l.queue.Remove(valueInMap)
			l.items[key] = l.queue.PushFront(value)
		}
		l.mtx.Unlock()
		return true
	}
	l.items[key] = l.queue.PushFront(value)
	if l.queue.Len() > l.capacity {
		lastVale := l.queue.Back()
		l.queue.Remove(lastVale)
		for key, value := range l.items {
			if value == lastVale {
				delete(l.items, key)
			}
		}
	}
	l.mtx.Unlock()
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mtx.Lock()
	valueInMap := l.items[key]
	if valueInMap != nil {
		l.queue.MoveToFront(valueInMap)
		l.mtx.Unlock()
		return valueInMap.Value, true
	}
	l.mtx.Unlock()
	return nil, false
}

func (l *lruCache) Clear() {
	l.mtx.Lock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
	l.mtx.Unlock()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
