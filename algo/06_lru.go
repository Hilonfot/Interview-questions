package main

import (
	"fmt"
	"sync"
)

// 使用双链表实现lru

// 定义一个节点
type Node struct {
	Prev  *Node // 前节点
	Next  *Node // 后节点
	Key   int
	Value int // 存储值
}

// 实例化一个节点
func NewNode(k, v int) *Node {
	return &Node{
		Key:   k,
		Value: v,
	}
}

// 定义一个LRUCache结构体
type LRUCache struct {
	cache map[int]*Node // 设置一个map，用于存储值
	head  *Node         // 头节点
	tail  *Node         // 尾节点
	len   int           // 链表长度
	cap   int           // 链表容量
	mx    sync.Mutex
}

type LRUCacher interface {
	String() string     // 打印
	addHead(n *Node)    // 添加头部节点
	remove(n *Node)     // 移除节点
	Get(key int) int    // 获取
	Put(key, value int) // 存储
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		cache: make(map[int]*Node),
		cap:   capacity,
	}
}
func (l *LRUCache) String() string {
	str := "lru节点信息:"
	p := l.head
	for p != nil {
		str += fmt.Sprintf("k:%d,v:%d", p.Key, p.Value)
		p = p.Next
		if p != nil {
			str += "=>"
		}
	}
	return str
}

func (l *LRUCache) addHead(n *Node) {
	if l.len == 0 {
		l.head = n
		l.tail = n
	} else {
		_node := l.head // 中转
		l.head = n
		_node.Prev = n
		l.head.Next = _node
	}
	l.len++
}

func (l *LRUCache) remove(n *Node) {
	if l.head == n { // 判断节点是否为头部节点
		l.head = l.head.Next
		l.head.Prev = nil
	} else if l.tail == n { // 判断节点是否为尾部节点
		l.tail = l.tail.Prev
		l.tail.Next = nil
	} else { // 中间节点
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
	}
	l.len--
}

func (l *LRUCache) Get(key int) int {
	l.mx.Lock()
	defer l.mx.Unlock()

	if node, ok := l.cache[key]; ok {
		// 删除节点
		l.remove(node)
		l.addHead(node)
		return node.Value
	}

	return -1
}

func (l *LRUCache) Put(key, value int) {
	l.mx.Lock()
	defer l.mx.Unlock()

	if node, ok := l.cache[key]; ok {
		// 节点放最头部
		l.remove(node)
		l.addHead(node)
	} else {
		if l.len == l.cap {
			l.remove(l.tail)
		}
		_node := NewNode(key, value)
		l.addHead(_node)
		l.cache[key] = _node
	}
}

var _ LRUCacher = (*LRUCache)(nil)
