package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func removeEmpty(in []string, remove string) []string {
	out := make([]string, 0)

	for _, s := range in {
		if s != remove {
			out = append(out, s)
		}
	}

	return out
}

func makeInputOutputArrays(input []string, output []string) ([]string, []string) {
	inputArray := removeEmpty(input, "")
	outArray := removeEmpty(output, "")

	return inputArray, outArray
}

func getDataArray() [][]string {
	file, err := os.Open("input.txt")
	// file, err := os.Open("day8test.txt")
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	mainArray := make([][]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		newLine := strings.Split(line, "|")
		input := strings.Split(newLine[0], " ")  // one line array
		output := strings.Split(newLine[1], " ") // one line array
		inputArray, outputArray := makeInputOutputArrays(input, output)
		mainArray = append(mainArray, inputArray)
		mainArray = append(mainArray, outputArray)
	}

	// fmt.Print(mainArray[0])

	file.Close()
	return mainArray
}

func main() {
	array := getDataArray()
	fmt.Println(array)
	// fmt.Println(array[1])
	// fmt.Println(array[2])
	// fmt.Println(array[3])

	// newArray := [4]string{"cg", "cg", "fdcagb", "cbg"}

	count := 0

	for i := 1; i < len(array); i += 2 {
		for j := 0; j < len(array[i]); j++ {
			switch {
			case len(array[i][j]) == 2:
				count++
			case len(array[i][j]) == 4:
				count++
			case len(array[i][j]) == 3:
				count++
			case len(array[i][j]) == 7:
				count++
			default:
				continue
			}
		}
	}

	fmt.Println(count)
}
