package _6

import "testing"

func TestQueueCache(t *testing.T) {
	lurCache := NewQueueCache(10)
	lurCache.Put(1, 2)
	t.Log(lurCache.Get(1))
}
