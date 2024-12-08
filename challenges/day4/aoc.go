package day4

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

// //go:embed sample_input
// var input string

var directions = [][2]int{
	{-1, 0},  // North
	{1, 0},   // South
	{0, 1},   // East
	{0, -1},  // West
	{-1, 1},  // Northwest
	{-1, -1}, // Northeast
	{1, 1},   // Southeast
	{1, -1},  // Southwest
}

func SolvePartOne() string {
	grid := parse()
	word := "XMAS"
	regex := regexp.MustCompile(word)
	var res []string
	for r, row := range grid {
		for c := range row {
			coordinates := []int{r, c}
			boundries := []int{len(grid) - 1, len(row) - 1}
			findWordOccurances(&res, grid, coordinates, boundries, regex, word)
		}
	}
	return strconv.Itoa(len(res))
}

func findWordOccurances(res *[]string, grid [][]string, cor []int, boundry []int, regex *regexp.Regexp, word string) {
	boundry_x := boundry[0]
	boundry_y := boundry[1]
	for _, dir := range directions {
		var build string
		x := cor[0]
		y := cor[1]
		for 0 <= x && x <= boundry_x && 0 <= y && y <= boundry_y {
			build += grid[x][y]
			x = x + dir[0]
			y = y + dir[1]
			if len(build) == len(word) {
				break
			}
		}

		if regex.MatchString(build) && len(build) == len(word) {
			*res = append(*res, build)
		}
	}
}

func SolvePartTwo() string {
	grid := parse()
	var sum int
	for r, row := range grid {
		for c, col := range row {
			if col == "A" {
				coordinates := []int{r, c}
				boundries := []int{len(grid) - 1, len(row) - 1}
				corners := crossCoordinates(coordinates)

				continue_to_next_anchor := false
				for _, v := range corners {
					if v[0] < 0 || v[0] > boundries[0] || v[1] < 0 || v[1] > boundries[1] {
						continue_to_next_anchor = true
						break
					}
				}
				if continue_to_next_anchor {
					continue
				}
				upper_left, lower_right := corners["upper_left"], corners["lower_right"]
				upper_right, lower_left := corners["upper_right"], corners["lower_left"]

				isValidPair := func(a, b string) bool {
					return (a == "M" && b == "S") || (a == "S" && b == "M")
				}

				if !isValidPair(grid[upper_left[0]][upper_left[1]], grid[lower_right[0]][lower_right[1]]) ||
					!isValidPair(grid[upper_right[0]][upper_right[1]], grid[lower_left[0]][lower_left[1]]) {
					continue
				}

				sum++
			}
		}
	}
	return strconv.Itoa(sum)
}

func crossCoordinates(coordinate []int) map[string][2]int {
	var res [][2]int
	for i := 4; i < len(directions); i++ {
		x := coordinate[0] + directions[i][0]
		y := coordinate[1] + directions[i][1]
		res = append(res, [2]int{x, y})

	}
	corners := make(map[string][2]int)
	for i, v := range []string{"upper_right", "upper_left", "lower_right", "lower_left"} {
		corners[v] = res[i]
	}
	return corners
}

func parse() [][]string {
	input = strings.TrimSpace(input)
	rows := strings.Split(input, "\n")
	var parsed [][]string
	for _, row := range rows {
		parsed = append(parsed, strings.Split(row, ""))
	}
	return parsed
}
