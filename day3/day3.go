package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func main() {
	binaryArray := getDataArray()
	numLines := len(binaryArray)
	bitWidth := len(binaryArray[0])
	slice := make([]int, bitWidth)

	gammarateBinString := ""
	epsilonBinString := ""

	for _, s := range binaryArray {
		for j, d := range s {
			if d == '1' {
				slice[j] += 1
			}
		}
	}
	fmt.Println(slice)

	for _, n := range slice {
		if numLines-n < n {
			gammarateBinString += "1"
			epsilonBinString += "0"

		} else {
			gammarateBinString += "0"
			epsilonBinString += "1"
		}
	}

	gammarate, _ := strconv.ParseInt(gammarateBinString, 2, 64) // 22
	epsilonrate, _ := strconv.ParseInt(epsilonBinString, 2, 64)
	powerConsumption := gammarate * epsilonrate

	fmt.Println(powerConsumption)
}
