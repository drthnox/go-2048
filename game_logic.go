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

var score = 0

func (gameLogic *GameLogic) GetScore() int {
	return score
}
