package hw04lrucache

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
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	valueInMap := l.items[key]
	if valueInMap != nil {
		if valueInMap.Value == value {
			l.queue.MoveToFront(valueInMap)
		} else {
			l.queue.Remove(valueInMap)
			l.items[key] = l.queue.PushFront(value)
		}
		return true
	} else {
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
		return false
	}

}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	valueInMap := l.items[key]
	if valueInMap != nil {
		l.queue.MoveToFront(valueInMap)
		return valueInMap.Value, true
	} else {
		return nil, false
	}
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
