package main

import "fmt"

func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

func ReverseRow(row []int) []int {
	newNumbers := make([]int, 0, len(row))
	for i := len(row) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, row[i])
	}
	fmt.Printf("> %v\n", newNumbers)
	return newNumbers
}

func Reverse(cells [][]int) [][]int {
	for i, row := range cells {
		r := ReverseRow(row)
		cells[i] = r
	}
	return cells
}

func Transpose(cells [][]int) [][]int {
	newArr := make([][]int, len(cells))
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[0]); j++ {
			newArr[j] = append(newArr[j], cells[i][j])
		}
	}
	return newArr
}
