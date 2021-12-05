package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getDataArray() []string {
	file, err := os.Open("day4test.txt")
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

func removeEmpty(in []string) []string {
	out := make([]string, 0)

	for _, s := range in {
		if s != "" {
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
		newSlice := removeEmpty(slice)

		currentBoard = append(currentBoard, newSlice)
	}

	if currentBoard != nil {
		allBoards = append(allBoards, currentBoard)
	}

	return allBoards
}

func checkWin() {

}

func markNumber(boards [][][]string, bingoNumber string) [][][]string {
	numBoards := len(boards)
	xyMax := len(boards[0])
	var x, y, z int

	for z = 0; z < numBoards; z++ {
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

func main() {
	array := getDataArray()
	bingoNumbers := strings.Split(array[0], ",")
	allBoards := getAllBoards(array)

	for _, b := range bingoNumbers {
		allBoards = markNumber(allBoards, b)
	}

	fmt.Println(allBoards)

}
