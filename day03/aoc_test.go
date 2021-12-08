package day03

import (
	"fmt"
	"log"
	"os"
	"testing"
)

const inputPath = "./input.txt"

const sInput = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

const sExpected = 198
const sExpected2 = 230

var input string

func init() {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	input = string(data)
}

func TestPart1(t *testing.T) {
	sGot := part1(sInput)
	if sGot != sExpected {
		t.Fatalf("got %d, want %d\n", sGot, sExpected)
	}

	result := part1(input)
	fmt.Println(result)
}

func TestPart2(t *testing.T) {
	sGot := part2(sInput)
	if sGot != sExpected2 {
		t.Fatalf("got %d, want %d\n", sGot, sExpected2)
	}

	result := part2(input)
	fmt.Println(result)
}
