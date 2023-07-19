package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	board := NewBoard()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		input = strings.ToLower(input)
		if input == "q" {
			break
		}
		switch input {
		case "?":
			showHelp()
		case "h":
			showHelp()
		case "w":
			moveUp()
		case "s":
			moveDown()
		case "a":
			moveLeft()
		case "d":
			moveRight()
		}
		board.Add2()
		fmt.Printf("%s\n", board.ToString())
	}
}

func NewBoard() *Board {
	board := new(Board)
	board.Init()
	return board
}

func moveRight() {

}

func moveLeft() {

}

func moveDown() {

}

func moveUp() {

}

func showHelp() {
	fmt.Println("W - up")
	fmt.Println("S - down")
	fmt.Println("A - left")
	fmt.Println("D - right")
	fmt.Println("h/? - help")
}
