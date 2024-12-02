package main

import (
	_ "embed"
	"fmt"
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

// func isSafe(report string) bool {
// 	levels := convertToInt(report)
	
// 	// Check if the report is ascending or descending
// 	if !slices.IsSorted(levels) &&
// 	   !slices.IsSortedFunc(levels, func(a, b int) int {
// 					return b - a
// 				}) {
// 					return false
// 				}
// 	// Any two adjacent levels differ by at least one and at most three
// 	for i, level := range levels[:len(levels)-1] {
// 		diff := int(math.Abs(float64(level - levels[i+1])))
// 		if  diff == 0 || diff > 3 {
// 			return false
// 		}
// 	}
// 				return true
// }

func isSafe(report string) bool {
	levels := convertToInt(report)
	
	ascending := levels[0] < levels[1]
	if ascending {
		for i := range len(levels) - 1 {
			diff := levels[i+1] - levels[i]
			if  diff <= 0 || diff > 3 {
				return false
			}
		}
	} else {
		for i := range len(levels) - 1 {
			diff := levels[i] - levels[i+1]
			if  diff <= 0 || diff > 3 {
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

func lineMinusOne(line []string, i int) []string {
	// Return line minus line[i]
	res := []string{}
	for j := range len(line) {
		if j != i {
			res = append(res, line[j])
		}
	}
	return res
}

func solution2(lines []string) int {
	safes := 0		
	for _, report := range lines {
		levels := strings.Fields(report)
		for i := range len(levels) {
			temp := lineMinusOne(levels, i)
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