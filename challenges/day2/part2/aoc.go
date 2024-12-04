package part2

import (
	_ "embed"
	"math"
	"regexp"
	"slices"
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
		for i := range report {
			// sub_slice := append(append([]string{}, report[:i]...), report[i+1:]...)
			copy_report_slice := slices.Clone(report)
			sub_report_slice := slices.Delete(copy_report_slice, i, i+1)
			if isSafe(sub_report_slice) {
				safe++
				break
			}
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
