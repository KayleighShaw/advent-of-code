package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getDataArray() []string {
	// file, err := os.Open("input.txt")
	file, err := os.Open("day8test.txt")
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	trimmedArray := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		newLine := strings.Split(line, "|")

		for _, l := range newLine {
			l = strings.Trim(l, " ")
			trimmedArray = append(trimmedArray, l)
		}
	}

	file.Close()
	return trimmedArray
}

func main() {
	array := getDataArray()
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[2])
	fmt.Println(array[3])
}
