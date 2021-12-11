package day11

import (
	"strings"
)

type energyGrid [10][10]int

func part1(input string) int {
	grid := parseEnergyGrid(input)
	const numSteps = 100

	var flashCount int
	for i := 0; i < numSteps; i++ {
		flashCount += grid.step()
	}

	return flashCount
}

func part2(input string) int {
	grid := parseEnergyGrid(input)

	var flashCount int
	synchronousStep := -1
	for i := 0; i < 100000; i++ {
		flashesThisStep := grid.step()
		if flashesThisStep == 10*10 {
			synchronousStep = i
			break
		}
		flashCount += flashesThisStep
	}

	// 1-indexed answer
	return synchronousStep + 1
}

// step returns the number of flashes that occurred during this step
func (grid *energyGrid) step() int {
	for r, row := range grid {
		for c := range row {
			grid[r][c]++
		}
	}

	flashed := make([]bool, 10*10)
	flashCount := 0
	for r, row := range grid {
		for c := range row {
			if grid[r][c] > 9 && !flashed[r*len(grid[r])+c] {
				grid.flash(r, c, flashed, &flashCount)
			}
		}
	}

	for idx, didFlash := range flashed {
		if didFlash {
			r := idx / len(grid)
			c := idx % len(grid)
			grid[r][c] = 0
		}
	}

	return flashCount
}

func (grid *energyGrid) flash(r, c int, flashed []bool, flashCount *int) {
	cellIdx := r*len(grid[r]) + c
	if flashed[cellIdx] {
		return
	}

	flashed[cellIdx] = true
	(*flashCount)++

	grid.flashIncrease(r-1, c-1, flashed, flashCount)
	grid.flashIncrease(r-1, c, flashed, flashCount)
	grid.flashIncrease(r-1, c+1, flashed, flashCount)
	grid.flashIncrease(r, c+1, flashed, flashCount)
	grid.flashIncrease(r+1, c+1, flashed, flashCount)
	grid.flashIncrease(r+1, c, flashed, flashCount)
	grid.flashIncrease(r+1, c-1, flashed, flashCount)
	grid.flashIncrease(r, c-1, flashed, flashCount)
}

// flashIncrease is used to increase a neighboring cell while a cell is flashing.
func (grid *energyGrid) flashIncrease(r, c int, flashed []bool, flashCount *int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
		return
	}

	grid[r][c]++
	if grid[r][c] > 9 {
		grid.flash(r, c, flashed, flashCount)
	}
}

func parseEnergyGrid(input string) *energyGrid {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var newGrid energyGrid
	for r, line := range lines {
		for c, num := range line {
			level := int(num - '0')
			newGrid[r][c] = level
		}
	}

	return &newGrid
}
