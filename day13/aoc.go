package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type grid [][]bool

type coordinate struct {
	x int
	y int
}

type fold struct {
	axis  byte // 'x' or 'y'
	value int
}

func part1(input string) int {
	// Parse input
	inputSplit := strings.Split(input, "\n\n")
	coordsText := inputSplit[0]
	coords := parseCoords(coordsText)
	foldsText := inputSplit[1]
	folds := parseFolds(foldsText)

	// Create paper of appropriate size
	maxX, maxY := getMaxXY(coords)
	paper := make(grid, maxY+1)
	for i := range paper {
		paper[i] = make([]bool, maxX+1)
	}

	// Plot points
	for _, coord := range coords {
		paper.plot(coord)
	}

	// Make folds
	paper.doFold(folds[0])

	// Count visible dots
	return paper.countVisibleDots()
}

func part2(input string) int {
	// Parse input
	inputSplit := strings.Split(input, "\n\n")
	coordsText := inputSplit[0]
	coords := parseCoords(coordsText)
	foldsText := inputSplit[1]
	folds := parseFolds(foldsText)

	// Create paper of appropriate size
	maxX, maxY := getMaxXY(coords)
	paper := make(grid, maxY+1)
	for i := range paper {
		paper[i] = make([]bool, maxX+1)
	}

	// Plot points
	for _, coord := range coords {
		paper.plot(coord)
	}

	// Make folds
	paper.print()
	for _, f := range folds {
		paper.doFold(f)
		paper.print()
	}

	// Count visible dots
	return paper.countVisibleDots()
}

func parseCoords(input string) []coordinate {
	coordsText := strings.Split(input, "\n")
	var coords []coordinate

	for _, coord := range coordsText {
		xy := strings.Split(coord, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		coords = append(coords, coordinate{x, y})
	}

	return coords
}

func parseFolds(input string) []fold {
	input = strings.TrimSpace(input)
	var folds []fold

	for _, line := range strings.Split(input, "\n") {
		lineSplit := strings.Split(line, " ")
		foldText := lineSplit[len(lineSplit)-1]
		foldToks := strings.Split(foldText, "=")
		foldAxis := foldToks[0][0]
		foldValue, _ := strconv.Atoi(foldToks[1])
		folds = append(folds, fold{foldAxis, foldValue})
	}

	return folds
}

// getMaxXY returns the maximum x and y values from the given coordinates.
func getMaxXY(coords []coordinate) (int, int) {
	var maxX, maxY int

	for _, coord := range coords {
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
	}

	return maxX, maxY
}

func (g *grid) plot(coord coordinate) {
	(*g)[coord.y][coord.x] = true
}

func (g *grid) doFold(f fold) {
	if f.axis == 'y' {
		for i := 0; i <= f.value; i++ {
			for x := 0; x < len((*g)[i]); x++ {
				low := (*g)[f.value-i][x]
				hi := (*g)[f.value+i][x]
				(*g)[f.value-i][x] = low || hi
			}
		}
		(*g) = (*g)[:f.value]
	} else if f.axis == 'x' {
		for i := 0; i <= f.value; i++ {
			for y := 0; y < len(*g); y++ {
				low := (*g)[y][f.value-i]
				hi := (*g)[y][f.value+i]
				(*g)[y][f.value-i] = low || hi
			}
		}
		for i := 0; i < len(*g); i++ {
			(*g)[i] = (*g)[i][:f.value]
		}
	}
}

func (g *grid) print() {
	for r := 0; r < len(*g); r++ {
		for c := 0; c < len((*g)[r]); c++ {
			if (*g)[r][c] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *grid) countVisibleDots() int {
	var visibleDots int

	for r := 0; r < len(*g); r++ {
		for c := 0; c < len((*g)[r]); c++ {
			if (*g)[r][c] {
				visibleDots++
			}
		}
	}

	return visibleDots
}
