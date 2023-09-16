package main

import (
	"fmt"
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	arr = sortArray(arr)
	target = sortArray(target)
	for i := 0; i < len(arr); i++ {
		if arr[i] != target[i] {
			return false
		}
	}
	return true
}
func canBeEqual1(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	count := [1001]int{}
	for i := range target {
		count[target[i]]++
		count[arr[i]]--
	}
	for i := range count {
		if count[i] != 0 {
			return false
		}
	}
	return true
}

func sortArray(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
	fmt.Println(canBeEqual1([]int{7}, []int{7}))
}
