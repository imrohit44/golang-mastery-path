package main

import (
	"container/list"
	"fmt"
)

type LRUCache[K comparable, V any] struct {
	capacity int
	items    map[K]*list.Element
	evictList *list.List
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
		evictList: list.New(),
	}
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	// If it exists, update it and move to front
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.(*entry[K, V]).value = value
		return
	}

	// If at capacity, evict the oldest item (back of the list)
	if c.evictList.Len() >= c.capacity {
		evict := c.evictList.Back()
		if evict != nil {
			c.evictList.Remove(evict)
			delete(c.items, evict.Value.(*entry[K, V]).key)
		}
	}

	// Add new item to front
	ent := &entry[K, V]{key, value}
	element := c.evictList.PushFront(ent)
	c.items[key] = element
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		return ent.Value.(*entry[K, V]).value, true
	}
	var zero V
	return zero, false
}

func main() {
	cache := NewLRUCache[string, int](2)

	cache.Put("A", 100)
	cache.Put("B", 200)
	
	cache.Get("A") // Access A, making B the least recently used
	
	cache.Put("C", 300) // This will evict B because capacity is 2

	if val, ok := cache.Get("B"); !ok {
		fmt.Println("B was successfully evicted.")
	} else {
		fmt.Println("B:", val)
	}
	
	val, _ := cache.Get("C")
	fmt.Println("C:", val)
}