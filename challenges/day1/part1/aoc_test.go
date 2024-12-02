package part1

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
