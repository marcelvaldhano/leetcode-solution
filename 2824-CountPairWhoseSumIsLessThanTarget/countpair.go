package main

import "fmt"

func countPairs(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				count++
			}
		}
	}
	return count
}

func main() {
	nums := []int{-1, 1, 2, 3, 1}
	target := 2
	fmt.Println(countPairs(nums, target))
	nums = []int{-6, 2, 5, -2, -7, -1, 3}
	target = -2
	fmt.Println(countPairs(nums, target))
}
