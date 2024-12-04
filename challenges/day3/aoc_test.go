package day3

import (
	_ "embed"
	"testing"
)

expected_part_one := "Replace here" 
expected_part_two := "Replace here" 

func TestSolvePartOne(t *testing.T) {
	result := Solve()
	if result != expected_part_one {
		t.Errorf("Expected %s, got %s", expected_part_one, result)
	}
}

func TestSolvePartTwo(t *testing.T) {
	result := Solve()
	if result != expected_part_two {
		t.Errorf("Expected %s, got %s", expected_part_two, result)
	}
}
