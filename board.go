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
	for _, row := range b.cells {
		for i, _ := range row {
			row[i] = 0
		}
	}
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

func (b *Board) Reverse() *Board {
	board := NewBoard()
	for i, row := range b.cells {
		r := b.ReverseRow(row)
		board.cells[i] = r
	}
	return board
}

func (b *Board) ReverseRow(row []int) []int {
	newNumbers := make([]int, 0, len(row))
	for i := len(row) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, row[i])
	}
	fmt.Printf("> %v\n", newNumbers)
	return newNumbers
}

func (b *Board) Set(row int, col int, val int) {
	b.cells[row][col] = val
}
