package _0

import (
	"testing"
)

func TestLinked_Push(t *testing.T) {
	list := NewLinkedList()
	list.InsertHeadLinkedList(NewLinkedListNode(1))
	list.InsertHeadLinkedList(NewLinkedListNode(2))
	list.InsertHeadLinkedList(NewLinkedListNode(3))
	list.InsertHeadLinkedList(NewLinkedListNode(4))

	t.Log(list.String())
	list.Reversed()
	t.Log(list.String())
}
