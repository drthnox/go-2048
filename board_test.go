package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	board := NewBoard()
	board.Set(1, 1, 2)
	board.Set(2, 2, 2)
	fmt.Println(board.ToString())

	board = board.Reverse()
	fmt.Println(board.ToString())

	assert.Equal(t, board.cells[1][1], 0)
	assert.Equal(t, board.cells[2][2], 0)
	assert.Equal(t, board.cells[1][2], 2)
	assert.Equal(t, board.cells[2][1], 2)
}
