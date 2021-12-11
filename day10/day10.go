package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getDataArray() [][]string {
	// file, err := os.Open("day10test.txt")
	// file, err := os.Open("shortertest.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	mainArray := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		newLine := strings.SplitAfter(line, "") //this will split per char
		mainArray = append(mainArray, newLine)
	}

	file.Close()
	return mainArray
}

// [ ( { ( < ( ( [ [ [ { [ { < ( <
// ) ) ] > ] ) > >
// [ ( { ( [ [ { { // No more bois
// { ( [ ( < [  } > { [ ] { [ ( < ( ) > // bad curly boi

func ignoreCorruptedLine(slice []string) bool {

	return false
}

func isOpening(bracket string, array []string) []string { // will append string to array if the bracket is an opening bracket
	array = append(array, bracket)
	return array
}

func isClosing(bracket string, slice []string) ([]string, string) {
	isCorrectClosing := false

	// check last element of the array
	lastOpeningBracket := slice[len(slice)-1]

	// check bracket matches opening bracket
	switch {
	case lastOpeningBracket == "{" && bracket == "}":
		isCorrectClosing = true
	case lastOpeningBracket == "(" && bracket == ")":
		isCorrectClosing = true
	case lastOpeningBracket == "<" && bracket == ">":
		isCorrectClosing = true
	case lastOpeningBracket == "[" && bracket == "]":
		isCorrectClosing = true
	default:
		isCorrectClosing = false
	}

	// this will remove the last element in the slice if the closing bracket matches
	if len(slice) > 0 && isCorrectClosing {
		slice = slice[:len(slice)-1]
	}

	if isCorrectClosing == false {
		return slice, bracket
	}

	return slice, "" // will return normal slice and empty string if no removals made
}

func main() {
	syntaxArray := getDataArray()
	// uncorruptedLines := make([][]string, 0)

	openingBracketArray := make([]string, 0)
	var incorrectClosing string
	var score int

	// fmt.Println(syntaxArray)

	for i := 0; i < len(syntaxArray); i++ { // iterates over each line array

		for _, s := range syntaxArray[i] { // iterates other each string in the line array
			if s == "{" || s == "(" || s == "<" || s == "[" {
				openingBracketArray = isOpening(s, openingBracketArray)
			}
			if s == "}" || s == ")" || s == ">" || s == "]" {
				openingBracketArray, incorrectClosing = isClosing(s, openingBracketArray)
			}

			fmt.Printf("Incorrect closing bracket: %v\n", incorrectClosing)

			if incorrectClosing != "" {
				break
			}
		}
		if incorrectClosing != "" {
			switch {
			case incorrectClosing == ")":
				score += 3
			case incorrectClosing == "]":
				score += 57
			case incorrectClosing == "}":
				score += 1197
			case incorrectClosing == ">":
				score += 25137
			default:
				fmt.Println("No change in score")
			}
		}
		incorrectClosing = ""
	}
	fmt.Println(incorrectClosing)
	fmt.Printf("The score for closing brackets: %d", score)
}
