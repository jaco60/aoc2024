package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// No error handling: input is supposed correct...

func leftsAndRights(lines []string) ([]int, []int) {
	var lefts, rights []int
	for _, pair := range lines {
		line := strings.Fields(pair)
		left, _ := strconv.Atoi(line[0])
		right, _ := strconv.Atoi(line[1])
		lefts = append(lefts, left)
		rights = append(rights, right)
	}
	return lefts, rights
}

func solution1(lines []string) int {
	lefts, rights := leftsAndRights(lines)
	slices.Sort(lefts)
	slices.Sort(rights)
	sum := 0
	for i, left := range lefts {
		sum += int(math.Abs(float64(rights[i] - left)))
	}
	return sum
}

func solution2(lines []string) int {
	lefts, rights := leftsAndRights(lines)
	counts := make(map[int]int)
	// Counting each values occurrences in rights
	for _, right := range rights {
		counts[right] += 1
	}
	// Sum up left values scores
	score := 0
	for _, left := range lefts {
		score += left * counts[left]
	}
	return score
}

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(solution1(lines))
	fmt.Println(solution2(lines))
}
