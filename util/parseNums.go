package util

import (
	"strconv"
	"strings"
)

func ParseNums(input string) []int {
	input = strings.TrimSpace(input)
	split := strings.Split(input, ",")

	var nums []int
	for _, num := range split {
		n, _ := strconv.Atoi(num)
		nums = append(nums, n)
	}

	return nums
}
