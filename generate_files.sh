#!/bin/bash

# Check argument
if [ "$#" -ne 1 ]; then
	echo "enter in a day as an arugment"
	exit 1
fi

# Check regex for 1..25
if [[ $1 =~ '^([1-9]|1[0-9]|2[0-5])$' ]]; then
	echo "day range is between 1-25"
	exit 1
fi

day_number=$1
path="day$day_number"

mkdir -p ./challenges/"$path"/part{1..2}

# Check curl maybe?

# Source env variables
. "./.env"
for part_number in {1..2}; do
	curl -Ls https://adventofcode.com/2024/day/$day_number/input -b "session=$AOC_SESSION_TOKEN" -o ./challenges/"$path"/part"$part_number"/input

	# Create solution file
    cat > ./challenges/"$path"/part"$part_number"/aoc.go << EOL
package part$part_number

import (
	_ "embed"
	"strings"
)

//go:embed input
var input string

func Solve() string {
	return "TODO"
}

func parse() {
	input = strings.TrimSpace(input)
}
EOL

    # Create test file
    cat > ./challenges/"$path"/part"$part_number"/aoc_test.go << EOL
package part$part_number

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed expected
var expected string

func TestSolve(t *testing.T) {
	expected = strings.TrimSpace(expected)
	result := Solve()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
EOL

    # Add placeholder expected file
    echo "ADD RESULT" > ./challenges/"$path"/part"$part_number"/expected
done
