package main

import (
	"bufio"
	"fmt"
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

func getDataArray() []Movement {
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

	for i := range text {
		results := re.FindStringSubmatch(text[i])
		var currentMove Movement
		currentMove.StartX, _ = strconv.Atoi(results[1])
		currentMove.StartY, _ = strconv.Atoi(results[2])
		currentMove.EndX, _ = strconv.Atoi(results[3])
		currentMove.EndY, _ = strconv.Atoi(results[4])
		moveList = append(moveList, currentMove)
	}

	fmt.Print(moveList)
	file.Close()
	return moveList
}

func main() {
	getDataArray()
}
