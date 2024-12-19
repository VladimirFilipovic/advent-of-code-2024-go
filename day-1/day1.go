package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func compareChiefHistorianLists(listLocation string) {
	wd, _ := os.Getwd()

	fileLocation := wd + listLocation
	file, err := os.Open(fileLocation)

	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	locationIdPairs := strings.Split(string(content), "\r\n")

	fmt.Println("Location id pairs: ", strings.Join(locationIdPairs, ", "))

	distance := 0
	list1, list2 := make([]int, len(locationIdPairs)), make([]int, len(locationIdPairs))

	// Populate the lists
	for i, locationIdPair := range locationIdPairs {
		locationIds := strings.Fields(locationIdPair)

		value, _ := strconv.Atoi(locationIds[0])
		value2, _ := strconv.Atoi(locationIds[1])

		list1[i] = value
		list2[i] = value2
	}

	slices.Sort(list1)
	slices.Sort(list2)

	for i := 0; i < len(list1); i++ {
		distance += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println("Distance: ", distance)
}

func compareChiefHistorianListsPart2(listLocation string) {
	wd, _ := os.Getwd()

	fileLocation := wd + listLocation
	file, err := os.Open(fileLocation)

	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	locationIdPairs := strings.Split(string(content), "\r\n")

	fmt.Println("Location id pairs: ", strings.Join(locationIdPairs, ", "))

	similarity := 0
	list1, list2 := make([]int, len(locationIdPairs)), make([]int, len(locationIdPairs))

	// Populate the lists
	for i, locationIdPair := range locationIdPairs {
		locationIds := strings.Fields(locationIdPair)

		value, _ := strconv.Atoi(locationIds[0])
		value2, _ := strconv.Atoi(locationIds[1])

		list1[i] = value
		list2[i] = value2
	}

	slices.Sort(list1)
	slices.Sort(list2)

	for i, locationId := range list1 {
		locationIndex := slices.Index(list2, list1[i])

		if locationIndex == -1 {
			continue
		}

		countOfItemOccurrence := 0
		for j := locationIndex; j < len(list2); j++ {
			if list2[j] == locationId {
				countOfItemOccurrence++
			}
			if list2[j] != locationId {
				similarity += locationId * countOfItemOccurrence
				break
			}
		}

	}

	fmt.Println("Similarity: ", similarity)
}

func main() {
	compareChiefHistorianLists("/day-1/input.txt")
	compareChiefHistorianListsPart2("/day-1/input.txt")
}
