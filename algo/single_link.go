package main

import "fmt"

// 单链表只有一个节点，即下一个节点
// 单链表分头节点，尾节点
// 单链表特点就是没一个节点有一个指向下一个的节点，除尾节点为nil

// 定义一个通用类型
type Object interface{}

// 创建一个单链表的节点
type SingleLinkNode struct {
	Next  *SingleLinkNode
	Value Object
}

func (s *SingleLinkNode) String() string {
	var str string = "节点结果:"
	p := s
	for p != nil {
		str += fmt.Sprintf("%s", p.Value)
		p = p.Next
		if p != nil {
			str += "=>"
		}
	}
	return str
}

// 实例化一个节点
func NewSingleLinkNode(v Object) *SingleLinkNode {
	return &SingleLinkNode{
		Value: v,
	}
}

// 创建一个单链表结构
// 特点： 有头节点，尾节点，长的
type SingleLink struct {
	head *SingleLinkNode // 头节点
	tail *SingleLinkNode // 尾节点
	len  int             // 链表长度
}

func (s *SingleLink) String() string {
	// 判断链表是否为空
	var str string = "链表结果:"
	if s.len == 0 {
		str += "链表为空"
	}
	p := s.head
	for p != nil {
		str += fmt.Sprintf("%s", p.Value)
		p = p.Next
		if p != nil {
			str += "=>"
		}
	}
	return str
}

func (s *SingleLink) AddHead(v Object) {
	node := NewSingleLinkNode(v)
	if s.len == 0 {
		s.head = node
		s.tail = node
	} else {
		_node := s.head
		s.head = node
		s.head.Next = _node // 将新的节点下一个节点指针旧头部节点
	}
	s.len++
}

func (s *SingleLink) Append(v Object) {
	node := NewSingleLinkNode(v)
	if s.len == 0 {
		s.head = node
		s.tail = node
	} else {
		_tail := s.tail
		_tail.Next = node
		s.tail = node
	}
	s.len++
}

func (s *SingleLink) DelHead() {
	// 判断当前链表是否为空
	if s.len == 0 {
		return
	}
	node := s.head
	// 当只有一个节点
	if node.Next == nil {
		s.head = nil
		s.tail = nil
	} else {

		s.head = node.Next // 将下一个节点设置为头节点
	}
	s.len--
}

// 删除尾部节点，每一个节点没有记录上一个节点，删除了最后一个节点，无法重新设置新的尾部节点
// 使用循环找到最后一个节点
// 时间复杂度O(n), 主要员是没有上一个节点的指针，所以只有一个办法就是从头部找到尾部
func (s *SingleLink) DelTail() {
	// 判断当时是否尾空的单链表
	if s.len == 0 {
		return
	}

	if s.len == 1 { // 节点 = 1
		s.head = nil
		s.tail = nil
	} else if s.len == 2 { // 节点 = 2
		s.tail = s.head
		s.head.Next = nil
	} else { // 节点 > 2
		p := s.head
		for i := 0; i < s.len-1; i++ {
			p = p.Next // 继续向下找
		}
		s.tail = p
		p.Next = nil
	}
	s.len--
}

// 创建一个空链表
func NewSingleLink() *SingleLink {
	return &SingleLink{}
}

// 定义一个单链表接口
type SingleLinker interface {
	String() string   // 打印
	AddHead(v Object) // 向头部添加节点
	Append(v Object)  // 向尾部添加节点
	DelHead()         // 删除头部节点
	DelTail()         // 删除尾部节点
}

var _ SingleLinker = (*SingleLink)(nil)
