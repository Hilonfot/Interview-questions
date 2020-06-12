package main

import "fmt"

func main() {
	fmt.Println(repeat([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(duplicateByHash([]int{2, 3, 1, 0, 2, 5, 3}))
}

// 第二种 hash方法 时间复杂度O(n)
// 以空间复杂度换时间复杂度
func duplicateByHash(arr []int) []int {
	rep := make([]int, 0)
	m := make(map[int]int, len(arr))
	for k, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = k
		} else {
			rep = append(rep, v)
		}
	}
	return rep
}

// 第一种 暴力方法 时间复杂度O(n^2)
func repeat(arr []int) []int {
	rep := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				rep = append(rep, arr[i])
				break
			}
		}
	}
	return rep
}
