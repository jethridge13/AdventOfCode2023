package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/jethridge13/AdventOfCode2023/util"
)

func getIntFromRuneStack(stack []rune) int {
	if len(stack) == 0 {
		return 0
	}
	digits := []rune{stack[0], stack[len(stack)-1]}
	n := string(digits)
	val, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	return val
}

func getSimpleDigitFromLine(line string) int {
	stack := make([]rune, 0)
	for _, c := range line {
		if unicode.IsDigit(c) {
			stack = append(stack, c)
		}
	}
	return getIntFromRuneStack(stack)
}

func getDigitFromWord(line string) (rune, error) {
	switch line {
	case "one":
		return '1', nil
	case "two":
		return '2', nil
	case "three":
		return '3', nil
	case "four":
		return '4', nil
	case "five":
		return '5', nil
	case "six":
		return '6', nil
	case "seven":
		return '7', nil
	case "eight":
		return '8', nil
	case "nine":
		return '9', nil
	}
	return '0', fmt.Errorf("No number found")
}

func getComplexDigitFromLine(line string, trie *util.Trie) int {
	stack := make([]rune, 0)
	current := trie.RootNode
	for _, c := range line {
		if unicode.IsDigit(c) {
			stack = append(stack, c)
			current = trie.RootNode
		} else {
			current = current.Traverse(c)
			if current == nil {
				current = trie.RootNode
				candidate := current.Traverse(c)
				if candidate != nil {
					current = candidate
				}
			} else if current.IsWord {
				d, err := getDigitFromWord(current.Word)
				util.ErrCheck(err)
				stack = append(stack, d)
				current = trie.RootNode
				candidate := current.Traverse(c)
				if candidate != nil {
					current = candidate
				}
			}
		}
	}
	return getIntFromRuneStack(stack)
}

func part1(path string) int {
	scanner := util.GetFileScanner(path)
	s := 0
	for scanner.Scan() {
		line := scanner.Text()
		val := getSimpleDigitFromLine(line)
		s += val
	}
	return s
}

func part2(path string) int {
	scanner := util.GetFileScanner(path)
	s := 0
	trie := util.NewTrie()
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, n := range numbers {
		trie.Insert(n)
	}
	for scanner.Scan() {
		line := scanner.Text()
		val := getComplexDigitFromLine(line, trie)
		s += val
	}
	return s
}

func main() {
	file := "input.txt"
	// Part 1: 55172
	fmt.Println(part1(file))
	// Part 2: 54925
	fmt.Println(part2(file))
}
