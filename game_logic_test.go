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

		assert.True(t, gameLogic.HasMetWinCon())
	})
}

//func TestWinCondition_WhenHasNotWon(t *testing.T) {
//
//}
