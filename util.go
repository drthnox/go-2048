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

func moveRight(grid [][]int) ([][]int, bool) {
	//# to move right we just reverse
	//# the matrix
	newGrid := Reverse(grid)

	//# then move left
	newGrid, changed := moveLeft(newGrid)

	//# then again reverse matrix will
	//# give us desired result
	newGrid = Reverse(newGrid)

	return newGrid, changed
}

// move all cells left
func moveLeft(grid [][]int) ([][]int, bool) {
	//# first compress the grid
	newGrid, changed1 := Compress(grid)

	//# then merge the cells.
	newGrid, changed2 := Merge(newGrid)

	changed := changed1 || changed2

	//# again compress after merging.
	newGrid, _ = Compress(newGrid)

	//# return new matrix and bool changed
	//# telling whether the grid is same
	//# or different
	return newGrid, changed
}

func moveDown(grid [][]int) ([][]int, bool) {
	// move all cells down
	//# to move down we take transpose
	newGrid := Transpose(grid)

	//# move right and then again
	newGrid, changed := moveRight(newGrid)

	//# take transpose will give desired
	//# results.
	newGrid = Transpose(newGrid)
	return newGrid, changed
}

func moveUp(grid [][]int) ([][]int, bool) {
	// move all cells up
	//# to move up we just take
	//# transpose of matrix
	newGrid := Transpose(grid)

	//# then move left (calling all
	//# included functions) then
	newGrid, changed := moveLeft(newGrid)

	//# again take transpose will give
	//# desired results
	newGrid = Transpose(newGrid)
	return newGrid, changed
}
