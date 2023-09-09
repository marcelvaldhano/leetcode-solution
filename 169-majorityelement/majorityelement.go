package main

import "fmt"

func majorityElement(nums []int) int {
	dict := map[int]int{}
	for i := 0; i < len(nums); i++ {
		dict[nums[i]] += 1
	}
	maks := 0
	element := -1
	for key, v := range dict {
		if v > maks {
			maks = v
			element = key
		}
	}
	return element
}

/* using booyer moore */
func majorityElement1(nums []int) int {
	candidate := nums[0]
	count := 1

	for _, v := range nums[1:] {
		if v != candidate {
			count--
		} else if v == candidate {
			count++
		}
		if count == 0 {
			candidate = v
			count = 1
		}
	}
	return candidate
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement1(nums))

}
