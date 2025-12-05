package problems_2025

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2025_01 struct{}

func (p Problem_2025_01) Input() string {
	filename := "data/2025/01.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2025_01) ParseInput(input string) [][]string {
	splits := strings.Split(input, "\n")
	parsed := make([][]string, len(splits))

	for i, v := range splits {
		parsed[i] = []string{string(v[0]), v[1:]}
	}

	return parsed
}

func (p Problem_2025_01) Solve_01() error {
	input := p.ParseInput(p.Input())
	current := 50
	counter := 0

	for _, v := range input {
		dirStr := v[0]
		num, _ := strconv.Atoi(v[1])

		var dir utils.CircularDirection
		switch dirStr {
		case "L":
			dir = utils.CircularLeft
		case "R":
			dir = utils.CircularRight
		}

		current = utils.CircularRotation(current, num, dir, 100)

		if current == 0 {
			counter++
		}
	}

	fmt.Println(counter)
	return nil
}

func (p Problem_2025_01) Solve_02() error {
	input := p.ParseInput(p.Input())
	current := 50
	counter := 0

	for _, v := range input {
		dirStr := v[0]
		num, _ := strconv.Atoi(v[1])

		var dir utils.CircularDirection
		switch dirStr {
		case "L":
			dir = utils.CircularLeft
		case "R":
			dir = utils.CircularRight
		}

		for range num {
			current = utils.CircularRotation(current, 1, dir, 100)

			if current == 0 {
				counter++
			}
		}
	}

	fmt.Println(counter)
	return nil
}
