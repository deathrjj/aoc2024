package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInputData() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func isNumber(char string) bool {
	return char >= "0" && char <= "9"
}

func calculateMuls(data string) int {
	total := 0
	for i := 0; i < len(data)-4; i++ { // Need at least "mul()" (4 chars)
		if data[i:i+4] == "mul(" {
			// Found start of potential mul expression
			j := i + 4 // Position after "mul("

			// Get first number
			num1Str := ""
			for j < len(data) && isNumber(string(data[j])) {
				num1Str += string(data[j])
				j++
			}

			// Check for comma
			if j < len(data) && data[j] == ',' {
				j++ // Move past comma

				// Get second number
				num2Str := ""
				for j < len(data) && isNumber(string(data[j])) {
					num2Str += string(data[j])
					j++
				}

				// Check for closing parenthesis
				if j < len(data) && data[j] == ')' {
					// Valid mul expression found
					num1, _ := strconv.Atoi(num1Str)
					num2, _ := strconv.Atoi(num2Str)
					total += num1 * num2
				}
			}
		}
	}
	return total
}

func calcEnabledMuls(data string) int {
	total := 0

	// Split the data on "do()" to get all the possible expressions. Enabled by default so the expression at index 0 is included
	splits := strings.Split(data, "do()")
	for _, split := range splits {
		// Get the part of the expression before the first don't()
		enabledExpression := strings.Split(split, "don't()")[0]

		// Calculate the muls in the enabled expression and add to total
		total += calculateMuls(enabledExpression)
	}
	return total
}

func main() {
	data := getInputData()
	fmt.Println("Day 3, Part 1:", calculateMuls(data))
	fmt.Println("Day 3, Part 2:", calcEnabledMuls(data))
}
