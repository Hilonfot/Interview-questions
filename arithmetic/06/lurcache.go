package _6

import (
	"sync"
)

type IQueue interface {
	Get(k interface{}) interface{}
	Put(k interface{}, v interface{})
}

type QueueData struct {
	key   interface{}
	value interface{}
}

func NewQueueData(k, v interface{}) *QueueData {
	return &QueueData{key: k, value: v}
}

type QueueCache struct {
	Data   []*QueueData
	cap    int // 容量
	len    int // 长度
	rwLock sync.RWMutex
}

func NewQueueCache(cap int) *QueueCache {
	queue := new(QueueCache)
	queue.cap = cap
	queue.Data = make([]*QueueData, 0, cap)
	return queue
}

func (q *QueueCache) Get(k interface{}) (val interface{}) {
	q.rwLock.RLock()

	if ks, isOk := q.IsExist(k); isOk {
		val = q.Data[ks.(int)].value
		q.len--
	}
	q.rwLock.RUnlock()
	return -1
}

func (q *QueueCache) Put(k interface{}, v interface{}) {
	if q.IsEmpty() {
		q.Data = q.Data[:q.len]
	}
	q.rwLock.Lock()
	q.Data = append(q.Data, NewQueueData(k, v))
	q.rwLock.Unlock()
	q.len++
}

func (q *QueueCache) IsExist(k interface{}) (interface{}, bool) {
	for ks, val := range q.Data {
		if val.key == k {
			return ks, true
		}
	}
	return nil, false
}

func (q *QueueCache) IsEmpty() bool {
	return q.len == q.cap
}
