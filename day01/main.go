package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// ReadCsvFile reads a CSV file and returns a 2D slice of strings
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath) // #nosec G304
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

// ParseInt converts a string to an integer
func parseInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Error converting string to int:", err)
	}
	return num
}

// Abs returns the absolute value of an integer
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// getOccurences returns the number of times a number appears in a sorted list
func getOccurences(searchNum int, list []int) int {
	occurences := 0
	for _, num := range list {
		if num == searchNum {
			occurences++
		}
		if num > searchNum {
			break
		}
	}
	return occurences
}

func getSortedLists() ([]int, []int) {
	inputData := readCsvFile("input.csv")
	var leftList, rightList []int
	for _, row := range inputData {
		leftList = append(leftList, parseInt(row[0]))
		rightList = append(rightList, parseInt(row[1]))
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	return leftList, rightList
}

func getTotalDistance(leftList, rightList []int) int {
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
	}
	return totalDistance
}

func getSimilarityScore(leftList, rightList []int) int {
	totalSimilarity := 0
	for _, leftNum := range leftList {
		occurences := getOccurences(leftNum, rightList)
		totalSimilarity += (leftNum * occurences)
	}

	return totalSimilarity
}

func main() {
	leftList, rightList := getSortedLists()
	fmt.Println("Day 1, Part 1: Total Distance =", getTotalDistance(leftList, rightList))
	fmt.Println("Day 1, Part 2: Similarity Score =", getSimilarityScore(leftList, rightList))
}
