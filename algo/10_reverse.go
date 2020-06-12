package main

// 单链反转，使用头插法
func ReverseLinked(head *SingleLinkNode) *SingleLinkNode {
	var curr *SingleLinkNode
	for head != nil {
		next := head.Next // 先将head另外存起来

		head.Next = curr // 类似 A = B + A
		curr = head

		head = next
	}
	return curr
}
