package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWinConditions(t *testing.T) {
	t.Run("Has Won", func(t *testing.T) {
		board := NewBoard()
		for _, row := range board.cells {
			for i, _ := range row {
				row[i] = 2048
			}
		}
		gameLogic := &GameLogic{board: board}

		assert.True(t, gameLogic.board.isFull())
		assert.True(t, gameLogic.HasMetWinCon())
	})

	t.Run("Has Lost", func(t *testing.T) {
		board := NewBoard()
		for _, row := range board.cells {
			for i, _ := range row {
				row[i] = 128
			}
		}
		gameLogic := &GameLogic{board: board}

		assert.True(t, gameLogic.board.isFull())
		assert.False(t, gameLogic.HasMetWinCon())
	})
}

//func TestWinCondition_WhenHasNotWon(t *testing.T) {
//
//}
