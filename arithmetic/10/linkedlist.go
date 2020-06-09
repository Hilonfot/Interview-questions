package _0

import "fmt"

type ILinkedListNode interface {
	Next() *LinkedListNode
	Value() int
}

type LinkedListNode struct {
	next *LinkedListNode
	Data int
}

func NewLinkedListNode(data int) *LinkedListNode {
	return &LinkedListNode{Data: data, next: nil}
}

func (l *LinkedListNode) Next() *LinkedListNode {
	return l.next
}

func (l *LinkedListNode) Value() int {
	return l.Data
}

type LinkedList struct {
	Head   *LinkedListNode
	Tail   *LinkedListNode
	Length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Head: NewLinkedListNode(0), Length: 0}
}

// 从头部插入
func (l *LinkedList) InsertHeadLinkedList(linked *LinkedListNode) {
	if l.Head == nil {
		l.Head.next = linked
		linked.next = nil
		l.Length++
	} else {
		head := l.Head
		linked.next = head.next
		head.next = linked
		l.Length++
	}
}

func (l *LinkedList) String() string {
	next := l.Head.next
	var str string
	for next != nil {
		if next.Value() != 0 {
			str += fmt.Sprintf(" --> %d", next.Value())
		}
		next = next.Next()
	}
	return str
}

func (l *LinkedList) Reversed()  {
	next := l.Head.next
	pre := &LinkedListNode{}
	for next != nil {
		temp := next.next
		next.next = pre
		pre = next
		next = temp
	}
	l.Head.next = pre
}
