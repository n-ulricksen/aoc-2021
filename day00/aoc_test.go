package day00

import (
	"fmt"
	"log"
	"os"
	"testing"
)

const inputPath = "./input.txt"

const sInput = `
`

const sExpected = 0
const sExpected2 = 0

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
