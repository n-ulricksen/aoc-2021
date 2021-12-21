package day14

import (
	"math"
	"strings"
)

func part1(input string) int {
	start, rules := parseInput(input)

	insertSteps := 10
	result := start
	for i := 0; i < insertSteps; i++ {
		result = runInsertionStep(result, rules)
	}

	charCount := make(map[rune]int)
	for _, ch := range result {
		charCount[ch]++
	}

	leastCommon := int(math.MaxInt64)
	mostCommon := int(math.MinInt64)
	for _, count := range charCount {
		leastCommon = min(leastCommon, count)
		mostCommon = max(mostCommon, count)
	}

	return mostCommon - leastCommon
}

func part2(input string) int {
	start, rules := parseInput(input)

	// count initial pairs
	pairCount := make(map[string]int)
	for i := 0; i < len(start)-1; i++ {
		pair := start[i : i+2]
		pairCount[pair]++
	}

	insertSteps := 40
	for i := 0; i < insertSteps; i++ {
		newPairCount := make(map[string]int)
		for k, v := range pairCount {
			newPairCount[k] = v
		}
		for pair, count := range pairCount {
			// each pair creates 2 new pairs, and then "deletes" itself
			p1 := pair[:1] + rules[pair]
			p2 := rules[pair] + pair[1:]
			newPairCount[p1] += count
			newPairCount[p2] += count
			if count > 0 {
				newPairCount[pair] -= count
			}
		}

		pairCount = newPairCount
	}

	// count the individual letters
	letterCount := make(map[byte]int)
	for pair, count := range pairCount {
		l1 := pair[0]
		l2 := pair[1]
		letterCount[l1] += count
		letterCount[l2] += count
	}
	for letter := range letterCount {
		if letter == start[0] || letter == start[len(start)-1] {
			letterCount[letter] += 1
		}
		letterCount[letter] /= 2
	}

	leastCommon := int(math.MaxInt64)
	mostCommon := int(math.MinInt64)
	for _, count := range letterCount {
		leastCommon = min(leastCommon, count)
		mostCommon = max(mostCommon, count)
	}

	return mostCommon - leastCommon
}

// parseInput returns the starting string and the map of insertion rules.
func parseInput(input string) (string, map[string]string) {
	inputSplit := strings.Split(input, "\n\n")
	start := inputSplit[0]
	rulesText := strings.TrimSpace(inputSplit[1])

	rules := make(map[string]string)
	for _, line := range strings.Split(rulesText, "\n") {
		lineSplit := strings.Split(line, " -> ")
		rules[lineSplit[0]] = lineSplit[1]
	}

	return start, rules
}

func runInsertionStep(prevString string, rules map[string]string) string {
	nextString := ""

	for i := 0; i < len(prevString)-1; i++ {
		pair := prevString[i : i+2]
		nextString += pair[:1]
		nextString += rules[pair]
	}
	nextString += prevString[len(prevString)-1:]

	return nextString
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
