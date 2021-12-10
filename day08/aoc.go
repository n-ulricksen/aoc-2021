package day08

import (
	"log"
	"sort"
	"strings"
)

var lengthOne = map[int]bool{
	2: true,
	4: true,
	3: true,
	7: true,
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	var simpleCount int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		outputValues := strings.Split(line, "|")[1]
		values := strings.Split(outputValues, " ")
		for _, val := range values {
			if lengthOne[len(val)] {
				simpleCount++
			}
		}
	}

	return simpleCount
}

func part2(input string) int {
	var outputSum int
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		lineSplit := strings.Split(line, "|")
		signalPatternsStr := strings.TrimSpace(lineSplit[0])

		signalPatterns := strings.Split(signalPatternsStr, " ")
		sort.Slice(signalPatterns, func(i, j int) bool {
			return len(signalPatterns[i]) < len(signalPatterns[j])
		})

		// bitmap -> digit
		digits := map[byte]int{
			encode(signalPatterns[0]): 1, // only pattern with 2 segments set
			encode(signalPatterns[1]): 7, // only pattern with 3 segments set
			encode(signalPatterns[2]): 4, // only pattern with 4 segments set
			encode(signalPatterns[9]): 8, // only pattern with 2 segments set
		}
		// digit -> bitmap
		bitmaps := make([]byte, 10)
		bitmaps[1] = encode(signalPatterns[0]) // length 2
		bitmaps[7] = encode(signalPatterns[1]) // length 3
		bitmaps[4] = encode(signalPatterns[2]) // length 4
		bitmaps[8] = encode(signalPatterns[9]) // length 7

		for i := 3; i <= 5; i++ {
			e := encode(signalPatterns[i])
			if e&bitmaps[1] == bitmaps[1] {
				// find '3' from patterns of length 5 (AND with 1, should = 1)
				bitmaps[3] = e
				digits[e] = 3
			} else if e|bitmaps[4] == bitmaps[8] {
				// find '2' from patterns of length 5 (OR with 4, should = 8)
				bitmaps[2] = e
				digits[e] = 2
			} else {
				bitmaps[5] = e
				digits[e] = 5
			}
		}

		for i := 6; i <= 8; i++ {
			e := encode(signalPatterns[i])
			if e&bitmaps[4] == bitmaps[4] {
				// find '9' from patterns of length 6 (AND with 4, should = 4)
				bitmaps[9] = e
				digits[e] = 9
			} else if e&bitmaps[7] == bitmaps[7] {
				// find '0' from patterns of length 6 (AND with 7, should = 7)
				bitmaps[0] = e
				digits[e] = 0
			} else {
				bitmaps[6] = e
				digits[e] = 6
			}
		}

		outputPatternsStr := strings.TrimSpace(lineSplit[1])
		outputPatterns := strings.Split(outputPatternsStr, " ")

		var output int
		for _, pattern := range outputPatterns {
			digit := digits[encode(pattern)]
			output *= 10
			output += digit
		}
		outputSum += output
	}

	return outputSum
}

// Encode signal as bitmap
func encode(input string) byte {
	var encoded byte

	var bit byte
	for _, ch := range input {
		switch ch {
		case 'a':
			bit = 1 << 6
		case 'b':
			bit = 1 << 5
		case 'c':
			bit = 1 << 4
		case 'd':
			bit = 1 << 3
		case 'e':
			bit = 1 << 2
		case 'f':
			bit = 1 << 1
		case 'g':
			bit = 1 << 0
		default:
			log.Fatal("unsupported character", ch)
		}
		encoded |= bit
	}

	return encoded
}
