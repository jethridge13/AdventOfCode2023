package main

import (
	"fmt"
	"strings"

	"github.com/jethridge13/AdventOfCode2023/util"
)

type Card struct {
	CardNumber     int
	Numbers        map[int]bool
	WinningNumbers map[int]bool
	Qty            int
}

func parseLine(line string) Card {
	parts := strings.Split(line, ": ")
	nPart := strings.Split(parts[0], " ")
	n := util.ParseInt(nPart[len(nPart)-1])
	numbers := strings.Split(parts[1], " | ")
	cardNumbers := make(map[int]bool)
	winningNumbers := make(map[int]bool)
	for _, number := range strings.Split(numbers[0], " ") {
		if len(number) < 1 {
			continue
		}
		cardNumbers[util.ParseInt(number)] = true
	}
	for _, number := range strings.Split(numbers[1], " ") {
		if len(number) < 1 {
			continue
		}
		winningNumbers[util.ParseInt(number)] = true
	}
	return Card{CardNumber: n, Numbers: cardNumbers, WinningNumbers: winningNumbers, Qty: 1}
}

func getCardScore(card Card) int {
	s := 0
	for n := range card.Numbers {
		if card.WinningNumbers[n] {
			if s == 0 {
				s = 1
			} else {
				s *= 2
			}
		}
	}
	return s
}

func getWinningCount(card Card) int {
	s := 0
	for n := range card.Numbers {
		if card.WinningNumbers[n] {
			s += 1
		}
	}
	return s
}

func part1(path string) int {
	scanner := util.GetFileScanner(path)
	cards := []Card{}
	for scanner.Scan() {
		card := parseLine(scanner.Text())
		cards = append(cards, card)
	}
	s := 0
	for _, card := range cards {
		s += getCardScore(card)
	}
	return s
}

func part2(path string) int {
	s := 0
	scanner := util.GetFileScanner(path)
	cards := []Card{}
	for scanner.Scan() {
		card := parseLine(scanner.Text())
		cards = append(cards, card)
	}
	for index, card := range cards {
		s += card.Qty
		winners := getWinningCount(card)
		for i := 0; i < winners; i++ {
			cards[index+i+1].Qty += card.Qty
		}
	}
	return s
}

func main() {
	file := "input.txt"
	// Part 1: 23673
	fmt.Println(part1(file))
	// Part 2: 12263631
	fmt.Println(part2(file))
}
