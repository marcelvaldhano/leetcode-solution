package main

import "fmt"

func containsDuplicate(nums []int) bool {
	dict := make(map[int]bool)

	for _, v := range nums {
		if _, ok := dict[v]; ok {
			return true
		}
		dict[v] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
