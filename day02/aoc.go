package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input string) int {
	lines := strings.Split(input, "\n")
	horizPos, depth := 0, 0

	for _, line := range lines {
		toks := strings.Split(line, " ")
		if len(toks) < 2 {
			continue
		}
		command, m := toks[0], toks[1]
		mag, _ := strconv.Atoi(m)

		switch command {
		case "forward":
			horizPos += mag
		case "down":
			depth += mag
		case "up":
			depth -= mag
		default:
			fmt.Println("unrecognized command:", command)
			os.Exit(1)
		}
	}

	return horizPos * depth
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	horizPos, depth, aim := 0, 0, 0

	for _, line := range lines {
		toks := strings.Split(line, " ")
		if len(toks) < 2 {
			continue
		}
		command, m := toks[0], toks[1]
		mag, _ := strconv.Atoi(m)

		switch command {
		case "forward":
			horizPos += mag
			depth += aim * mag
		case "down":
			aim += mag
		case "up":
			aim -= mag
		default:
			fmt.Println("unrecognized command:", command)
			os.Exit(1)
		}
	}

	return horizPos * depth
}
