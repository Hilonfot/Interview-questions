package main

// 冒泡交换,交换数据，不是交换节点
func SortByBubble(link *SingleLink) *SingleLink {
	len := link.len
	if len == 0 || len == 1 {
	} else if len == 2 {
		link.head, link.tail = link.tail, link.head
	} else {
		p1 := link.head
		for p1 != nil {
			p2 := p1.Next
			for p2 != nil {
				if p1.Value.(int) < p2.Value.(int) {
					p1.Value, p2.Value = p2.Value, p1.Value
				}
				p2 = p2.Next
			}
			p1 = p1.Next
		}
	}
	return link
}
