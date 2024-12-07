package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() ([]int, [][]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var results []int
	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Split the line by colon and get the second part
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		// Add the first part to the results array
		result, _ := strconv.Atoi(parts[0])
		results = append(results, result)

		// Split the numbers by space and convert to integers
		numStrs := strings.Fields(strings.TrimSpace(parts[1]))
		numbers := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			var num int
			_, err := fmt.Sscanf(numStr, "%d", &num)
			if err != nil {
				continue
			}
			numbers[i] = num
		}
		data = append(data, numbers)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return results, data
}

func generateAllPossibleOperations(numberOfOperations int) [][]byte {
	// Calculate total number of possibilities (2^n)
	total := 1 << numberOfOperations
	operations := make([][]byte, total)

	// Generate all possible combinations
	for i := 0; i < total; i++ {
		ops := make([]byte, numberOfOperations)
		for j := 0; j < numberOfOperations; j++ {
			// Check if jth bit is set in i
			if (i & (1 << j)) != 0 {
				ops[j] = 1
			} else {
				ops[j] = 0
			}
		}
		operations[i] = ops
	}

	return operations
}

func generateAllPossibleOperations2(numberOfOperations int) [][]string {
	// Calculate total number of possibilities (3^n)
	total := 1
	for i := 0; i < numberOfOperations; i++ {
		total *= 3
	}
	operations := make([][]string, total)

	// Generate all possible combinations
	for i := 0; i < total; i++ {
		ops := make([]string, numberOfOperations)
		temp := i
		for j := 0; j < numberOfOperations; j++ {
			// Get value for this position (0, 1, or 2)
			ops[j] = string('0' + byte(temp%3))
			temp /= 3
		}
		operations[i] = ops
	}

	return operations
}
func tryPossibleOperations(result int, data []int) bool {
	numberOfOperations := len(data) - 1
	operations := generateAllPossibleOperations(numberOfOperations)
	for _, operation := range operations {
		testResult := data[0]
		for i, op := range operation {
			if op == 0 {
				testResult += data[i+1]
			} else {
				testResult *= data[i+1]
			}
		}
		if testResult == result {
			return true
		}
	}
	return false
}

func tryPossibleOperations2(result int, data []int) bool {
	numberOfOperations := len(data) - 1
	operations := generateAllPossibleOperations2(numberOfOperations)
	for _, operation := range operations {
		testResult := data[0]
		for i, op := range operation {
			if op == "0" {
				testResult += data[i+1]
			} else if op == "1" {
				testResult *= data[i+1]
			} else {
				strNum1 := strconv.Itoa(testResult)
				strNum2 := strconv.Itoa(data[i+1])
				testResult, _ = strconv.Atoi(strNum1 + strNum2)
			}
		}
		if testResult == result {
			return true
		}
	}
	return false
}

func day07_1(results []int, data [][]int) int {
	total := 0
	for i, result := range results {
		if tryPossibleOperations(result, data[i]) {
			total += result
		}
	}
	return total
}

func day07_2(results []int, data [][]int) int {
	total := 0
	for i, result := range results {
		if tryPossibleOperations2(result, data[i]) {
			total += result
		}
	}
	return total
}

func main() {
	results, data := getInput()
	fmt.Println("Day 7 Part 1: Total of possibly valid results:", day07_1(results, data))
	fmt.Println("Day 7 Part 2: Total of possibly valid results:", day07_2(results, data))
}
