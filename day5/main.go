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

func isCorrect(updates, rules []string) bool {
	for i := 0; i < len(updates)-1; i++ {
		first := updates[i]
		for j := i + 1; j < len(updates); j++ {
			second := updates[j]
			if !slices.Contains(rules, first+"|"+second) {
				return false
			}
		}
	}
	return true
}

func solution1(updates, rules []string) int {
	res := 0
	for _, update := range updates {
		pages := strings.Split(update, ",")
		if isCorrect(pages, rules) {
			mid, _ := strconv.Atoi(string(pages[len(pages)/2]))
			res += mid
		}
	}
	return res
}

func repair(pages, rules []string) []string {
	return slices.SortedFunc(slices.Values(pages), func(a, b string) int {
		if slices.Contains(rules, a+"|"+b) {
			return -1
		}
		if slices.Contains(rules, b+"|"+a) {
			return 1
		}
		return 0
	})
}

func solution2(updates, rules []string) int {
	res := 0
	for _, update := range updates {
		pages := strings.Split(update, ",")
		if !isCorrect(pages, rules) {
			mid, _ := strconv.Atoi(repair(pages, rules)[len(pages)/2])
			res += mid
		}
	}
	return res
}

func main() {
	data := strings.Split(input, "\n\n")
	rules := strings.Fields(data[0])
	updates := strings.Fields(data[1])

	fmt.Println(solution1(updates, rules))
	fmt.Println(solution2(updates, rules))
}
