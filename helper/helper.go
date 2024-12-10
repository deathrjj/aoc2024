package helper

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// Position represents a 2D grid position
type Position struct {
	Row, Col int
}

// Grid represents a 2D grid of strings
type Grid [][]string

// ReadInput reads a file and returns its contents as a Grid
func ReadInput(filename string) Grid {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data Grid
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

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Distance calculates the Euclidean distance between two points
func Distance(x1, y1, x2, y2 int) float64 {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	return math.Sqrt(dx*dx + dy*dy)
}

// ManhattanDistance calculates the Manhattan distance between two points
func ManhattanDistance(x1, y1, x2, y2 int) int {
	return Abs(x2-x1) + Abs(y2-y1)
}

// Abs returns the absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// IsCollinear checks if three points are collinear within a tolerance
func IsCollinear(x1, y1, x2, y2, x3, y3 float64, tolerance float64) bool {
	// Calculate vectors from first point to others
	v1x := x2 - x1
	v1y := y2 - y1
	v2x := x3 - x1
	v2y := y3 - y1

	// Cross product should be close to 0 for collinear points
	// Normalize by the distance to handle larger numbers
	dist := math.Sqrt(v1x*v1x + v1y*v1y)
	if dist == 0 {
		return false
	}
	crossProduct := (v1x*v2y - v1y*v2x) / dist
	return math.Abs(crossProduct) < tolerance
}

// HasDistanceRatio checks if a point has the required distance ratio to two other points
func HasDistanceRatio(x1, y1, x2, y2, px, py int, targetRatio float64, tolerance float64) bool {
	d1 := Distance(px, py, x1, y1)
	d2 := Distance(px, py, x2, y2)

	if d1 == 0 || d2 == 0 {
		return false
	}

	ratio := d2 / d1
	return math.Abs(ratio-targetRatio) < tolerance
}

// Grid methods

// Height returns the height of the grid
func (g Grid) Height() int {
	return len(g)
}

// Width returns the width of the grid
func (g Grid) Width() int {
	if len(g) == 0 {
		return 0
	}
	return len(g[0])
}

// IsInBounds checks if a position is within the grid bounds
func (g Grid) IsInBounds(pos Position) bool {
	return pos.Row >= 0 && pos.Row < g.Height() && pos.Col >= 0 && pos.Col < g.Width()
}

// Get returns the value at a position in the grid
func (g Grid) Get(pos Position) string {
	if !g.IsInBounds(pos) {
		return ""
	}
	return g[pos.Row][pos.Col]
}

// FindAll returns all positions in the grid that match a predicate
func (g Grid) FindAll(predicate func(string) bool) []Position {
	var positions []Position
	for row := range g {
		for col := range g[row] {
			if predicate(g[row][col]) {
				positions = append(positions, Position{Row: row, Col: col})
			}
		}
	}
	return positions
}

// Position methods

// Add returns a new position offset by dr, dc
func (p Position) Add(dr, dc int) Position {
	return Position{Row: p.Row + dr, Col: p.Col + dc}
}

// DistanceTo returns the Euclidean distance to another position
func (p Position) DistanceTo(other Position) float64 {
	return Distance(p.Row, p.Col, other.Row, other.Col)
}

// ManhattanDistanceTo returns the Manhattan distance to another position
func (p Position) ManhattanDistanceTo(other Position) int {
	return ManhattanDistance(p.Row, p.Col, other.Row, other.Col)
}

// IsCollinearWith checks if this position is collinear with two other positions
func (p Position) IsCollinearWith(p1, p2 Position, tolerance float64) bool {
	// For grid positions, Row is Y and Col is X
	return IsCollinear(
		float64(p1.Row), float64(p1.Col),
		float64(p2.Row), float64(p2.Col),
		float64(p.Row), float64(p.Col),
		tolerance,
	)
}

// StringsToInts converts a slice of strings to a slice of ints
func StringsToInts(strings []string) ([]int, error) {
	var numbers []int
	for _, s := range strings {
		var num int
		if _, err := fmt.Sscanf(s, "%d", &num); err != nil {
			return nil, fmt.Errorf("failed to parse number '%s': %v", s, err)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

// IsEven returns true if the number is even
func IsEven(n int) bool {
	return n%2 == 0
}
