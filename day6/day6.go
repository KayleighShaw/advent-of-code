package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getDataArray() []int {
	file, err := os.Open("input.txt")
	// file, err := os.Open("day6test.txt")
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(file)
	var text []string

	for scanner.Scan() {
		line := scanner.Text()
		text = strings.Split(line, ",")
	}

	var intArray []int

	for i := range text {
		var currentInt int
		currentInt, _ = strconv.Atoi(text[i])
		intArray = append(intArray, currentInt)
	}

	file.Close()
	return intArray
}

func generateAgeTracker(fish []int) [9]int {
	var ageTracker [9]int

	for _, f := range fish {
		ageTracker[f]++
	}
	return ageTracker
}

func sumAgeList(ageList [9]int) int {
	var sum int

	for _, c := range ageList {
		sum += c
	}

	return sum
}

func main() {
	fishArray := getDataArray()
	ageList := generateAgeTracker(fishArray)
	fmt.Println(ageList)

	// iterate through days
	for i := 0; i < 256; i++ {
		var updatedAgeList [9]int

		// iterate through ageList
		for j, f := range ageList {
			switch {
			case j == 0:
				updatedAgeList[6] = f
				updatedAgeList[8] = f
			default:
				updatedAgeList[j-1] += f
			}
		}
		ageList = updatedAgeList
	}

	count := sumAgeList(ageList)
	fmt.Printf("There are %d fish!", count)
}
