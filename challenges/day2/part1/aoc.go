package part1

import (
	_ "embed"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

// //go:embed sample_input
// var input string

func Solve() string {
	reports := parse()
	var safe int
	for _, report := range reports {
		if isSafe(report) {
			safe++
		}
	}
	return strconv.Itoa(safe)
}

func isSafe(numbers []string) bool {
	a, _ := strconv.Atoi(numbers[0])
	b, _ := strconv.Atoi(numbers[len(numbers)-1])
	increasing := isIncreasing(a, b)
	isSafe := true
	for i := 1; i < len(numbers); i++ {
		a, _ := strconv.Atoi(numbers[i-1])
		b, _ := strconv.Atoi(numbers[i])
		sum := int(math.Abs(float64(a) - float64(b)))
		if increasing != isIncreasing(a, b) || sum == 0 || sum > 3 {
			isSafe = false
			break
		}
	}
	return isSafe
}

func isIncreasing(a, b int) bool {
	return a < b
}

func parse() [][]string {
	input = strings.TrimSpace(input)
	rows := strings.Split(input, "\n")
	regex := regexp.MustCompile(`\s+`)

	arr := make([][]string, len(rows))

	for i, row := range rows {
		row := regex.Split(row, -1)
		arr[i] = row
	}
	return arr
}
