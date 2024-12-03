package part2

import (
	_ "embed"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed input
var input string

func Solve() string {
	left, right := parse()
	sortArr(&left)
	sortArr(&right)

	var sum int
	for i := 0; i < len(left); i++ {
		value := left[i]
		similarity := 0
		for j := 0; j < len(right); j++ {
			if value == right[j] {
				similarity++
			}
		}
		sum += value * similarity
		similarity = 0
	}
	return strconv.Itoa(sum)
}

func parse() ([]int, []int) {
	input = strings.TrimSpace(input)
	rows := strings.Split(input, "\n")

	regex := regexp.MustCompile(`\s+`)

	left := make([]int, len(rows))
	right := make([]int, len(rows))

	for i, row := range rows {
		row := regex.Split(row, -1)
		l, _ := strconv.Atoi(row[0])
		r, _ := strconv.Atoi(row[1])
		left[i] = l
		right[i] = r
	}
	return left, right
}

func sortArr(s *[]int) {
	sort.Slice(*s, func(i, j int) bool {
		return (*s)[i] < (*s)[j]
	})
}
