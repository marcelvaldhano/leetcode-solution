package main

import (
	"fmt"
	"strings"
)

var (
	belowTen = map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
	}
	belowTwenty = map[int]string{
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
	}
	belowHundred = map[int]string{
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
	}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	return convert(num)
}

func convert(num int) string {
	result := ""
	switch {
	case num < 10:
		result = belowTen[num]
	case num < 20:
		result = belowTwenty[num]
	case num < 100:
		result = belowHundred[num/10] + " " + belowTen[num%10]
	case num < 1000:
		result = convert(num/100) + " Hundred " + convert(num%100)
	case num < 1000000:
		result = convert(num/1000) + " Thousand " + convert(num%1000)
	case num < 1000000000:
		result = convert(num/1000000) + " Million " + convert(num%1000000)
	case num < 1000000000000:
		result = convert(num/1000000000) + " Billion " + convert(num%1000000000)
	}
	return strings.Trim(result, " ")

}

func main() {
	fmt.Println(numberToWords(1234567))
}
