package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	result := make([]int, len(nums))

	start := 0
	end := len(nums) - 1
	for i := range nums {
		if nums[i]&1 == 0 {
			result[start] = nums[i]
			start++
			continue
		}
		result[end] = nums[i]
		end--
	}
	return result

}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{0}))
}
