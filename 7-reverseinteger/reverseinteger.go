package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	isNumberMinus := false
	if x < 0 {
		x = x * (-1)
		isNumberMinus = true
	}
	result := 0
	for x > 0 {
		remainder := x % 10
		result = result*10 + remainder
		x /= 10
	}
	if isNumberMinus {
		result = result * (-1)
	}
	if result > int(math.Pow(2, 31)-1) {
		return 0
	} else if result < (int(math.Pow(2, 31)-1) * -1) {
		return 0
	}
	return result
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}
