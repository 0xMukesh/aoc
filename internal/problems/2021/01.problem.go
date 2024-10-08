package problems_2021

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

var filename = "data/2021/01.txt"

type Problem_2021_01 struct {
}

func (p Problem_2021_01) Input() string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2021_01) ParseInput(input string) []int {
	lines := strings.Split(input, "\n")
	parsed := make([]int, len(lines))

	for i := range lines {
		num, _ := strconv.Atoi(lines[i])
		parsed[i] = num
	}

	return parsed
}

func (p Problem_2021_01) Solve_01() error {
	input := p.ParseInput(p.Input())
	result := 0

	for i := range input {
		if i != 0 {
			if input[i] > input[i-1] {
				result++
			}
		}
	}

	fmt.Println(result)

	return nil
}

func (p Problem_2021_01) Solve_02() error {
	input := p.ParseInput(p.Input())
	result := 0

	for i := range input {
		if i >= 3 {
			firstWindowSum := input[i-1] + input[i-2] + input[i-3]
			secondWindowSum := input[i] + input[i-1] + input[i-2]

			if secondWindowSum > firstWindowSum {
				result++
			}
		}
	}

	fmt.Println(result)

	return nil
}
