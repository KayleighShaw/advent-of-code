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
	file, err := os.Open("day7test.txt")
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

func sumList(list []int) int {
	var sum int

	for _, c := range list {
		sum += c
	}

	return sum
}

func calculateFuelConsumption(crabArray []int, position int) int {
	fuelConsumption := make([]int, 0)
	for _, c := range crabArray {
		if c > position {
			fuelConsumption = append(fuelConsumption, c-position)
		}
		if c < position {
			fuelConsumption = append(fuelConsumption, position-c)
		}
		if c == position {
			fuelConsumption = append(fuelConsumption, 0)
		}
	}

	return sumList(fuelConsumption)
}

func getLargestPossibleMovement(array []int) int {
	var largestNumber int

	for _, c := range array {
		if c == 0 || c > largestNumber {
			largestNumber = c
		}
	}
	return largestNumber
}

func main() {
	crabArray := getDataArray()
	var smallestFuelUsage int
	largestMovement := getLargestPossibleMovement(crabArray)

	for position := 1; position < largestMovement; position++ {
		currentFuelUsage := calculateFuelConsumption(crabArray, position)
		if smallestFuelUsage == 0 || smallestFuelUsage > currentFuelUsage {
			smallestFuelUsage = currentFuelUsage
		}
	}
	fmt.Printf("Smallest fuel usage: %d\n", smallestFuelUsage)
}
