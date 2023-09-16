package main

import (
	"fmt"
	"math"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	var count int
	for i := 0; i < len(nums); i += 2 {
		count += int(math.Min(float64(nums[i]), float64(nums[i+1])))
	}
	return count
}

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
	fmt.Println(arrayPairSum([]int{6, 2, 6, 5, 1, 2}))
}
