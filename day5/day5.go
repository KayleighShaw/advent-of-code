package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Movement struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
}

func getLargest(m Movement) int {
	largest := m.StartX

	if largest < m.StartY {
		largest = m.StartY
	}
	if largest < m.EndX {
		largest = m.EndX
	}

	if largest < m.EndY {
		largest = m.EndY
	}

	return largest
}

func getMoveList() ([]Movement, int) {
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

	re := regexp.MustCompile(`(\d*),(\d*)\s\-\>\s(\d*),(\d*)`)
	var moveList []Movement
	var largestNumber int

	for i := range text {
		results := re.FindStringSubmatch(text[i])
		var currentMove Movement
		currentMove.StartX, _ = strconv.Atoi(results[1])
		currentMove.StartY, _ = strconv.Atoi(results[2])
		currentMove.EndX, _ = strconv.Atoi(results[3])
		currentMove.EndY, _ = strconv.Atoi(results[4])
		currentLargest := getLargest(currentMove)
		if currentLargest > largestNumber {
			largestNumber = currentLargest
		}
		moveList = append(moveList, currentMove)
	}

	file.Close()
	return moveList, largestNumber
}

func makeGrid(maxCoord int) [][]int {
	grid := make([][]int, maxCoord)

	for i := 0; i < maxCoord; i++ {
		grid[i] = make([]int, maxCoord)
	}
	return grid
}

func markGrid(grid [][]int, move Movement) [][]int {

}

func main() {
	moveList, maxCoord := getMoveList()
	grid := makeGrid(maxCoord)

	for _, m := range moveList {
		grid = markGrid(grid, m)
	}

}
