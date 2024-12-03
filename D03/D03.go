package D03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Solution() {
	fmt.Println("\n--- Day 3: Mull It Over ---")

	file, err := os.ReadFile("D03/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rex, err := regexp.Compile(`mul\((\d*),(\d*)\)`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := rex.FindAllSubmatch(file, -1)
	if len(res) == 0 {
		fmt.Println("<!> invalid input.txt")
		return
	}

	mulsum := 0
	for _, v := range res {
		v1, err := strconv.Atoi(string(v[1]))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		v2, err := strconv.Atoi(string(v[2]))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		mulsum += v1 * v2
	}

	fmt.Println("P1:", mulsum)
}
