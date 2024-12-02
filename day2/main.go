package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func convertToInt(report string) []int {
	levels := strings.Fields(report)
	res := make([]int, len(levels))
	for i, level := range levels {
		res[i], _ = strconv.Atoi(level)
	}
	return res
}

func isSafe(report string) bool {
	levels := convertToInt(report)

	ascending := levels[0] < levels[1]
	if ascending {
		for i := range len(levels) - 1 {
			diff := levels[i+1] - levels[i]
			if diff <= 0 || diff > 3 {
				return false
			}
		}
	} else {
		for i := range len(levels) - 1 {
			diff := levels[i] - levels[i+1]
			if diff <= 0 || diff > 3 {
				return false
			}
		}
	}
	return true
}

func solution1(lines []string) int {
	safes := 0
	for _, report := range lines {
		if isSafe(report) {
			safes++
		}
	}
	return safes
}

func solution2(lines []string) int {
	safes := 0
	for _, report := range lines {
		levels := strings.Fields(report)
		for i := range len(levels) {
			// Remove levels[i] and test the safety of levels without this elt
			temp := slices.Clone(levels)
			temp = append(temp[:i], temp[i+1:]...)
			if isSafe(strings.Join(temp, " ")) {
				safes++
				break
			}
		}
	}
	return safes
}

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(solution1(lines))
	fmt.Println(solution2(lines))
}
