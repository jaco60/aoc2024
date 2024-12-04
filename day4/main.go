package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// checkCoords tests if (x,y) is in the grid
func checkCoords(grid []string, x, y int) bool {
	return (x >= 0 && y >= 0) && (x < len(grid[0]) && y < len(grid))
}

// findMax checks if "XMAS" is found in grid, from start coords
// in dir direction
func findXmas(grid []string, startX, startY int, dir []int) bool {
	// We already know that grid[startX][startY] == 'X'
	// We have to find MAS in direction dir
	xCurr := startX + dir[0]
	yCurr := startY + dir[1]
	numChar := 0
	for checkCoords(grid, xCurr, yCurr) {
		c := grid[xCurr][yCurr]
		switch numChar {
		case 0: // Should be a 'M'
			if c != 'M' {
				return false
			}
		case 1: // Should be a 'A'
			if c != 'A' {
				return false
			}
		case 2: // Should be a 'S'
			return c == 'S'
		}
		xCurr += dir[0]
		yCurr += dir[1]
		numChar++
	}
	return false
}

// nbXmas return the number of "XMAS" in the grid, starting from start in any direction
func nbXmas(grid []string, startX, startY int) int {
	var directions = [][]int{
		{0, -1},  // Up
		{0, 1},   // Down
		{-1, 0},  // Left
		{1, 0},   // Right
		{-1, -1}, // Diag Up-Left
		{1, 1},   // Diag Down-Right
		{-1, 1},  // Diag Up-Right
		{1, -1},  // Diag Down-Left
	}
	count := 0
	for _, dir := range directions {
		if grid[startX][startY] == 'X' && findXmas(grid, startX, startY, dir) {
			count++
		}
	}
	return count
}

func solution1(grid []string) int {
	count := 0
	xMax := len(grid[0])
	yMax := len(grid)
	for startX := range xMax {
		for startY := range yMax {
			count += nbXmas(grid, startX, startY)
		}
	}
	return count
}

func findMas(grid []string, x, y int) bool {
	// We already know that grid[x][y] == 'A'
	// Searching for MS...
	return isMs(grid[x-1][y-1], grid[x+1][y+1]) &&
		isMs(grid[x-1][y+1], grid[x+1][y-1])
}

func isMs(c1, c2 byte) bool {
	return (c1 == 'M' && c2 == 'S') ||
		(c1 == 'S' && c2 == 'M')
}

func solution2(grid []string) int {
	count := 0
	xMax := len(grid[0])
	yMax := len(grid)
	// 'A' cannot be in line 0/col 0...
	for startX := 1; startX < xMax-1; startX++ {
		for startY := 1; startY < yMax-1; startY++ {
			if grid[startX][startY] == 'A' && findMas(grid, startX, startY) {
				count++
			}
		}
	}
	return count
}

func main() {
	grid := strings.Fields(input)
	fmt.Println(solution1(grid))
	fmt.Println(solution2(grid))
}
