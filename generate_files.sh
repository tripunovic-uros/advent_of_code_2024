#!/bin/bash

# Check argument
if [ "$#" -ne 1 ]; then
	echo "enter in a day as an arugment"
	exit 1
fi

# Check regex for 1..25
if ! [[ $1 =~ ^([1-9]|1[0-9]|2[0-5])$ ]]; then
	echo "day range is between 1-25"
	exit 1
fi

# Check if im getting a head of myself
day_number=$1
today_number=$(date -u | awk '{print $3}')
if [[ "$day_number" -gt "$today_number" ]]; then 
	echo "the challenge has not been released yet"
	exit 1
fi

path="day$day_number"
mkdir -p ./challenges/"$path"/

# Source env variables
. "./.env"
curl -Ls https://adventofcode.com/2024/day/$day_number/input -b "session=$AOC_SESSION_TOKEN" -o ./challenges/"$path"/input

# create a sample input file
echo "ADD sample" >  ./challenges/"$path"/sample_input

# Create solution file
cat > ./challenges/"$path"/aoc.go << EOL
package $path

import (
	_ "embed"
	"regexp"
	"strings"
)

// //go:embed input
// var input string

//go:embed sample_input
var input string

func SolvePartOne() string {
	return "TODO"
}

func SolvePartTwo() string {
	return "TODO"
}

func parse() {
	input = strings.TrimSpace(input)
	rows := strings.Split(input, "\n")
	regex := regexp.MustCompile(\`\s+\`)

	for i, row := range rows {
		row := regex.Split(row, -1)
	}
}
EOL

# Create test file
cat > ./challenges/"$path"/aoc_test.go << EOL
package $path

import (
	_ "embed"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	expected_part_one := "Replace here"
	result := SolvePartOne()
	if result != expected_part_one {
		t.Errorf("Expected %s, got %s", expected_part_one, result)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected_part_two := "Replace here"
	result := SolvePartTwo()
	if result != expected_part_two {
		t.Errorf("Expected %s, got %s", expected_part_two, result)
	}
}
EOL
