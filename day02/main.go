package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInputData() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func getIntList(list []string) []int {
	intList := make([]int, len(list))
	for i, num := range list {
		intList[i], _ = strconv.Atoi(num)
	}
	return intList
}

func checkRule1(list []int) bool {
	increasing := false
	decreasing := false
	for i := 0; i < len(list)-1; i++ {
		// If these are the first values, check if the list is increasing or decreasing
		if !increasing && !decreasing {
			// If the first two values are equal, the list is invalid
			if list[i] == list[i+1] {
				return false
			}
			// If the first value is less than the second value, the list is increasing
			if list[i] < list[i+1] {
				increasing = true
			}
			// If the first value is greater than the second value, the list is decreasing
			if list[i] > list[i+1] {
				decreasing = true
			}
		}
		// If the previous values were increasing make sure the next value continues to be increasing
		if increasing {
			if list[i] >= list[i+1] {
				return false
			}
		}
		// If the previous values were decreasing make sure the next value continues to be decreasing
		if decreasing {
			if list[i] <= list[i+1] {
				return false
			}
		}
	}
	// If the loop completes without returning, the list is valid
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkRule2(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		difference := abs(list[i+1] - list[i])
		// If the difference is greater than 3 or less than 1, the list is invalid so return the index of the error
		if difference > 3 || difference < 1 {
			return false
		}
	}
	// If the loop completes without returning, the list is valid so the error index is -1
	return true
}

func checkRules(list []int) bool {
	if !checkRule1(list) {
		return false
	}
	if !checkRule2(list) {
		return false
	}
	return true
}

func day02_1(inputData [][]string) int {
	safeReports := 0
	for _, row := range inputData {
		// Convert row to int list
		intList := getIntList(row)

		// Check both rules and if both pass increment safe reports counter
		if checkRules(intList) {
			safeReports++
		}
	}
	return safeReports
}

func removeIndex(list []int, index int) []int {
	newList := []int{}
	for i := 0; i < len(list); i++ {
		if i != index {
			newList = append(newList, list[i])
		}
	}
	return newList
}

func day02_2(inputData [][]string) int {
	safeReports := 0
	for _, row := range inputData {
		intList := getIntList(row)
		if checkRules(intList) {
			safeReports++
		} else { // If the row is invalid, check each possible row with one value removed
			for i := 0; i < len(row); i++ {
				newList := removeIndex(intList, i)
				if checkRules(newList) {
					safeReports++
					break
				}
			}
		}
	}
	return safeReports
}

func main() {
	inputData := getInputData()
	fmt.Println("Day 2, Part 1: Total safe reports = ", day02_1(inputData))
	fmt.Println("Day 2, Part 2: Total safe reports after dampening = ", day02_2(inputData))
}
