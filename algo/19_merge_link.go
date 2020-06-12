package main

import "fmt"

func main() {

	s1 := NewSingleLink()
	s1.Append(1)
	s1.Append(7)
	s1.Append(8)
	fmt.Println(s1)
	s2 := NewSingleLink()
	s2.Append(2)
	s2.Append(4)
	s2.Append(6)
	s2.Append(8)
	s2.Append(10)
	fmt.Println(s2)

	s1.head = MergeLinked2(s1.head, s2.head)
	s1.len = s1.len + s2.len
	fmt.Println(s1)
	fmt.Println(i)
}

var i = 1

// 递归实现链表合并,并排序
func MergeLinked2(n1, n2 *SingleLinkNode) *SingleLinkNode {
	i++
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Value.(int) < n2.Value.(int) {
		n1.Next = MergeLinked2(n1.Next, n2)
		return n1
	} else {
		n2.Next = MergeLinked2(n1, n2.Next)
		return n2
	}
}
