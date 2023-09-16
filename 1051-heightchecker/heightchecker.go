package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(heights)
	count := 0
	for i := range heights {
		if heights[i] != expected[i] {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
	fmt.Println(heightChecker([]int{5, 1, 2, 3, 4}))
}
