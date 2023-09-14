package main

import "fmt"

func missingNumber(nums []int) int {
	dict := make(map[int]bool, len(nums))

	for i := 0; i <= len(nums); i++ {
		dict[i] = false
	}

	for _, v := range nums {
		dict[v] = true
	}

	for i, v := range dict {
		if !v {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 6, 5, 7, 0, 1}))
}
