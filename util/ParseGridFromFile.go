package util

import "strings"

func ParseGridFromFile(path string) [][]string {
	grid := [][]string{}
	scanner := GetFileScanner(path)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}
