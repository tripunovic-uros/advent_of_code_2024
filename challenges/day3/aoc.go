package day3

import (
	_ "embed"
	"regexp"
	"strings"
)

//go:embed input
var input string

// //go:embed sample_input
// var input string

func SolvePartOne() string {
	return "TODO"
}

func SolvePartTwo() string {
	return "TODO"
}

func parse() {
	input = strings.TrimSpace(input)
	rows := strings.Split(input, "\n")
	regex := regexp.MustCompile(`\s+`)

	for i, row := range rows {
		row := regex.Split(row, -1)
	}
}
