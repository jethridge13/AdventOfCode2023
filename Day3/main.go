package main

import (
	"fmt"
	"strings"

	"github.com/jethridge13/AdventOfCode2023/util"
)

type Key struct {
	X, Y int
}

type Gear struct {
	X, Y    int
	Numbers []int
}

func findSymbolsInGrid(grid [][]string) map[Key]string {
	m := make(map[Key]string)
	for i, row := range grid {
		for j, cell := range row {
			if !util.IsStringDigit(cell) && cell != "." {
				key := &Key{X: i, Y: j}
				m[*key] = string(cell)
			}
		}
	}
	return m
}

func getNumbersAdjacentToSymbols(symbols map[Key]string, grid [][]string) map[Key]string {
	m := make(map[Key]string)
	var checkGrid = func(x, y int) {
		if util.IsStringDigit(grid[x][y]) {
			k := Key{X: x, Y: y}
			if _, ok := m[k]; !ok {
				m[k] = grid[x][y]
			}
		}
	}
	for key := range symbols {
		// NW
		x := key.X - 1
		y := key.Y - 1
		if x >= 0 && y >= 0 {
			checkGrid(x, y)
		}
		// N
		x = key.X - 1
		y = key.Y
		if x >= 0 {
			checkGrid(x, y)
		}
		// NE
		x = key.X - 1
		y = key.Y + 1
		if x >= 0 && y < len(grid[x]) {
			checkGrid(x, y)
		}
		// E
		x = key.X
		y = key.Y + 1
		if y < len(grid[x]) {
			checkGrid(x, y)
		}
		// SE
		x = key.X + 1
		y = key.Y + 1
		if x < len(grid) && y < len(grid[x]) {
			checkGrid(x, y)
		}
		// S
		x = key.X + 1
		y = key.Y
		if x < len(grid) {
			checkGrid(x, y)
		}
		// SW
		x = key.X + 1
		y = key.Y - 1
		if x < len(grid) && y >= 0 {
			checkGrid(x, y)
		}
		// W
		x = key.X
		y = key.Y - 1
		if y >= 0 {
			checkGrid(x, y)
		}
	}
	return m
}

func getNumbersAdjacentToGears(symbols map[Key]string, grid [][]string) []Gear {
	gears := []Gear{}
	visited := make(map[Key]bool)
	var checkGrid = func(x, y int, numbers []int) []int {
		if visited[Key{X: x, Y: y}] {
			return numbers
		}
		if !util.IsStringDigit(grid[x][y]) {
			visited[Key{X: x, Y: y}] = true
			return numbers
		}
		// Left
		tempY := y
		for tempY > 0 && util.IsStringDigit(grid[x][tempY-1]) {
			tempY -= 1
			key := Key{X: x, Y: tempY}
			visited[key] = true
		}
		left := tempY
		// Right
		tempY = y
		for tempY < len(grid[x])-1 && util.IsStringDigit(grid[x][tempY+1]) {
			tempY += 1
			key := Key{X: x, Y: tempY}
			visited[key] = true
		}
		right := tempY
		visited[Key{X: x, Y: y}] = true
		stringNumber := grid[x][left : right+1]
		digits := strings.Join(stringNumber, "")
		number := util.ParseInt(digits)
		return append(numbers, number)
	}
	for key, val := range symbols {
		numbers := []int{}
		if val != "*" {
			continue
		}
		// NW
		x := key.X - 1
		y := key.Y - 1
		if x >= 0 && y >= 0 {
			numbers = checkGrid(x, y, numbers)
		}
		// N
		x = key.X - 1
		y = key.Y
		if x >= 0 {
			numbers = checkGrid(x, y, numbers)
		}
		// NE
		x = key.X - 1
		y = key.Y + 1
		if x >= 0 && y < len(grid[x]) {
			numbers = checkGrid(x, y, numbers)
		}
		// E
		x = key.X
		y = key.Y + 1
		if y < len(grid[x]) {
			numbers = checkGrid(x, y, numbers)
		}
		// SE
		x = key.X + 1
		y = key.Y + 1
		if x < len(grid) && y < len(grid[x]) {
			numbers = checkGrid(x, y, numbers)
		}
		// S
		x = key.X + 1
		y = key.Y
		if x < len(grid) {
			numbers = checkGrid(x, y, numbers)
		}
		// SW
		x = key.X + 1
		y = key.Y - 1
		if x < len(grid) && y >= 0 {
			numbers = checkGrid(x, y, numbers)
		}
		// W
		x = key.X
		y = key.Y - 1
		if y >= 0 {
			numbers = checkGrid(x, y, numbers)
		}
		gear := Gear{X: key.X, Y: key.Y, Numbers: numbers}
		gears = append(gears, gear)
	}
	return gears
}

func part1(path string) int {
	s := 0
	grid := util.ParseGridFromFile(path)
	symbols := findSymbolsInGrid(grid)
	numbers := getNumbersAdjacentToSymbols(symbols, grid)
	visited := make(map[Key]bool)
	q := []Key{}
	for number := range numbers {
		q = append(q, number)
	}
	for _, n := range q {
		if _, ok := visited[n]; ok {
			continue
		}
		// Left
		y := n.Y
		for y > 0 && util.IsStringDigit(grid[n.X][y-1]) {
			y -= 1
			key := Key{X: n.X, Y: y}
			visited[key] = true
		}
		left := y
		// Right
		y = n.Y
		for y < len(grid[n.X])-1 && util.IsStringDigit(grid[n.X][y+1]) {
			y += 1
			key := Key{X: n.X, Y: y}
			visited[key] = true
		}
		right := y
		visited[n] = true
		stringNumber := grid[n.X][left : right+1]
		digits := strings.Join(stringNumber, "")
		s += util.ParseInt(digits)
	}
	return s
}

func part2(path string) int {
	s := 0
	grid := util.ParseGridFromFile(path)
	symbols := findSymbolsInGrid(grid)
	gears := getNumbersAdjacentToGears(symbols, grid)
	for _, gear := range gears {
		if len(gear.Numbers) == 2 {
			s += gear.Numbers[0] * gear.Numbers[1]
		}
	}
	return s
}

func main() {
	file := "input.txt"
	// Part 1: 544664
	fmt.Println(part1(file))
	// Part 2: 84495585
	fmt.Println(part2(file))
}
