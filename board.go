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
	b.Add2()
}

func (b *Board) Add2() {
	row := int64(rand.Float64() * 4.0)
	col := int64(rand.Float64() * 4.0)
	for {
		if b.cells[row][col] == 0 {
			b.cells[row][col] = 2
			break
		}
		row = int64(rand.Float64() * 4.0)
		col = int64(rand.Float64() * 4.0)
	}
}

func (b *Board) ToString() string {
	s := ""
	for i := range b.cells {
		s += fmt.Sprintf("%v\n", b.cells[i])
	}
	return s
}

func (b *Board) isEmpty() bool {
	for cell, _ := range b.cells {
		if cell != 0 {
			return false
		}
	}
	return true
}

func (b *Board) isFull() bool {
	for _, row := range b.cells {
		for _, val := range row {
			if val == 0 {
				return false
			}
		}
	}
	return true
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
