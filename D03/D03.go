package D03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solution() {
	fmt.Println("\n--- Day 3: Mull It Over ---")

	file, err := os.ReadFile("D03/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	part_one(file)
	part_two(file)
}

func part_two(file []byte) {
	rex, err := regexp.Compile(`don't\(\)(?:.|\n)*?do\(\)`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := rex.FindAllSubmatch(file, -1)
	if len(res) == 0 {
		fmt.Println("<!> invalid input.txt")
		return
	}

	text := string(file)
	for _, do_not := range res {
		text = strings.ReplaceAll(text, string(do_not[0]), "do()")
	}

	P2, err := mulsum([]byte(text))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("P2:", P2)
}

func part_one(file []byte) {
	P1, err := mulsum(file)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("P1:", P1)
	}
}

func mulsum(file []byte) (int, error) {
	rex, err := regexp.Compile(`mul\((\d*),(\d*)\)`)
	if err != nil {
		return -1, err
	}
	res := rex.FindAllSubmatch(file, -1)
	if len(res) == 0 {
		fmt.Println("<!> invalid input.txt")
		return -1, err
	}

	mulsum := 0
	for _, v := range res {
		v1, err := strconv.Atoi(string(v[1]))
		if err != nil {
			fmt.Println(err.Error())
			return -1, err
		}
		v2, err := strconv.Atoi(string(v[2]))
		if err != nil {
			fmt.Println(err.Error())
			return -1, err
		}

		mulsum += v1 * v2
	}

	return mulsum, nil
}
