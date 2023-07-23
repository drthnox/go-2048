package main

const (
	FULL string = "FULL"
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
	// check if board is full
	if gameLogic.board.isFull() {
		return FULL
	}

	return WIN
}

func HasMetWinCon(board *Board) bool {
	return false
}

func (gameLogic *GameLogic) moveRight() {
	// move all cells right
	// iterate over each rowNum
	//for _, row := range board.cells {
	//	for i, cell := range row {
	//		if i < len(row) {
	//			if row[i+1] == 0 {
	//				row[i+1] = row[i]
	//				row[i] = 0
	//			}
	//		}
	//	}
	//}
}

func (gameLogic *GameLogic) moveLeft() {
	// move all cells left
}

func (gameLogic *GameLogic) moveDown() {
	// move all cells down
}

func (gameLogic *GameLogic) moveUp() {
	// move all cells up
}

func (gameLogic *GameLogic) GetResult() string {
	return gameLogic.GameState()
}

func (gameLogic *GameLogic) isFinished() bool {
	return gameLogic.board.isFull()
}