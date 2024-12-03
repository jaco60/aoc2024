package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func solution1(input string) int {
	res := 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	mults := re.FindAllString(input, -1)
	for _, mult := range mults {
		args := strings.Fields(re.ReplaceAllString(mult, "$1 $2"))
		arg1, _ := strconv.Atoi(string(args[0]))
		arg2, _ := strconv.Atoi(string(args[1]))
		res += arg1 * arg2
	}
	return res
}

func solution2(input string) int {
	mults := ""    // To reuse solution1...
    re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)
	matches := re.FindAllString(input, -1)
	ok := true
	for _, match := range matches {
		if match == "don't()" {
			ok = false
			continue
		}
		if match == "do()" {
			ok = true
			continue
		}
		if ok {
			mults = mults + match
		}
	}
	return solution1(mults)
}

func main() {
	fmt.Println(solution1(input))
	fmt.Println(solution2(input))

}
