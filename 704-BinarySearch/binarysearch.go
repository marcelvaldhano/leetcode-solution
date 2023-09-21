package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	mid := (start + end) / 2

	for start <= end {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
			mid = (start + end) / 2
		} else if nums[mid] > target {
			end = mid - 1
			mid = (start + end) / 2
		}
	}
	return -1

}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))

}
