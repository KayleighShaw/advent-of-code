package main

import (
	"bufio"
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

func checkBoard(boards []int, bingoNumber int) {

}

func main() {
	array := getDataArray()
	// bingoNumbers := strings.Split(array[0], ",")
	// fmt.Println(bingoNumbers)

	getAllBoards(array)
	// iterate over bingoNumbers and check each

}
