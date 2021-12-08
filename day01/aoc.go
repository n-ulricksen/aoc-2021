package day01

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	nums := strToInts(input)
	increasing := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			increasing++
		}
	}

	return increasing
}

func part2(input string) int {
	nums := strToInts(input)
	increasing := 0

	prevSum := nums[0] + nums[1] + nums[2]
	for i := 3; i < len(nums); i++ {
		sum := prevSum - nums[i-3] + nums[i]
		if sum > prevSum {
			increasing++
		}
		prevSum = sum
	}

	return increasing
}

func strToInts(str string) []int {
	var ret []int

	lines := strings.Split(str, "\n")
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		ret = append(ret, num)
	}

	return ret
}
