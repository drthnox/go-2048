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
	Zero(b.cells)
}

func Zero(b [][]int) {
	for _, row := range b {
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
	reversedCells := Reverse(b.cells)
	board := NewBoard()
	board.cells = reversedCells
	return board
}

func (b *Board) Set(row int, col int, val int) {
	b.cells[row][col] = val
}

// function to compress the grid
// after every step before and
// after merging cells.
func Compress(grid [][]int) ([][]int, bool) {
	changed := false
	new_mat := Make2D[int](4, 4)
	Zero(new_mat)

	//# here we will shift entries
	//# of each cell to it's extreme
	//# left row by row
	//# loop to traverse rows
	for i, row := range grid {
		pos := 0
		//# loop to traverse each column
		//# in respective row
		for j, _ := range row {
			if grid[i][j] != 0 {

				//# if cell is non empty then
				//# we will shift it's number to
				//# previous empty cell in that row
				//# denoted by pos variable
				new_mat[i][pos] = grid[i][j]

				if j != pos {
					changed = true
				}
				pos++
			}
		}
	}
	return new_mat, changed
}

func Merge(grid [][]int) ([][]int, bool) {
	changed := false

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {

			//# if current cell has same value as
			//# next cell in the row and they
			//# are non empty then
			if grid[i][j] == grid[i][j+1] && grid[i][j] != 0 {

				//# double current cell value and
				//# empty the next cell
				grid[i][j] = grid[i][j] * 2
				grid[i][j+1] = 0

				//# make bool variable True indicating
				//# the new grid after merging is
				//# different.
				changed = true
			}
		}
	}

	return grid, changed
}
