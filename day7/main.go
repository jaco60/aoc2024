package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func mapToInts(input string) []int {
	strs := strings.Fields(input)
	res := make([]int, len(strs))
	for i, str := range strs {
		res[i], _ = strconv.Atoi(str)
	}
	return res
}

func ops(n int) []string {
	if n == 1 {
		return []string{"+", "*"}
	}
	combinations := []string{}
	for _, comb := range ops(n - 1) {
		combinations = append(combinations, comb+"+")
		combinations = append(combinations, comb+"*")
	}
	return combinations
}

func ops2(n int) []string {
	if n == 1 {
		return []string{"+", "*", "|"}
	}
	combinations := []string{}
	for _, comb := range ops2(n - 1) {
		combinations = append(combinations, comb+"+")
		combinations = append(combinations, comb+"*")
		combinations = append(combinations, comb+"|")
	}
	return combinations
}

func solve(expected int, values []int, ops []string) bool {
	for _, seqOps := range ops {
		res := values[0]
		for i, op := range seqOps {
			if op == '+' {
				res += values[i+1]
			} else if op == '*' {
				res *= values[i+1]
			} else if op == '|' {
				res, _ = strconv.Atoi(fmt.Sprintf("%d%d", res, values[i+1]))
			}
		}
		if res == expected {
			return true
		}
	}
	return false
}

func solution(lines []string, num int) int {
	var (
		res        = 0
		expected   int
		values     []int
		operations []string
	)

	for _, line := range lines {
		tmp := strings.Split(line, ":")
		expected, _ = strconv.Atoi(tmp[0])
		values = mapToInts(tmp[1])
		if num == 1 {
			operations = ops(len(values) - 1)
		} else if num == 2 {
			operations = ops2(len(values) - 1)
		}
		if solve(expected, values, operations) {
			res += expected
		}
	}
	return res
}

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(solution(lines, 1))
	fmt.Println(solution(lines, 2))

}
