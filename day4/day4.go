package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getDataArray() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()
	return text
}

func removeEmpty(in []string, remove string) []string {
	out := make([]string, 0)

	for _, s := range in {
		if s != remove {
			out = append(out, s)
		}
	}

	return out
}

func getAllBoards(inputLines []string) [][][]string {
	var allBoards [][][]string
	var currentBoard [][]string

	for i, s := range inputLines {
		if i == 0 || i == 1 {
			continue
		}

		if s == "" {
			allBoards = append(allBoards, currentBoard)
			currentBoard = nil
			continue
		}

		slice := make([]string, 0)
		slice = strings.Split(s, " ")
		newSlice := removeEmpty(slice, "")

		currentBoard = append(currentBoard, newSlice)
	}

	if currentBoard != nil {
		allBoards = append(allBoards, currentBoard)
	}

	return allBoards
}

func rowWin(board [][]string) bool {
	lengthBoard := len(board[0])
	for row := 0; row < lengthBoard; row++ {
		if board[row][0] != "X" {
			continue
		}

		for col := 0; col < lengthBoard; col++ {
			if board[row][col] != "X" {
				break
			}

			if col == lengthBoard-1 {
				return true
			}
		}
	}
	return false
}

func columnWin(board [][]string) bool {
	lengthBoard := len(board[0])
	for col := 0; col < lengthBoard; col++ {
		if board[0][col] != "X" {
			continue
		}

		for row := 0; row < lengthBoard; row++ {
			if board[row][col] != "X" {
				break
			}

			if row == lengthBoard-1 {
				return true
			}
		}
	}
	return false
}

func checkWin(boards [][][]string) (bool, int) {
	for i, b := range boards {
		if b == nil {
			continue
		}
		colWin := columnWin(b)
		if colWin == true {
			return true, i
		}
		rowWin := rowWin(b)
		if rowWin == true {
			return true, i
		}
	}
	return false, -1
}

func markNumber(boards [][][]string, bingoNumber string) [][][]string {
	numBoards := len(boards)
	var x, y, z int

	for z = 0; z < numBoards; z++ {
		if boards[z] == nil {
			continue
		}
		xyMax := len(boards[z])
		for y = 0; y < xyMax; y++ {
			for x = 0; x < xyMax; x++ {
				if boards[z][y][x] == bingoNumber {
					boards[z][y][x] = "X"
				}
			}
		}
	}

	return boards
}

func sumMDArray(board [][]string) int {
	xyMax := len(board[0])
	var x, y, sum int

	for x = 0; x < xyMax; x++ {
		for y = 0; y < xyMax; y++ {
			if board[x][y] != "X" {
				number, _ := strconv.ParseInt(board[x][y], 10, 64)
				sum += int(number)
			}
		}
	}
	return sum
}

func main() {
	array := getDataArray()
	bingoNumbers := strings.Split(array[0], ",")
	allBoards := getAllBoards(array)
	var count int

	for _, b := range bingoNumbers {
		allBoards = markNumber(allBoards, b)

		for moreWinners := true; moreWinners; {
			isWinner, indexOfBoard := checkWin(allBoards)
			if isWinner == true {
				count++
				fmt.Print("Win count: ")
				fmt.Println(count)
				fmt.Print("The board index is: ")
				fmt.Println(indexOfBoard)
				finalNumber, _ := strconv.ParseInt(b, 10, 64)
				score := sumMDArray(allBoards[indexOfBoard]) * int(finalNumber)
				allBoards[indexOfBoard] = nil
				fmt.Print("The score is: ")
				fmt.Println(score)
			} else {
				moreWinners = false
			}
		}
	}
}
