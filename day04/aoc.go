package day04

import (
	"regexp"
	"strconv"
	"strings"
)

const reExp = `(\d+)`

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(reExp)
}

type bingoBoard struct {
	nums   [25]int
	marked [25]bool
}

func (bb *bingoBoard) markCell(numCalled int) {
	for i, n := range bb.nums {
		if n == numCalled {
			bb.marked[i] = true
			break
		}
	}
}

func (bb *bingoBoard) isWinner() bool {
	// rows
	for i := 0; i < 25; i += 5 {
		for j := 0; j < 5; j++ {
			// If any cell is unmarked in this row, break and check the next row.
			if !bb.marked[i+j] {
				break
			}
			if j == 4 {
				// Reaching this point means every cell in the row was marked. Bingo!
				return true
			}
		}
	}

	// columns
	for i := 0; i < 5; i++ {
		for j := 0; j < 25; j += 5 {
			if !bb.marked[i+j] {
				break
			}
			if j == 20 {
				return true
			}
		}
	}

	return false
}

func (bb *bingoBoard) getScore(lastCalled int) int {
	unmarkedSum := 0
	for i, num := range bb.nums {
		if !bb.marked[i] {
			unmarkedSum += num
		}
	}

	return lastCalled * unmarkedSum
}

func parseBingoBoards(input string) []*bingoBoard {
	boards := strings.Split(input, "\n\n")[1:]

	var bingoBoards []*bingoBoard
	for _, board := range boards {
		newBoard := &bingoBoard{
			nums:   [25]int{},
			marked: [25]bool{},
		}
		nums := re.FindAllString(board, -1)
		for i, num := range nums {
			n, _ := strconv.Atoi(num)
			newBoard.nums[i] = n
		}
		bingoBoards = append(bingoBoards, newBoard)
	}

	return bingoBoards
}

func parseDrawnNumbers(input string) []int {
	line0 := strings.Split(input, "\n")[0]
	drawnNumbers := strings.Split(line0, ",")
	numbers := []int{}
	for _, num := range drawnNumbers {
		n, _ := strconv.Atoi(num)
		numbers = append(numbers, n)
	}

	return numbers
}

func part1(input string) int {
	drawnNumbers := parseDrawnNumbers(input)

	bingoBoards := parseBingoBoards(input)

	// Play bingo
	var winner *bingoBoard
	var lastCalled int
BingoLoop:
	for _, num := range drawnNumbers {
		for _, board := range bingoBoards {
			board.markCell(num)
			if board.isWinner() {
				winner = board
				lastCalled = num
				break BingoLoop
			}
		}
	}

	return winner.getScore(lastCalled)
}

func part2(input string) int {
	drawnNumbers := parseDrawnNumbers(input)

	bingoBoards := parseBingoBoards(input)
	alreadyWon := make([]bool, len(bingoBoards))

	// Play bingo
	var lastWinner *bingoBoard
	var lastCalled int
	winners := 0
BingoLoop:
	for _, num := range drawnNumbers {
		for boardIdx, board := range bingoBoards {
			board.markCell(num)
			if board.isWinner() {
				if !alreadyWon[boardIdx] {
					winners++
				}
				alreadyWon[boardIdx] = true
				if winners == len(bingoBoards) {
					lastWinner = board
					lastCalled = num
					break BingoLoop
				}
			}
		}
	}

	return lastWinner.getScore(lastCalled)
}
