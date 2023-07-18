package main

import (
	"fmt"
	"math/rand"
)

type Board struct {
	cells [][]int
}

func (b *Board) Init() {
	b.cells = Make2D[int](4, 4)
	startCount := 2
	i := 0
	for i < startCount {
		row := int64(rand.Float64() * 4.0)
		col := int64(rand.Float64() * 4.0)
		fmt.Printf("%d:%d\n", row, col)
		b.cells[row][col] = 2
		i++
	}

}

func (b *Board) ToString() string {
	s := ""
	for i := range b.cells {
		s += fmt.Sprintf("%v\n", b.cells[i])
	}
	return s
}

func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}
