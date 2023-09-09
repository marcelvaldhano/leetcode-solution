package main

import "fmt"

func removeDuplicates(nums []int) int {
	index := 0
	dict := map[int]int{}

	for _, v := range nums {
		if dict[v] < 2 {
			nums[index] = v
			index++
			dict[v] += 1
		}
	}
	return index

}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))

}
