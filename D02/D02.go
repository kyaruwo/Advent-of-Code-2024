package D02

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const line_count int = 1000

func Solution() {
	fmt.Println("\n--- Day 2: Red-Nosed Reports ---")

	file, err := os.ReadFile("D02/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	file_string := string(file)
	reports_string := strings.Split(file_string, "\n")
	if len(reports_string) != line_count {
		fmt.Println("<!> invalid input.txt")
		return
	}

	reports := make([][]int, len(reports_string))
	for ri, report := range reports_string {
		levels_string := strings.Split(report, " ")
		levels := make([]int, len(levels_string))

		for i, level := range levels_string {
			level = strings.TrimRight(level, "\r")
			n, err := strconv.Atoi(level)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			levels[i] = n
		}
		reports[ri] = levels
	}

	part_one(reports)
	part_two(reports)
}

func part_two(reports [][]int) {
	safe := 0

	for _, levels := range reports {
		if safety_check(levels) {
			safe += 1
			continue
		}

		for i := range len(levels) {
			new_levels := make([]int, len(levels))
			copy(new_levels, levels)
			new_levels = slices.Delete(new_levels, i, i+1)

			if safety_check(new_levels) {
				safe += 1
				break
			}
		}
	}

	fmt.Println("P2:", safe)
}

func part_one(reports [][]int) {
	safe := 0

	for _, levels := range reports {
		if safety_check(levels) {
			safe += 1
		}
	}

	fmt.Println("P1:", safe)
}

func safety_check(levels []int) bool {
	all_increasing := false
	if levels[0] < levels[1] {
		all_increasing = true
	}

	if levels[0] == levels[1] {
		return false
	}

	for i, level := range levels[1:] {
		if all_increasing {
			if levels[i]+4 > level && level > levels[i] {
				continue
			}
		} else {
			if levels[i]-4 < level && level < levels[i] {
				continue
			}
		}

		return false
	}

	return true
}
