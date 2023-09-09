package main

import "fmt"

func removeElement(nums []int, val int) int {
	counter := 0
	for _, v := range nums {
		if v != val {
			nums[counter] = v
			counter++
		}
	}
	return counter

}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	fmt.Println(removeElement(nums, val))

}
