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

func checkRight(data [][]string, row, col int) int {
	if col+3 >= len(data[row]) {
		return 0
	}
	if data[row][col] == "X" && data[row][col+1] == "M" && data[row][col+2] == "A" && data[row][col+3] == "S" {
		return 1
	}
	return 0
}

func checkRightDown(data [][]string, row, col int) int {
	if row+3 >= len(data) || col+3 >= len(data[row]) {
		return 0
	}
	if data[row][col] == "X" && data[row+1][col+1] == "M" && data[row+2][col+2] == "A" && data[row+3][col+3] == "S" {
		return 1
	}
	return 0
}

func checkDown(data [][]string, row, col int) int {
	if row+3 >= len(data) {
		return 0
	}
	if data[row][col] == "X" && data[row+1][col] == "M" && data[row+2][col] == "A" && data[row+3][col] == "S" {
		return 1
	}
	return 0
}

func checkLeftDown(data [][]string, row, col int) int {
	if row+3 >= len(data) || col-3 < 0 {
		return 0
	}
	if data[row][col] == "X" && data[row+1][col-1] == "M" && data[row+2][col-2] == "A" && data[row+3][col-3] == "S" {
		return 1
	}
	return 0
}

func checkLeft(data [][]string, row, col int) int {
	if col-3 < 0 {
		return 0
	}
	if data[row][col] == "X" && data[row][col-1] == "M" && data[row][col-2] == "A" && data[row][col-3] == "S" {
		return 1
	}
	return 0
}

func checkLeftUp(data [][]string, row, col int) int {
	if row-3 < 0 || col-3 < 0 {
		return 0
	}
	if data[row][col] == "X" && data[row-1][col-1] == "M" && data[row-2][col-2] == "A" && data[row-3][col-3] == "S" {
		return 1
	}
	return 0
}

func checkUp(data [][]string, row, col int) int {
	if row-3 < 0 {
		return 0
	}
	if data[row][col] == "X" && data[row-1][col] == "M" && data[row-2][col] == "A" && data[row-3][col] == "S" {
		return 1
	}
	return 0
}

func checkRightUp(data [][]string, row, col int) int {
	if row-3 < 0 || col+3 >= len(data[row]) {
		return 0
	}
	if data[row][col] == "X" && data[row-1][col+1] == "M" && data[row-2][col+2] == "A" && data[row-3][col+3] == "S" {
		return 1
	}
	return 0
}

func findXMAS(data [][]string) int {
	total := 0
	for i, row := range data {
		for j := range row {
			total += checkRight(data, i, j)
			total += checkRightDown(data, i, j)
			total += checkDown(data, i, j)
			total += checkLeftDown(data, i, j)
			total += checkLeft(data, i, j)
			total += checkLeftUp(data, i, j)
			total += checkUp(data, i, j)
			total += checkRightUp(data, i, j)
		}
	}
	return total
}

// M.M
// .A.
// S.S
func checkMTop(data [][]string, row, col int) int {
	if data[row][col] == "M" && data[row][col+2] == "M" &&
		data[row+1][col+1] == "A" &&
		data[row+2][col] == "S" && data[row+2][col+2] == "S" {
		return 1
	}
	return 0
}

// S.M
// .A.
// S.M
func checkMRight(data [][]string, row, col int) int {
	if data[row][col] == "S" && data[row][col+2] == "M" &&
		data[row+1][col+1] == "A" &&
		data[row+2][col] == "S" && data[row+2][col+2] == "M" {
		return 1
	}
	return 0
}

// S.S
// .A.
// M.M
func checkMBottom(data [][]string, row, col int) int {
	if data[row][col] == "S" && data[row][col+2] == "S" &&
		data[row+1][col+1] == "A" &&
		data[row+2][col] == "M" && data[row+2][col+2] == "M" {
		return 1
	}
	return 0
}

// M.S
// .A.
// M.S
func checkMLeft(data [][]string, row, col int) int {
	if data[row][col] == "M" && data[row][col+2] == "S" &&
		data[row+1][col+1] == "A" &&
		data[row+2][col] == "M" && data[row+2][col+2] == "S" {
		return 1
	}
	return 0
}

func findXshapedMAS(data [][]string) int {
	total := 0
	for i, row := range data {
		if i+2 >= len(data) {
			continue
		}
		for j := range row {
			if j+2 >= len(data[i]) {
				continue
			}
			total += checkMTop(data, i, j)
			total += checkMRight(data, i, j)
			total += checkMBottom(data, i, j)
			total += checkMLeft(data, i, j)
		}
	}
	return total
}

func main() {
	data := getInputData()
	fmt.Println("Day 4 - Part 1:", findXMAS(data))
	fmt.Println("Day 4 - Part 2:", findXshapedMAS(data))
}
