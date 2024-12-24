package main

import (
	"fmt"
	"time"
)

var pow10 = [19]int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000,
	10000000000, 100000000000, 1000000000000, 10000000000000,
	100000000000000, 1000000000000000, 10000000000000000,
	100000000000000000, 1000000000000000000}

func applyRules(value int) []int {
	if value == 0 {
		return []int{1}
	}
	if value < pow10[1] {
		return []int{value * 2024}
	}
	if value < pow10[2] {
		return []int{value / pow10[1], value % pow10[1]}
	}
	if value < pow10[3] {
		return []int{value * 2024}
	}
	if value < pow10[4] {
		return []int{value / pow10[2], value % pow10[2]}
	}
	if value < pow10[5] {
		return []int{value * 2024}
	}
	if value < pow10[6] {
		return []int{value / pow10[3], value % pow10[3]}
	}
	if value < pow10[7] {
		return []int{value * 2024}
	}
	if value < pow10[8] {
		return []int{value / pow10[4], value % pow10[4]}
	}
	if value < pow10[9] {
		return []int{value * 2024}
	}
	if value < pow10[10] {
		return []int{value / pow10[5], value % pow10[5]}
	}
	if value < pow10[11] {
		return []int{value * 2024}
	}
	if value < pow10[12] {
		return []int{value / pow10[6], value % pow10[6]}
	}
	if value < pow10[13] {
		return []int{value * 2024}
	}
	if value < pow10[14] {
		return []int{value / pow10[7], value % pow10[7]}
	}
	if value < pow10[15] {
		return []int{value * 2024}
	}
	if value < pow10[16] {
		return []int{value / pow10[8], value % pow10[8]}
	}
	if value < pow10[17] {
		return []int{value * 2024}
	}
	if value < pow10[18] {
		return []int{value / pow10[9], value % pow10[9]}
	}
	return []int{value * 2024}
}

func main() {

	totalIterations := 75
	results := make([][]int, totalIterations+1)
	results[0] = []int{773, 79858, 0, 71, 213357, 2937, 1, 3998391}

	fmt.Println("Iteration 0 ->", len(results[0]), "values")

	nextMap := make(map[int][]int)
	for iteration := 1; iteration <= totalIterations; iteration++ {
		iterationStart := time.Now()
		previousValues := &results[iteration-1]
		currentValues := &results[iteration]
		for _, previousValue := range *previousValues {
			if _, exists := nextMap[previousValue]; !exists {
				nextMap[previousValue] = applyRules(previousValue)
			}
			*currentValues = append(*currentValues, nextMap[previousValue]...)
		}
		fmt.Println("Iteration", iteration, "->", len(*currentValues), "values", "( Calculated in ", time.Since(iterationStart), ")")
	}
}
