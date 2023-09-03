package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
			continue
		} else if i%5 == 0 {
			result = append(result, "Buzz")
			continue
		} else if i%3 == 0 {
			result = append(result, "Fizz")
			continue
		}
		result = append(result, strconv.Itoa(i))
	}
	return result

}

func main() {
	fmt.Println(fizzBuzz(3))
}
