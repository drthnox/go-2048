package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	board := NewBoard()
	gameLogic := NewGameLogic(board)
	reader := bufio.NewReader(os.Stdin)
	for {
		input := getInput(reader)
		if input == "q" {
			break
		}
		switch input {
		case "?":
			showHelp()
		case "h":
			showHelp()
		case "w":
			gameLogic.moveUp()
		case "s":
			gameLogic.moveDown()
		case "a":
			gameLogic.moveLeft()
		case "d":
			gameLogic.moveRight()
		}
		board.Add2()
		fmt.Printf("%s\n", board.ToString())
		if gameLogic.isFinished() {
			fmt.Println("Board is full, cannot continue")
			fmt.Printf("Final Score: %d\n", gameLogic.GetScore())
			break
		} else {
			fmt.Println("Empty cells remain, can continue")
		}
	}
	fmt.Printf("Game Over! Result is %v\n", gameLogic.GetResult())
}

func getInput(reader *bufio.Reader) string {
	fmt.Print("> ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	input = strings.ToLower(input)
	return input
}

func NewBoard() *Board {
	board := new(Board)
	board.Init()
	return board
}

func showHelp() {
	fmt.Println("W - up")
	fmt.Println("S - down")
	fmt.Println("A - left")
	fmt.Println("D - right")
	fmt.Println("h/? - help")
}
