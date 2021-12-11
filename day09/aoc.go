package day00

import (
	"sort"
	"strings"
)

func part1(input string) int {
	input = strings.TrimSpace(input)
	heightMap := parseHeightMap(input)

	var riskLevelSum int
	for r, row := range heightMap {
		for c, height := range row {
			if isLowPoint(heightMap, r, c) {
				riskLevel := height + 1
				riskLevelSum += riskLevel
			}
		}
	}

	return riskLevelSum
}

func part2(input string) int {
	input = strings.TrimSpace(input)
	heightMap := parseHeightMap(input)

	var basinSizes []int
	for r, row := range heightMap {
		for c, _ := range row {
			if isLowPoint(heightMap, r, c) {
				basinSize := calcBasinSize(heightMap, r, c)
				basinSizes = append(basinSizes, basinSize)
			}
		}
	}

	sort.Ints(basinSizes)

	return basinSizes[len(basinSizes)-1] *
		basinSizes[len(basinSizes)-2] *
		basinSizes[len(basinSizes)-3]
}

func parseHeightMap(input string) [][]int {
	lines := strings.Split(input, "\n")

	var heightMap [][]int
	for _, line := range lines {
		var row []int
		for _, ch := range line {
			n := int(ch - '0')
			row = append(row, n)
		}
		heightMap = append(heightMap, row)
	}

	return heightMap
}

func isLowPoint(heightMap [][]int, r int, c int) bool {
	point := heightMap[r][c]

	// check up, down, left, right
	if r-1 >= 0 && heightMap[r-1][c] <= point {
		return false
	}
	if r+1 < len(heightMap) && heightMap[r+1][c] <= point {
		return false
	}
	if c-1 >= 0 && heightMap[r][c-1] <= point {
		return false
	}
	if c+1 < len(heightMap[r]) && heightMap[r][c+1] <= point {
		return false
	}

	return true
}

func calcBasinSize(heightMap [][]int, r int, c int) int {
	// Starting at the given low point, search outwards until either:
	//   - the next value decreases
	//   - the next value is a 9
	visited := make([]bool, len(heightMap)*len(heightMap[0]))

	var dfs func(int, int, int) int
	dfs = func(r, c, prev int) int {
		if r < 0 || r >= len(heightMap) || c < 0 || c >= len(heightMap[r]) {
			return 0
		}

		height := heightMap[r][c]
		if height <= prev || height == 9 {
			return 0
		}

		visitIdx := r*len(heightMap[0]) + c
		if visited[visitIdx] {
			return 0
		}
		visited[visitIdx] = true

		return 1 +
			dfs(r-1, c, height) +
			dfs(r+1, c, height) +
			dfs(r, c-1, height) +
			dfs(r, c+1, height)
	}

	return dfs(r, c, -1)
}
