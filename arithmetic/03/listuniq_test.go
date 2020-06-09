package _3

import "testing"

func TestListUniq(t *testing.T) {
	list := []int{1, 2, 3, 3, 1, 1, 2, 5, 6, 7}
	ListUniq(list)
	ListUniq1(list)
	t.Log(list)
}
