package main

import "fmt"

func maxProductDifference(nums []int) int {
	max := make([]int, 2)
	min := make([]int, 2)
	min[0], min[1] = nums[0], nums[0]

	for i := range nums {
		if nums[i] >= max[1] {
			max[1] = nums[i]
			if max[1] > max[0] {
				max[0], max[1] = max[1], max[0]
			}
		}
		if nums[i] <= min[1] {
			min[1] = nums[i]
			if min[1] < min[0] {
				min[0], min[1] = min[1], min[0]
			}
		}
	}
	fmt.Println("maks: ", max[0], max[1])
	fmt.Println("maks: ", min[0], min[1])
	return (max[0] * max[1]) - (min[1] * min[0])
}

// catsandog

func main() {
	fmt.Println(maxProductDifference([]int{1, 6, 7, 5, 2, 4, 10, 6, 4}))
}
