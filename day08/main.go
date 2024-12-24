package main

import (
	"advent-of-code-2024/day08/antenna"
	"advent-of-code-2024/helper"
	"fmt"
)

func main() {
	// Read and parse input
	grid := helper.ReadInputToGrid("day08/input.txt")

	// Part 1: Find antinodes based on distance ratios
	antinodeCount := antenna.FindAntinodes(grid)
	fmt.Printf("Day 08 Part 1: Total Antinodes: %d\n", antinodeCount)

	// Part 2: Find antinodes considering resonant harmonics
	resonantCount := antenna.FindAntinodesWithResonance(grid)
	fmt.Printf("Day 08 Part 2: Total Antinodes with Resonance: %d\n", resonantCount)
}
