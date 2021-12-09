package day05

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

var re *regexp.Regexp

const reExp = `(\d+)`

func init() {
	re = regexp.MustCompile(reExp)
}

type lineSegment struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func part1(input string) int {
	// Create a board using the max x/y values found
	lineSegments := parseLineSegments(input)

	maxX, maxY := 0, 0
	for _, segment := range lineSegments {
		maxX = max(maxX, segment.x1)
		maxX = max(maxX, segment.x2)
		maxY = max(maxY, segment.y1)
		maxY = max(maxY, segment.y2)
	}

	ventMap := make([][]int, maxY+1)
	for i := range ventMap {
		ventMap[i] = make([]int, maxX+1)
	}

	// Draw lines to the vent map
	for _, segment := range lineSegments {
		if segment.x1 == segment.x2 {
			x := segment.x1
			y1 := min(segment.y1, segment.y2)
			y2 := max(segment.y1, segment.y2)
			for y := y1; y <= y2; y++ {
				ventMap[y][x]++
			}
		} else if segment.y1 == segment.y2 {
			y := segment.y1
			x1 := min(segment.x1, segment.x2)
			x2 := max(segment.x1, segment.x2)
			for x := x1; x <= x2; x++ {
				ventMap[y][x]++
			}
		}
	}

	return findDangerousSpots(ventMap)
}

func part2(input string) int {
	// Create a board using the max x/y values found
	lineSegments := parseLineSegments(input)

	maxX, maxY := 0, 0
	for _, segment := range lineSegments {
		maxX = max(maxX, segment.x1)
		maxX = max(maxX, segment.x2)
		maxY = max(maxY, segment.y1)
		maxY = max(maxY, segment.y2)
	}

	ventMap := make([][]int, maxY+1)
	for i := range ventMap {
		ventMap[i] = make([]int, maxX+1)
	}

	// Draw lines to the vent map
	var dx, dy int
	var slope int
	for _, segment := range lineSegments {
		dx = segment.x1 - segment.x2
		dy = segment.y1 - segment.y2
		if dx == 0 {
			slope = int(math.Inf(1))
		} else {
			slope = dy / dx
		}
		if segment.x1 == segment.x2 {
			x := segment.x1
			y1 := min(segment.y1, segment.y2)
			y2 := max(segment.y1, segment.y2)
			for y := y1; y <= y2; y++ {
				ventMap[y][x]++
			}
		} else if segment.y1 == segment.y2 {
			y := segment.y1
			x1 := min(segment.x1, segment.x2)
			x2 := max(segment.x1, segment.x2)
			for x := x1; x <= x2; x++ {
				ventMap[y][x]++
			}
		} else if slope == 1 || slope == -1 {
			x := segment.x1
			y := segment.y1
			for x <= segment.x2 {
				ventMap[y][x]++
				x += 1
				y += slope
			}
		}
	}

	return findDangerousSpots(ventMap)
}

func parseLineSegments(input string) []lineSegment {
	lines := strings.Split(input, "\n")

	var lineSegments []lineSegment
	for _, line := range lines {
		nums := re.FindAllString(line, -1)
		if len(nums) < 4 {
			continue
		}
		x1, _ := strconv.Atoi(nums[0])
		y1, _ := strconv.Atoi(nums[1])
		x2, _ := strconv.Atoi(nums[2])
		y2, _ := strconv.Atoi(nums[3])
		if x1 > x2 {
			// Lower x-value first
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		lineSegments = append(lineSegments, lineSegment{
			x1, y1, x2, y2,
		})
	}

	return lineSegments
}

func findDangerousSpots(ventMap [][]int) int {
	dangerousSpots := 0
	for y := 0; y < len(ventMap); y++ {
		for x := 0; x < len(ventMap[y]); x++ {
			if ventMap[y][x] >= 2 {
				dangerousSpots++
			}
		}
	}

	return dangerousSpots
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
