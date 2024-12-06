package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData() ([][]int, [][]int) {
	rules := [][]int{}
	updates := [][]int{}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			// Rule line
			parts := strings.Split(line, "|")
			rule := make([]int, 2)
			rule[0], _ = strconv.Atoi(parts[0])
			rule[1], _ = strconv.Atoi(parts[1])
			rules = append(rules, rule)
		} else if line != "" {
			// Value line
			parts := strings.Split(line, ",")
			update := make([]int, len(parts))
			for i, p := range parts {
				update[i], _ = strconv.Atoi(p)
			}
			updates = append(updates, update)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return rules, updates
}

func getRule(rules [][]int, value int) []int {
	ruleValues := []int{}
	for _, rule := range rules {
		if rule[0] == value {
			ruleValues = append(ruleValues, rule[1])
		}
	}
	return ruleValues
}

func checkRule(values []int, rule int) bool {
	for _, value := range values {
		if value == rule {
			return false
		}
	}
	return true
}

func checkOrder(values []int, rules [][]int) int {
	for j, value := range values {
		ruleValues := getRule(rules, value)
		before := values[:j]
		for _, rule := range ruleValues {
			if !checkRule(before, rule) {
				return rule
			}
		}
	}
	return -1
}

func getMiddleValue(update []int) int {
	length := len(update)
	middleIndex := (length - 1) / 2
	return update[middleIndex]
}

func day05_1(rules [][]int, values [][]int) int {
	total := 0
	for _, update := range values {
		if checkOrder(update, rules) == -1 {
			total += getMiddleValue(update)
		}
	}
	return total
}

func moveBrokenRuleValueToEnd(update []int, brokenRule int) []int {
	updated := []int{}
	for _, value := range update {
		if value != brokenRule {
			updated = append(updated, value)
		}
	}
	updated = append(updated, brokenRule)
	return updated
}

func moveUntilCorrectOrder(update []int, rules [][]int) []int {
	brokenRule := checkOrder(update, rules)
	if brokenRule == -1 {
		return update
	}
	newUpdate := moveBrokenRuleValueToEnd(update, brokenRule)
	return moveUntilCorrectOrder(newUpdate, rules)
}

func day05_2(rules [][]int, values [][]int) int {
	total := 0
	for _, update := range values {
		if checkOrder(update, rules) == -1 {
			continue
		}
		ordered := moveUntilCorrectOrder(update, rules)
		total += getMiddleValue(ordered)
	}
	return total
}

func main() {
	rules, updates := getInputData()
	fmt.Println("Day 05 - Part 1: Total of correctly ordered update's middle values =", day05_1(rules, updates))
	fmt.Println("Day 05 - Part 2: Total of incorrectly ordered update's corrected middle values =", day05_2(rules, updates))
}
