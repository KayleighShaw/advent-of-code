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
	// file, err := os.Open("input.txt")
	file, err := os.Open("day6test.txt")
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanLines)
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

func main() {
	fishArray := getDataArray()

	// fmt.Println(fishArray)

	// var babyFish []int //make this a slice
	babyFish := make([]int, 0)
	newFishiesList := make([]int, 0)

	for i := 0; i < 80; i++ {
		for _, f := range fishArray {
			// fmt.Println(f)
			switch {
			case f == 0:
				newFishiesList = append(newFishiesList, 6)
				babyFish = append(babyFish, 8)
			default:
				f--
				newFishiesList = append(newFishiesList, f)
			}
			// newFishiesList = append(newFishiesList, babyFish...)
			// fmt.Print("New Fish List: ")
			// fmt.Println(newFishiesList)
			// fmt.Print("Baby Fish: ")
			// fmt.Println(babyFish)
		}
		fishArray = newFishiesList
		fishArray = append(fishArray, babyFish...)
		// fmt.Printf("Day %d: %v\n", i+1, fishArray)
		babyFish = nil
		newFishiesList = nil
	}

	fmt.Printf("The number of fishies is: %d", len(fishArray))
}
