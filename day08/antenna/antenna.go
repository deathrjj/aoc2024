package antenna

import (
	"advent-of-code-2024/helper"
	"math"
)

const (
	targetRatio = 2.0
	tolerance   = 0.01
)

// FindAntinodes finds all antinodes in the antenna map (Part 1)
func FindAntinodes(data helper.Grid) int {
	processed := make(map[string]bool)          // Track processed antenna characters
	antinodes := make(map[helper.Position]bool) // Track unique antinode positions

	// Find all antenna positions
	antennaPositions := data.FindAll(func(s string) bool {
		return s != "." && s != "#" && !processed[s]
	})

	// Process each antenna
	for _, pos := range antennaPositions {
		char := data.Get(pos)
		processed[char] = true
		findMatchingAntennas(data, pos, antinodes)
	}

	return len(antinodes)
}

// FindAntinodesWithResonance finds all antinodes considering resonant harmonics (Part 2)
func FindAntinodesWithResonance(data helper.Grid) int {
	processed := make(map[string]bool)          // Track processed antenna characters
	antinodes := make(map[helper.Position]bool) // Track unique antinode positions

	// Find all antenna positions
	antennaPositions := data.FindAll(func(s string) bool {
		return s != "." && s != "#" && !processed[s]
	})

	// Process each antenna
	for _, pos := range antennaPositions {
		char := data.Get(pos)
		processed[char] = true
		findMatchingAntennasWithResonance(data, pos, antinodes)
	}

	return len(antinodes)
}

// findMatchingAntennas finds all matching antennas and their antinodes (Part 1)
func findMatchingAntennas(data helper.Grid, pos helper.Position, antinodes map[helper.Position]bool) {
	antennaChar := data.Get(pos)

	// Find all matching antennas
	matches := data.FindAll(func(s string) bool {
		return s == antennaChar
	})

	// Process each matching antenna
	for _, match := range matches {
		if match == pos {
			continue
		}

		// Calculate distance between antennas
		distance := pos.DistanceTo(match)

		// Skip if distance is 0
		if distance == 0 {
			continue
		}

		// Search for antinodes in a reasonable area around the antennas
		minRow := helper.Min(pos.Row, match.Row) - int(distance*2)
		maxRow := helper.Max(pos.Row, match.Row) + int(distance*2)
		minCol := helper.Min(pos.Col, match.Col) - int(distance*2)
		maxCol := helper.Max(pos.Col, match.Col) + int(distance*2)

		// Ensure search bounds are within grid
		minRow = helper.Max(minRow, 0)
		maxRow = helper.Min(maxRow, data.Height()-1)
		minCol = helper.Max(minCol, 0)
		maxCol = helper.Min(maxCol, data.Width()-1)

		// Search for points that satisfy the distance requirements
		for r := minRow; r <= maxRow; r++ {
			for c := minCol; c <= maxCol; c++ {
				candidate := helper.Position{Row: r, Col: c}

				// Skip if point is at same position as either antenna
				if candidate == pos || candidate == match {
					continue
				}

				// Check if point is collinear with antennas
				if !candidate.IsCollinearWith(pos, match, tolerance) {
					continue
				}

				d1 := candidate.DistanceTo(pos)
				d2 := candidate.DistanceTo(match)

				// Check both possible ratios (either distance could be the larger one)
				ratio1 := d1 / d2
				ratio2 := d2 / d1

				if math.Abs(ratio1-2.0) < tolerance || math.Abs(ratio2-2.0) < tolerance {
					antinodes[candidate] = true
				}
			}
		}
	}
}

// findMatchingAntennasWithResonance finds all matching antennas and their antinodes (Part 2)
func findMatchingAntennasWithResonance(data helper.Grid, pos helper.Position, antinodes map[helper.Position]bool) {
	antennaChar := data.Get(pos)

	// Find all matching antennas
	matches := data.FindAll(func(s string) bool {
		return s == antennaChar
	})

	// If there's more than one antenna of this frequency, the antenna positions themselves are antinodes
	if len(matches) > 1 {
		for _, match := range matches {
			antinodes[match] = true
		}
	}

	// Process each pair of matching antennas
	for i, antenna1 := range matches {
		for j := i + 1; j < len(matches); j++ {
			antenna2 := matches[j]

			// Search the entire grid for collinear points
			for r := 0; r < data.Height(); r++ {
				for c := 0; c < data.Width(); c++ {
					candidate := helper.Position{Row: r, Col: c}

					// Check if point is collinear with antennas
					if candidate.IsCollinearWith(antenna1, antenna2, tolerance) {
						antinodes[candidate] = true
					}
				}
			}
		}
	}
}
