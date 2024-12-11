package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"advent-of-code-2024/helper"
)

// Block represents a contiguous section of the disk
type Block struct {
	fileID int
	length int
}

// CompactDisk moves file blocks from the end to the leftmost free space
func CompactDisk(diskMap []string) []string {
	for i := len(diskMap) - 1; i >= 0; i-- { // Start from the rightmost block
		if diskMap[i] != "." { // If it's a file block
			// Find the leftmost free space
			for j := 0; j < i; j++ {
				if diskMap[j] == "." {
					// Move the block to the free space
					diskMap[j] = diskMap[i]
					diskMap[i] = "."
					break // Stop searching for free space
				}
			}
		}
	}
	return diskMap
}

// CalculateChecksum computes the checksum of the disk map
func CalculateChecksum(diskMap []string) int {
	checksum := 0
	for idx, block := range diskMap {
		if block != "." {
			fileID, err := strconv.Atoi(block)
			if err != nil {
				log.Fatalf("Invalid file ID '%s' in disk map: %v", block, err)
			}
			checksum += idx * fileID
		}
	}
	return checksum
}

func CompactDiskPart2(diskMap []string) []string {
	files := helper.IdentifyFiles(diskMap)
	// Sort files by file ID descending
	sort.Slice(files, func(i, j int) bool {
		return files[i].FileID > files[j].FileID
	})

	for _, f := range files {
		fileLen := f.End - f.Start + 1
		// Attempt to find a free space span to the left of f.Start that can hold fileLen blocks
		if f.Start > 0 {
			if freeStart := findFreeSpaceToTheLeft(diskMap, 0, f.Start-1, fileLen); freeStart != -1 {
				// Move the entire file
				moveFile(diskMap, f.Start, f.End, freeStart)
			}
		}
	}

	return diskMap
}

// findFreeSpaceToTheLeft searches within diskMap[start..end] for a contiguous run of '.' of length neededSize.
// Returns the start index of that run if found, otherwise -1.
func findFreeSpaceToTheLeft(diskMap []string, start, end, neededSize int) int {
	freeCount := 0
	freeStart := -1
	for i := start; i <= end; i++ {
		if diskMap[i] == "." {
			if freeStart == -1 {
				freeStart = i
			}
			freeCount++
			if freeCount == neededSize {
				// Found a suitable span
				return freeStart
			}
		} else {
			freeCount = 0
			freeStart = -1
		}
	}
	return -1
}

// moveFile moves the file blocks from [oldStart..oldEnd] to [newStart..newStart+(length-1)]
// and replaces the old file positions with '.'.
func moveFile(diskMap []string, oldStart, oldEnd, newStart int) {
	length := oldEnd - oldStart + 1
	// Copy file content
	for i := 0; i < length; i++ {
		diskMap[newStart+i] = diskMap[oldStart+i]
	}
	// Clear old location
	for i := oldStart; i <= oldEnd; i++ {
		diskMap[i] = "."
	}
}

func main() {
	// Determine if we're running from the day directory or root
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	// Check if we're in the day09 directory
	inDayDir := filepath.Base(currentDir) == "day09"

	// Adjust log directory path based on current location
	logDir := "logs"
	if !inDayDir {
		logDir = "day09/logs"
	}

	// Create logs directory
	if err := os.MkdirAll(logDir, 0750); err != nil {
		log.Fatalf("Failed to create logs directory: %v", err)
	}

	// Create timestamped log file
	timestamp := time.Now().Format("2006-01-02-150405")
	logPath := filepath.Join(logDir, fmt.Sprintf("debug-%s.log", timestamp))

	// Initialize logger
	logFile, err := helper.InitLogger(logPath)
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer logFile.Close()

	// Try to read input file from either path
	var input string
	if inDayDir {
		input, err = helper.ReadInputAsString("input.txt")
	} else {
		input, err = helper.ReadInputAsString("day09/input.txt")
	}
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Parse the disk map
	diskMap := helper.ParseDiskMap(input)
	log.Printf("Parsed Disk Map: %v\n", diskMap)

	// Part 1
	compactDiskPart1 := CompactDisk(diskMap)
	checksumPart1 := CalculateChecksum(compactDiskPart1)
	log.Printf("Part 1 Compacted Disk: %v\n", compactDiskPart1)
	log.Printf("Part 1 Checksum Calculation: %v\n", checksumPart1)
	fmt.Printf("Day 09 Part 1: Checksum = %d\n", checksumPart1)

	// Part 2
	compactDiskPart2 := CompactDiskPart2(helper.ParseDiskMap(input)) // Re-parse to avoid mutating diskMap
	checksumPart2 := CalculateChecksum(compactDiskPart2)
	log.Printf("Part 2 Compacted Disk: %v\n", compactDiskPart2)
	log.Printf("Part 2 Checksum Calculation: %v\n", checksumPart2)
	fmt.Printf("Day 09 Part 2: Checksum = %d\n", checksumPart2)
}
