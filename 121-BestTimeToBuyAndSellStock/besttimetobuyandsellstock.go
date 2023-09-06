package main

import "fmt"

func maxProfit(prices []int) int {
	maks := 0
	minPrices := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrices {
			minPrices = prices[i]
		} else if prices[i]-minPrices > maks {
			maks = prices[i] - minPrices
		}
	}
	return maks
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
