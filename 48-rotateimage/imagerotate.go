package main

import "fmt"

func rotate(matrix [][]int) [][]int {
	lenMatrix := len(matrix)
	result := [][]int{}
	for i := 0; i < len(matrix[0]); i++ {
		new_row := []int{}
		for j := lenMatrix - 1; j >= 0; j-- {
			new_row = append(new_row, matrix[j][i])
		}
		result = append(result, new_row)
	}
	copy(matrix, result)
	return matrix
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(rotate(matrix))
}
