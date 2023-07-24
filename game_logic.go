package main

const (
	LOSS string = "LOSS"
	WIN  string = "WIN"
)

type GameLogic struct {
	board *Board
}

func NewGameLogic(board *Board) *GameLogic {
	gameLogic := &GameLogic{board}
	return gameLogic
}

func (gameLogic *GameLogic) GameState() string {

	if gameLogic.HasMetWinCon() {
		return WIN
	}

	return LOSS
}

func (gameLogic *GameLogic) HasMetWinCon() bool {
	if gameLogic.isFinished() {
		for _, row := range gameLogic.board.cells {
			for i, _ := range row {
				if row[i] != 2048 {
					return false
				}
			}
		}
		return true
	}
	return false
}

func (gameLogic *GameLogic) GetResult() string {
	return gameLogic.GameState()
}

func (gameLogic *GameLogic) isFinished() bool {
	return gameLogic.board.isFull()
}

var score = 0

func (gameLogic *GameLogic) GetScore() int {
	return score
}
