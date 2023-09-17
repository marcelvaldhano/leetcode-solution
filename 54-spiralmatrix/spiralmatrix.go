package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	top := 0
	left := 0
	buttom := len(matrix)
	right := len(matrix[0])

	res := []int{}

	for top < buttom && left < right {
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for i := top; i < buttom; i++ {
			res = append(res, matrix[i][right-1])
		}
		right--

		if left >= right || top >= buttom {
			break
		}

		for i := right - 1; i >= left; i-- {
			res = append(res, matrix[buttom-1][i])
		}
		buttom--
		for i := buttom - 1; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}
