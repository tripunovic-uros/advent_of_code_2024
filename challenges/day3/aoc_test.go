package day3

import (
	_ "embed"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	expected_part_one := "187194524"
	result := SolvePartOne()
	if result != expected_part_one {
		t.Errorf("Expected %s, got %s", expected_part_one, result)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected_part_two := "127092535"
	result := SolvePartTwo()
	if result != expected_part_two {
		t.Errorf("Expected %s, got %s", expected_part_two, result)
	}
}
