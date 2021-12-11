package day07

import (
	"math"

	"github.com/n-ulricksen/aoc-2021/util"
)

func part1(input string) int {
	positions := util.ParseNums(input)
	posCounts := make(map[int]int)
	maxPos := 0
	for _, pos := range positions {
		maxPos = max(maxPos, pos)
		posCounts[pos]++
	}

	var fuelCosts []int
	var fuelCost, distance int
	for i := 0; i <= maxPos; i++ {
		fuelCost = 0
		for pos, count := range posCounts {
			distance = abs(pos - i)
			fuelCost += distance * count
		}
		fuelCosts = append(fuelCosts, fuelCost)
	}

	minFuel := math.MaxInt64
	for _, cost := range fuelCosts {
		minFuel = min(minFuel, cost)
	}

	return minFuel
}

func part2(input string) int {
	positions := util.ParseNums(input)
	posCounts := make(map[int]int)
	maxPos := 0
	for _, pos := range positions {
		maxPos = max(maxPos, pos)
		posCounts[pos]++
	}

	var fuelCosts []int
	var fuelCost, distance int
	for i := 0; i <= maxPos; i++ {
		fuelCost = 0
		for pos, count := range posCounts {
			distance = abs(pos - i)
			fuelCost += sumFirstN(distance) * count
		}
		fuelCosts = append(fuelCosts, fuelCost)
	}

	minFuel := math.MaxInt64
	for _, cost := range fuelCosts {
		minFuel = min(minFuel, cost)
	}

	return minFuel
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func sumFirstN(n int) int {
	return (n * (n + 1)) / 2
}
