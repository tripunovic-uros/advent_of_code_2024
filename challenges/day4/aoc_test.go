package day4

import (
	_ "embed"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	expected_part_one := "2336"
	result := SolvePartOne()
	if result != expected_part_one {
		t.Errorf("Expected %s, got %s", expected_part_one, result)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected_part_two := "1831"
	result := SolvePartTwo()
	if result != expected_part_two {
		t.Errorf("Expected %s, got %s", expected_part_two, result)
	}
}
