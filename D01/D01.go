package D01

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func Solution() {
	fmt.Println("\n--- Day 1: Historian Hysteria ---")

	// check file
	file, err := os.ReadFile("D01/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// parse each line
	rex, err := regexp.Compile(`(\d{5})   (\d{5})`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := rex.FindAllSubmatch(file, -1)
	if len(res) != 1000 {
		fmt.Println("<!> invalid input.txt")
		return
	}

	// convert to a slice of ints
	L1, L2 := make([]int, 1000), make([]int, 1000)
	for i, v := range res {
		n, err := strconv.Atoi(string(v[1]))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		L1[i] = n

		n, err = strconv.Atoi(string(v[2]))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		L2[i] = n
	}

	// lowest to highest
	slices.Sort(L1)
	slices.Sort(L2)

	total_distance := 0
	for i := range 1000 {
		if L1[i] > L2[i] {
			total_distance += L1[i] - L2[i]
		} else {
			total_distance += L2[i] - L1[i]
		}

	}

	fmt.Println("P1:", total_distance)
}
