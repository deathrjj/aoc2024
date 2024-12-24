package main

import (
	"advent-of-code-2024/helper"
	"fmt"
	"os"
	"path/filepath"
)

func getInputData() [][]int {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	}
	inDayDir := filepath.Base(currentDir) == "day10"
	var data [][]int
	if inDayDir {
		data = helper.ReadInputToInt2DArray("input.txt")
	} else {
		data = helper.ReadInputToInt2DArray("day10/input.txt")
	}
	return data
}

type cell struct {
	r, c int
}

func day10Part1(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	// Gather all height 9 cells
	var nines []cell
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 9 {
				nines = append(nines, cell{r, c})
			}
		}
	}

	// Map each 9 cell to an ID
	nineID := make(map[cell]int)
	for i, pos := range nines {
		nineID[pos] = i
	}

	// reachable[r][c] holds a boolean slice marking which 9s are reachable
	reachable := make([][][]bool, rows)
	for r := 0; r < rows; r++ {
		reachable[r] = make([][]bool, cols)
		for c := 0; c < cols; c++ {
			reachable[r][c] = make([]bool, len(nines))
		}
	}

	// Initialize for height 9 cells
	for i, pos := range nines {
		reachable[pos.r][pos.c][i] = true
	}

	// Directions for neighbors
	dirs := []cell{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// Process heights from 8 down to 0
	for h := 8; h >= 0; h-- {
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] == h {
					// Union of all reachable sets from neighbors of height h+1
					for _, d := range dirs {
						nr, nc := r+d.r, c+d.c
						if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
							continue
						}
						if grid[nr][nc] == h+1 {
							for i := range reachable[nr][nc] {
								if reachable[nr][nc][i] {
									reachable[r][c][i] = true
								}
							}
						}
					}
				}
			}
		}
	}

	// Calculate the sum of scores for all trailheads (height 0)
	totalScore := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				count := 0
				for i := range reachable[r][c] {
					if reachable[r][c][i] {
						count++
					}
				}
				totalScore += count
			}
		}
	}

	return totalScore
}

func day10Part2(grid [][]int) uint64 {
	rows := len(grid)
	cols := len(grid[0])

	ways := make([][]uint64, rows)
	for r := 0; r < rows; r++ {
		ways[r] = make([]uint64, cols)
	}

	// Initialize for height 9 cells
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 9 {
				ways[r][c] = 1
			}
		}
	}

	dirs := []cell{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// Propagate downward: from height 8 to 0
	for h := 8; h >= 0; h-- {
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] == h {
					var total uint64 = 0
					for _, d := range dirs {
						nr, nc := r+d.r, c+d.c
						if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
							continue
						}
						if grid[nr][nc] == h+1 {
							total += ways[nr][nc]
						}
					}
					ways[r][c] = total
				}
			}
		}
	}

	// Sum the ways for all trailheads (height 0)
	var sumRatings uint64 = 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				sumRatings += ways[r][c]
			}
		}
	}

	return sumRatings
}

func main() {
	data := getInputData()
	result := day10Part1(data)
	fmt.Println("Day 10 Part 1:", result)
	result2 := day10Part2(data)
	fmt.Println("Day 10 Part 2:", result2)

}
