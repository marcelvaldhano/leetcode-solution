package main

import "fmt"

func maxProduct(nums []int) int {
	res := make([]int, len(nums))
	for i := range nums {
		if nums[i] >= res[1] {
			res[1] = nums[i]
			if res[0] < res[1] {
				res[0], res[1] = res[1], res[0]
			}
		}
	}
	return (res[0] - 1) * (res[1] - 1)
}

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
	fmt.Println(maxProduct([]int{1, 5, 4, 5}))
	fmt.Println(maxProduct([]int{3, 7}))
}
