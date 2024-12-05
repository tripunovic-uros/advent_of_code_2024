package day3

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

// //go:embed sample_input
// var input string

func SolvePartOne() string {
	muls, _ := parse()
	sum := compute(muls)
	return strconv.Itoa(sum)
}

func SolvePartTwo() string {
	_, muls := parse()
	sum := compute(muls)
	return strconv.Itoa(sum)
}

func compute(muls [][]string) int {
	regex := regexp.MustCompile(`\d{1,3}`)
	var sum int
	for _, row := range muls {
		for _, mul := range row {
			parameters := regex.FindAllString(mul, -1)
			lhs := parameters[0]
			rhs := parameters[1]
			l, _ := strconv.Atoi(lhs)
			r, _ := strconv.Atoi(rhs)
			sum += l * r
		}
	}
	return sum
}

func parse() ([][]string, [][]string) {
	input = strings.TrimSpace(input)
	rows := strings.Split(input, "\n")

	var parsed_part_one [][]string
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	var parsed_part_two [][]string
	regex_two := regexp.MustCompile(`(do\(\)|don't\(\)|(mul\(\d{1,3},\d{1,3}\)))`)
	should_append := true

	for _, row := range rows {
		part_one := regex.FindAllString(row, -1)
		parsed_part_one = append(parsed_part_one, part_one)

		part_two := regex_two.FindAllString(row, -1)
		var new_part_two []string
		for _, instruction := range part_two {
			if instruction == "do()" {
				should_append = true
				continue
			}
			if instruction == "don't()" {
				should_append = false
				continue
			}
			if !should_append {
				continue
			}
			new_part_two = append(new_part_two, instruction)
		}
		parsed_part_two = append(parsed_part_two, new_part_two)
	}

	return parsed_part_one, parsed_part_two
}
