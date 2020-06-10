package _6

import "testing"

func TestQueueCache(t *testing.T) {
	lurCache := NewQueueCache(5)
	lurCache.Put(1, 1)
	lurCache.Put(2, 2)
	lurCache.Put(3, 3)
	lurCache.Put(4, 4)
	lurCache.Put(5, 5)
	lurCache.Put(6, 6)
	lurCache.Put(7, 7)
	t.Log(lurCache.Get(1))
}
