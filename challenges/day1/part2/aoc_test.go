package part2

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
