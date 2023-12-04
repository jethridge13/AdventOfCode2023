package main

import (
	"fmt"
	"strings"

	"github.com/jethridge13/AdventOfCode2023/util"
)

type Game struct {
	Id    int
	Red   int
	Green int
	Blue  int
}

func parseGame(line string) *Game {
	parts := strings.Split(line, ": ")
	id := util.ParseInt(strings.Split(parts[0], " ")[1])
	games := strings.Split(parts[1], "; ")
	red := 0
	green := 0
	blue := 0
	for _, game := range games {
		cubes := strings.Split(game, ", ")
		for _, cube := range cubes {
			cubeParts := strings.Split(cube, " ")
			s := util.ParseInt(cubeParts[0])
			switch cubeParts[1] {
			case "red":
				red = util.Max(red, s)
			case "green":
				green = util.Max(green, s)
			case "blue":
				blue = util.Max(blue, s)
			}
		}
	}
	return &Game{Id: id, Red: red, Green: green, Blue: blue}
}

func isValidGame(baseGame *Game, game *Game) bool {
	return game.Red <= baseGame.Red && game.Green <= baseGame.Green && game.Blue <= baseGame.Blue
}

func part1(path string) int {
	scanner := util.GetFileScanner(path)
	baseGame := &Game{Id: 0, Red: 12, Green: 13, Blue: 14}
	s := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)
		if isValidGame(baseGame, game) {
			s += game.Id
		}
	}
	return s
}

func part2(path string) int {
	scanner := util.GetFileScanner(path)
	s := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)
		power := game.Red * game.Green * game.Blue
		s += power
	}
	return s
}

func main() {
	file := "input.txt"
	// Part 1: 2278
	fmt.Println(part1(file))
	// Part 2: 67953
	fmt.Println(part2(file))
}
