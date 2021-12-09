package day06

import (
	"strconv"
	"strings"
)

const daysP1 int = 80
const daysP2 int = 256

func part1(input string) int {
	fishes := parseNums(input)

	for i := 0; i < daysP1; i++ {
		for fishIdx, fish := range fishes {
			if fish == 0 {
				fishes[fishIdx] = 7
				fishes = append(fishes, 8)
			}
			fishes[fishIdx]--
		}
	}

	return len(fishes)
}

func part2(input string) int {
	fishes := parseNums(input)

	// Part1 naive solution is too slow for part 2. Instead we should keep
	// count of # of fish at each stage of their life (0-8).
	fishCount := make([]int, 9)
	for _, fish := range fishes {
		fishCount[fish]++
	}

	for i := 0; i < daysP2; i++ {
		reproducing := fishCount[0]

		for countIdx := 1; countIdx < len(fishCount); countIdx++ {
			fishCount[countIdx-1] = fishCount[countIdx]
		}

		fishCount[8] = reproducing
		fishCount[6] += reproducing
	}

	totalFish := 0
	for _, count := range fishCount {
		totalFish += count
	}

	return totalFish
}

func parseNums(input string) []int {
	input = strings.TrimSpace(input)
	split := strings.Split(input, ",")

	var nums []int
	for _, num := range split {
		n, _ := strconv.Atoi(num)
		nums = append(nums, n)
	}

	return nums
}
