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

type cacheItem struct {
	key   Key
	value interface{}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if element, exists := l.items[key]; exists {
		l.queue.MoveToFront(element)
		element.Value.(*cacheItem).value = value
		return true
	}

	if l.queue.Len() == l.capacity {
		if element := l.queue.Back(); element != nil {
			l.queue.Remove(element)
			delete(l.items, element.Value.(*cacheItem).key)
		}
	}

	item := &cacheItem{
		key:   key,
		value: value,
	}

	element := l.queue.PushFront(item)
	l.items[item.key] = element
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	element, exists := l.items[key]
	if !exists {
		return nil, false
	}
	l.queue.MoveToFront(element)
	return element.Value.(*cacheItem).value, true
}

func (l *lruCache) Clear() {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
