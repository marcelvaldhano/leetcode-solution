package main

import "fmt"

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	if target > nums[end] {
		return end + 1
	}

	for start <= end {
		middle := start + (end-start)/2
		if target == nums[middle] {
			return middle
		}
		if target < nums[middle] {
			end = middle - 1
			continue
		}
		start = middle + 1
	}
	return start
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
