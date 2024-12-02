package D02

import (
	"fmt"
	"os"
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
	reports := strings.Split(file_string, "\n")
	if len(reports) != line_count {
		fmt.Println("<!> invalid input.txt")
		return
	}

	safe := 0
	for _, report := range reports {
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

		if levels[0] == levels[1] {
			continue
		}

		all_increasing := false
		if levels[0] < levels[1] {
			all_increasing = true
		}

		unsafe := false
		for i, level := range levels[1:] {
			if level == levels[i] {
				unsafe = true
				break
			}

			if all_increasing {
				if levels[i]+4 > level && level > levels[i] {
					continue
				}
			} else {
				if levels[i]-4 < level && level < levels[i] {
					continue
				}
			}

			unsafe = true
			break
		}

		if unsafe {
			continue
		}
		safe += 1
	}

	fmt.Println("P1:", safe)
}
