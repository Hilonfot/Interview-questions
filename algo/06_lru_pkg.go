package main

import (
	"container/list"
	"sync"
)

// 项目中使用，双向链表采用标准库
type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.Element
	OnEvicted func(key string, value Value)
	mx    sync.Mutex
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry) // (*entry)  类型断言，作用于interface{}
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item
func (c *Cache) RemoveOldest() {
	if ele := c.ll.Back(); ele != nil {
		c.DeleteEle(ele)
	}
}

// Remove remove the item
func (c *Cache) Remove(key string) {
	if ele, ok := c.cache[key]; ele != nil && ok {
		c.DeleteEle(ele)
	}
}

// DeleteEle delete ele *Element
func (c *Cache) DeleteEle(ele *list.Element) {
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add adds a value to the cache.
func (c *Cache) Add(key string, value Value) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		kv.value = value
		return
	}
	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele
	c.nbytes += int64(len(key)) + int64(value.Len())
	// maxBytes == 0 时， 不限制字节大小
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// Len the number of cache entries
func (c *Cache) Len() int {
	return c.ll.Len()
}