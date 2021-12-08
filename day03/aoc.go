package day03

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	lines := strings.Split(input, "\n")

	gamma := 0
	epsilon := 0

	zeroes := make([]int, len(lines[0]))
	ones := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, n := range line {
			bit := n - '0'
			if bit == 0 {
				zeroes[i]++
			} else {
				ones[i]++
			}
		}
	}

	for i := range zeroes {
		gamma <<= 1
		epsilon <<= 1
		zeroCount, oneCount := zeroes[i], ones[i]
		if oneCount > zeroCount {
			gamma += 1
		} else {
			epsilon += 1
		}
	}

	return gamma * epsilon
}

func part2(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	oxygen := 0
	co2 := 0

	oxyPotentials := make([]string, len(lines))
	co2Potentials := make([]string, len(lines))
	copy(oxyPotentials, lines)
	copy(co2Potentials, lines)

	zeroes, ones := []string{}, []string{}

	bitIdx := 0 // from left
	for len(oxyPotentials) > 1 {
		for _, line := range oxyPotentials {
			bit := line[bitIdx] - '0'
			if bit == 0 {
				zeroes = append(zeroes, line)
			} else {
				ones = append(ones, line)
			}
		}
		if len(zeroes) > len(ones) {
			oxyPotentials = zeroes
		} else {
			oxyPotentials = ones
		}
		zeroes, ones = []string{}, []string{}
		bitIdx++
	}

	bitIdx = 0
	for len(co2Potentials) > 1 {
		for _, line := range co2Potentials {
			bit := line[bitIdx] - '0'
			if bit == 0 {
				zeroes = append(zeroes, line)
			} else {
				ones = append(ones, line)
			}
		}
		if len(ones) < len(zeroes) {
			co2Potentials = ones
		} else {
			co2Potentials = zeroes
		}
		zeroes, ones = []string{}, []string{}
		bitIdx++
	}

	n, _ := strconv.ParseInt(oxyPotentials[0], 2, 64)
	oxygen = int(n)

	n, _ = strconv.ParseInt(co2Potentials[0], 2, 64)
	co2 = int(n)

	return oxygen * co2
}
