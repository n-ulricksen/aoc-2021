package day12

import (
	"strings"
)

const (
	start = "start"
	end   = "end"
)

type cave map[string][]string

func part1(input string) int {
	c := parseCave(input)

	visited := make(map[string]int)
	paths := c.findPathsToEnd(start, visited)

	return paths
}

func part2(input string) int {
	c := parseCave(input)

	visited := make(map[string]int)
	paths := c.findPathsToEnd2(start, visited, false)

	return paths
}

func (c cave) findPathsToEnd(room string, visited map[string]int) int {
	if room == end {
		return 1
	}
	if visited[room] >= 1 {
		return 0
	}

	// Only lowercase rooms have a visit limit
	if isLowerCase(room[0]) {
		visited[room]++
	}

	var pathsToEnd int
	for _, adj := range c[room] {
		pathsToEnd += c.findPathsToEnd(adj, visited)
	}

	if isLowerCase(room[0]) {
		visited[room]--
	}

	return pathsToEnd
}

// same as above solution, but exactly one small room may be visited twice
func (c cave) findPathsToEnd2(room string, visited map[string]int, visitedSmallRoomTwice bool) int {
	if room == end {
		return 1
	}
	if visited[room] >= 1 {
		if visitedSmallRoomTwice || room == start {
			return 0
		}
		visitedSmallRoomTwice = true
	}

	// Only lowercase rooms have a visit limit
	if isLowerCase(room[0]) {
		visited[room]++
	}

	var pathsToEnd int
	for _, adj := range c[room] {
		pathsToEnd += c.findPathsToEnd2(adj, visited, visitedSmallRoomTwice)
	}

	if isLowerCase(room[0]) {
		visited[room]--
	}

	return pathsToEnd
}

func parseCave(input string) cave {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	newCave := make(map[string][]string)
	for _, line := range lines {
		nodes := strings.Split(line, "-")
		n1, n2 := nodes[0], nodes[1]
		newCave[n1] = append(newCave[n1], n2)
		newCave[n2] = append(newCave[n2], n1)
	}

	return newCave
}

func isLowerCase(b byte) bool {
	return b >= 97 && b <= 122
}
