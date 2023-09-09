package main

import "fmt"

func rotate(nums []int, k int) []int {
	for i := 0; i < k; i++ {
		nums = swap(nums)
	}
	return nums
}

func swap(nums []int) []int {
	lengthNums := len(nums) - 1
	for i := 0; i < lengthNums; i++ {
		nums[i], nums[lengthNums] = nums[lengthNums], nums[i]
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	fmt.Println(rotate(nums, k))

	nums = []int{-1, -100, 3, 99}
	k = 2
	fmt.Println(rotate(nums, k))
}
