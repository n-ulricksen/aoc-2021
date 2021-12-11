package day10

import (
	"log"
	"sort"
	"strings"
)

var syntaxErrorScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var syntaxAutocompleteScore = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var matchingSymbol = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func part1(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var errorSum int
	for _, line := range lines {
		var stack []rune
	CheckSymbol:
		for _, symbol := range line {
			switch symbol {
			case '(', '[', '{', '<':
				stack = append(stack, symbol)
			case ')', ']', '}', '>':
				expected := matchingSymbol[stack[len(stack)-1]]
				if symbol != expected {
					errorSum += syntaxErrorScore[symbol]
					break CheckSymbol
				} else {
					stack = stack[:len(stack)-1]
				}
			default:
				log.Fatal("unsopported character:", symbol)
			}
		}
	}

	return errorSum
}

func part2(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var completionScores []int
	for _, line := range lines {
		var stack []rune
		var corrupted bool
	CheckSymbol:
		for _, symbol := range line {
			switch symbol {
			case '(', '[', '{', '<':
				stack = append(stack, symbol)
			case ')', ']', '}', '>':
				expected := matchingSymbol[stack[len(stack)-1]]
				if symbol != expected {
					// Corrupted line
					corrupted = true
					break CheckSymbol
				} else {
					stack = stack[:len(stack)-1]
				}
			default:
				log.Fatal("unsopported character:", symbol)
			}
		}

		if !corrupted && len(stack) > 0 {
			score := 0
			// Autocomplete line
			for len(stack) > 0 {
				symbol := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				match := matchingSymbol[symbol]
				score = score*5 + syntaxAutocompleteScore[match]
			}
			completionScores = append(completionScores, score)
		}
	}

	sort.Ints(completionScores)
	midIdx := len(completionScores) / 2

	return completionScores[midIdx]
}
