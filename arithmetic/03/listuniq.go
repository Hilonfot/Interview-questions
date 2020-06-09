package _3

//* 列表去重 O(1)
func ListUniq(list []int) []int {
	listMap := make(map[int]int, len(list))
	newList := make([]int, 0, len(list))
	for _, val := range list {
		if _, ok := listMap[val]; !ok {
			listMap[val] = val
			newList = append(newList, val)
		}
	}
	return newList
}

// ** 采用冒泡思路
// 循环对比O(n)
func ListUniq1(list []int) []int {
	newList := make([]int, 0, len(list))
	for i := 0; i < len(list); i++ {
		flag := true
		for j := 0; j < len(newList); j++ {
			if list[i] == list[j] {
				flag = false
			}
		}
		if flag{
			newList = append(newList, list[i])
		}
	}
	return newList
}
