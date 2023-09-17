package main

import "fmt"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	dict1, dict2, dict3 := make(map[int]int, 101), make(map[int]int, 101), make(map[int]int, 101)
	for i := range nums1 {
		dict1[nums1[i]]++
	}
	for i := range nums2 {
		dict2[nums2[i]]++
	}
	for i := range nums3 {
		dict3[nums3[i]]++
	}
	result := []int{}
	for i := 0; i < 101; i++ {
		if dict1[i] != 0 && dict2[i] != 0 {
			result = append(result, i)
		} else if dict2[i] != 0 && dict3[i] != 0 {
			result = append(result, i)
		} else if dict3[i] != 0 && dict1[i] != 0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	fmt.Println(twoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3}))
	fmt.Println(twoOutOfThree([]int{3, 1}, []int{2, 3}, []int{1, 2}))
	fmt.Println(twoOutOfThree([]int{1, 2, 2}, []int{4, 3, 3}, []int{5}))
}
