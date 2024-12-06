package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func isIn(grid []string, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

// whereIsGuard returns initial guard coords
func whereIsGuard(grid []string) (int, int) {
	for x := range len(grid) {
		y := strings.IndexAny(grid[x], "^v><")
		if y != -1 {
			return x, y
		}
	}
	return 0, 0 // never reached
}

// nextCoord returns the next coords of the guard in its current direction
func nextCoords(x, y int, dir byte) (int, int) {
	switch dir {
	case '^':

		return x - 1, y
	case 'v':

		return x + 1, y
	case '<':
		return x, y - 1
	}
	return x, y + 1
}

// rotate turns the guard to the right
func rotate(dir byte) byte {
	var res byte
	switch dir {
	case '^':
		res = '>'
	case 'v':
		res = '<'
	case '<':
		res = '^'
	case '>':
		res = 'v'
	}
	return res
}

func checkBlock(grid []string, x, y int) bool {
	return grid[x][y] == '#'
}

func solution1(grid []string) map[string]struct{} {
	xGuard, yGuard := whereIsGuard(grid)
	guardDir := grid[xGuard][yGuard]
	path := map[string]struct{}{}
	var nextX, nextY int

	for isIn(grid, xGuard, yGuard) {
		path[fmt.Sprintf("(%d,%d)", xGuard, yGuard)] = struct{}{}
		nextX, nextY = nextCoords(xGuard, yGuard, guardDir)
		if isIn(grid, nextX, nextY) && checkBlock(grid, nextX, nextY) {
			guardDir = rotate(guardDir)
		} else {
			xGuard, yGuard = nextX, nextY
		}
	}
	return path
}

func main() {
	grid := strings.Fields(input)
	posX, posY := whereIsGuard(grid)
	fmt.Printf("starting from (%d, %d)\n", posX, posY)
	fmt.Println(len(solution1(grid)))
}
