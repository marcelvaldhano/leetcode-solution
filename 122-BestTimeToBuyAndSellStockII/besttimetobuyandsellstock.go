package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > 0 {
			profit += prices[i] - min
			min = prices[i]
		}
	}
	return profit
}

func maxProfit2(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	price := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(price))
	price = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(price))
}
