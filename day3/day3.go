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

func cullArray(bitPosition int, desiredValue byte, array []string) []string {
	newArray := make([]string, 0)

	for _, s := range array {
		if s[bitPosition] == desiredValue {
			newArray = append(newArray, s)
		}
	}

	return newArray
}

func getMostCommonValue(array []string) (string, string) {
	bitWidth := len(array[0])
	slice := make([]int, bitWidth)

	var mostCommon string
	var leastCommon string

	// iterating through array, counts number of 1s in line
	for _, s := range array {
		for j, d := range s {
			if d == '1' {
				slice[j] += 1
			}
		}
	}

	// for each line, returns a string for most and least common
	for _, n := range slice {
		if len(array)-n == n {
			mostCommon += "1"
			leastCommon += "0"
		}
		if len(array)-n < n {
			mostCommon += "1"
			leastCommon += "0"
		} else {
			mostCommon += "0"
			leastCommon += "1"
		}
	}
	return mostCommon, leastCommon
}

func main() {
	binaryArray := getDataArray()
	bitWidth := len(binaryArray[0])

	mostCommon, leastCommon := getMostCommonValue(binaryArray)

	filteredArray := binaryArray

	for i := 0; i < bitWidth; i++ {
		filteredArray = cullArray(i, mostCommon[i], filteredArray)
		if len(filteredArray) == 1 {
			break
		}
		mostCommon, leastCommon = getMostCommonValue(filteredArray)
	}

	oxyRate, _ := strconv.ParseInt(filteredArray[0], 2, 64)
	fmt.Print("Oxygen Generator Rate: ")
	fmt.Println(oxyRate)

	filteredArray = binaryArray

	for i := 0; i < bitWidth; i++ {
		filteredArray = cullArray(i, leastCommon[i], filteredArray)
		if len(filteredArray) == 1 {
			break
		}
		mostCommon, leastCommon = getMostCommonValue(filteredArray)
	}

	scrubRate, _ := strconv.ParseInt(filteredArray[0], 2, 64)
	fmt.Print("CO2 Scrubber Rate: ")
	fmt.Println(scrubRate)

	fmt.Println(oxyRate * scrubRate)

}

// oxygen generator rating 01001
