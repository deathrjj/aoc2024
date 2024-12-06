package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		var chars []string
		for _, c := range line {
			chars = append(chars, string(c))
		}
		data = append(data, chars)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func findStartingPosition(data [][]string) (int, int) {
	for i := range data {
		for j := range data[i] {
			if data[i][j] == "^" {
				return i, j
			}
		}
	}
	return -1, -1
}

func checkIfWall(data [][]string, row, col int) bool {
	if data[row][col] == "#" {
		return true
	}
	return false
}

func rotate(direction string) string {
	switch direction {
	case "N":
		return "E"
	case "E":
		return "S"
	case "S":
		return "W"
	case "W":
		return "N"
	}
	return ""
}

func moveInDirection(row, col int, direction string) (int, int) {
	switch direction {
	case "N":
		return row - 1, col
	case "E":
		return row, col + 1
	case "S":
		return row + 1, col
	case "W":
		return row, col - 1
	}
	return -1, -1
}

func getNextPosition(data [][]string, row, col int, direction string) (int, int, string) {
	// Find the coordinates after moving in the current direction
	newRow, newCol := moveInDirection(row, col, direction)
	// Check if the new position is out of bounds
	if newRow < 0 || newRow >= len(data) || newCol < 0 || newCol >= len(data[newRow]) {
		return -1, -1, ""
	}
	// If the new position is a wall, stay in the same position and rotate
	if data[newRow][newCol] == "#" {
		newRow = row
		newCol = col
		direction = rotate(direction)
	}
	return newRow, newCol, direction
}

func printData(data [][]string, row, col int, direction string) {
	fmt.Print("\033[H\033[2J")
	switch direction {
	case "N":
		data[row][col] = "^"
	case "E":
		data[row][col] = ">"
	case "S":
		data[row][col] = "v"
	case "W":
		data[row][col] = "<"
	}
	for i := range data {
		fmt.Println()
		for j := range data[i] {
			fmt.Print(data[i][j])
		}
	}
}

func moveUntilOutOfBounds(data [][]string, row, col int, direction string) [][]string {
	maxSteps := len(data) * len(data[0]) * 4 // Maximum possible unique positions
	return moveWithLimit(data, row, col, direction, maxSteps)
}

func moveWithLimit(data [][]string, row, col int, direction string, stepsLeft int) [][]string {
	if stepsLeft <= 0 {
		data[row][col] = "Z"
		return data // Emergency exit if we've taken too many steps
	}

	data[row][col] = direction

	newRow, newCol, newDirection := getNextPosition(data, row, col, direction)
	if newRow == -1 || newCol == -1 {
		return data
	}

	// If the new position is the same as the new direction, we have been in this exact position before and have therefore looped
	if data[newRow][newCol] == newDirection {
		data[newRow][newCol] = "Z"
		return data
	}

	return moveWithLimit(data, newRow, newCol, newDirection, stepsLeft-1)
}

func countDistinctPositions(data [][]string) int {
	count := 0
	for i := range data {
		for j := range data[i] {
			if data[i][j] == "N" || data[i][j] == "E" || data[i][j] == "S" || data[i][j] == "W" {
				count++
			}
		}
	}
	return count
}

func day06_1(data [][]string) int {
	row, col := findStartingPosition(data)
	updatedData := moveUntilOutOfBounds(data, row, col, "N")
	count := countDistinctPositions(updatedData)
	return count
}

func checkForTimeLoop(data [][]string) bool {
	for i := range data {
		for j := range data[i] {
			if data[i][j] == "Z" {
				return true
			}
		}
	}
	return false
}

func copyData(data [][]string) [][]string {
	newData := make([][]string, len(data))
	for i := range data {
		newData[i] = make([]string, len(data[i]))
		copy(newData[i], data[i])
	}
	return newData
}

func day06_2(data [][]string) int {
	timeLoops := 0
	startRow, startCol := findStartingPosition(data)
	for i := range data {
		for j := range data[i] {
			dataWithBlock := copyData(data) // Create a deep copy
			dataWithBlock[i][j] = "#"
			updatedData := moveUntilOutOfBounds(dataWithBlock, startRow, startCol, "N")
			if checkForTimeLoop(updatedData) {
				timeLoops++
			}
		}
	}
	return timeLoops
}

func main() {
	data := getInputData()
	fmt.Println("Day 06 - Part 1: Distinct Positions = ", day06_1(data))
	data = getInputData()
	fmt.Println("Day 06 - Part 2: Time Loops = ", day06_2(data))
}
