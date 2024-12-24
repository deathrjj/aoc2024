package helper

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Position represents a 2D grid position
type Position struct {
	Row, Col int
}

// Grid represents a 2D grid of strings
type Grid [][]string

// ReadInputToGrid reads a file and returns its contents as a Grid
func ReadInputToGrid(filename string) Grid {
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

func ReadInputToInt2DArray(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var nums []int
		for _, c := range line {
			// Convert rune to string then to int
			num, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}
		data = append(data, nums)
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

// CountBlocks counts contiguous blocks of non-"." entries in the disk map
func CountBlocks(diskMap []string) int {
	blocks := 0
	inBlock := false
	for _, v := range diskMap {
		if v != "." {
			if !inBlock {
				blocks++
				inBlock = true
			}
		} else {
			inBlock = false
		}
	}
	return blocks
}

// CountGaps counts contiguous runs of "." in the disk map
func CountGaps(diskMap []string) int {
	gaps := 0
	inGap := false
	for _, v := range diskMap {
		if v == "." {
			if !inGap {
				gaps++
				inGap = true
			}
		} else {
			inGap = false
		}
	}
	return gaps
}

// IsFreeSpace checks if diskMap[start..end] is entirely "."
func IsFreeSpace(diskMap []string, start, end int) bool {
	if start < 0 || end >= len(diskMap) {
		return false
	}
	for i := start; i <= end; i++ {
		if diskMap[i] != "." {
			return false
		}
	}
	return true
}

// CollectBlocks identifies all contiguous file-block segments and returns them as start, end, size
func CollectBlocks(diskMap []string) [][]int {
	var blocks [][]int
	blockStart := -1
	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] != "." {
			if blockStart == -1 {
				blockStart = i
			}
		} else if blockStart != -1 {
			blocks = append(blocks, []int{blockStart, i - 1, i - blockStart})
			blockStart = -1
		}
	}
	if blockStart != -1 {
		blocks = append(blocks, []int{blockStart, len(diskMap) - 1, len(diskMap) - blockStart})
	}
	return blocks
}

// CanFit checks if a block of a given size can fit starting at a given position
func CanFit(start, size int, diskMap []string) bool {
	if start+size > len(diskMap) {
		return false
	}
	for i := start; i < start+size; i++ {
		if diskMap[i] != "." {
			return false
		}
	}
	return true
}

// InitLogger initializes the logger and returns a file pointer
func InitLogger(filePath string) (*os.File, error) {
	// If the file exists, remove it
	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			return nil, err
		}
	}

	// Create/open a new file for logging
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) // #nosec G304
	if err != nil {
		return nil, err
	}

	log.SetOutput(logFile)
	return logFile, nil
}

// ParseDiskMap parses the input string into a disk map representation
func ParseDiskMap(input string) []string {
	var diskMap []string
	length := len(input)

	// Iterate through the input in pairs
	for i := 0; i < length; i += 2 {
		fileSize := int(input[i] - '0') // Convert char to int
		freeSpace := 0

		// Handle odd-length input (last digit as free space if no pair)
		if i+1 < length {
			freeSpace = int(input[i+1] - '0')
		}

		// Add file blocks
		for j := 0; j < fileSize; j++ {
			diskMap = append(diskMap, fmt.Sprintf("%d", i/2))
		}

		// Add free space
		for j := 0; j < freeSpace; j++ {
			diskMap = append(diskMap, ".")
		}
	}

	return diskMap
}

// FileBlock represents a contiguous file block in the disk map
type FileBlock struct {
	FileID int
	Start  int
	End    int
}

// IdentifyFiles returns a slice of FileBlocks representing contiguous file segments
func IdentifyFiles(diskMap []string) []FileBlock {
	var files []FileBlock
	inFile := false
	currentFileID := -1
	currentStart := -1

	for i, v := range diskMap {
		if v == "." {
			// If we were in a file, close it off
			if inFile {
				files = append(files, FileBlock{
					FileID: currentFileID,
					Start:  currentStart,
					End:    i - 1,
				})
				inFile = false
			}
		} else {
			// We have a file block
			fid, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Invalid file ID '%s' in disk map: %v", v, err)
			}
			if !inFile {
				inFile = true
				currentFileID = fid
				currentStart = i
			} else if fid != currentFileID {
				// Encountered a new file unexpectedly
				files = append(files, FileBlock{
					FileID: currentFileID,
					Start:  currentStart,
					End:    i - 1,
				})
				currentFileID = fid
				currentStart = i
			}
		}
	}
	// Close off last file if we ended in one
	if inFile {
		files = append(files, FileBlock{
			FileID: currentFileID,
			Start:  currentStart,
			End:    len(diskMap) - 1,
		})
	}

	return files
}

// ReadInputAsString reads a file and returns its contents as a trimmed string
func ReadInputAsString(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}
	return strings.TrimSpace(string(file)), nil
}
