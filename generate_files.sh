#!/bin/bash

# Check arguments
if [ "$#" -ne 1 ]; then
	echo "enter in a day as an arugment"
	exit 1
fi

# Check regex for 1..25

day=$1
path="day$day"
mkdir -p ./challenges/"$path"/part{1..2}
# Source env variables
. "./.env"
for part in {1..2}; do
	curl -Ls https://adventofcode.com/2024/day/$day/input -b "session=$AOC_SESSION_TOKEN" -o ./challenges/"$path"/part"$part"/input

	# Create solution file
    cat > ./challenges/"$path"/part"$part"/aoc.go << EOL
package part$part

import (
	_ "embed"
)

//go:embed input
var input string

func Solve() string {
	return "TODO"
}
EOL

    # Create test file
    cat > ./challenges/"$path"/part"$part"/aoc_test.go << EOL
package part$part

import (
	_ "embed"
	"testing"
)

//go:embed expected
var expected string

func TestSolve(t *testing.T) {
	result := Solve()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
EOL

    # Add placeholder expected file
    echo "ADD RESULT" > ./challenges/"$path"/part"$part"/expected
done
