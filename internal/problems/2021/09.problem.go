package problems_2021

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2021_09 struct{}

func (p Problem_2021_09) Input() string {
	filename := "data/2021/09.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2021_09) ParseInput(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")
	parsed := make([][]int, len(lines))

	for i, line := range lines {
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}

			parsed[i] = append(parsed[i], num)
		}
	}

	return parsed, nil
}

func (p Problem_2021_09) Solve_01() error {
	input, err := p.ParseInput(p.Input())
	if err != nil {
		return err
	}

	numOfRows := len(input)
	numOfCols := len(input[0])

	var lowPoints []int

	for i := range input {
		for j := range input[i] {
			num := input[i][j]
			res := 1

			var (
				top    *int
				bottom *int
				left   *int
				right  *int
			)

			if i != 0 {
				top = &input[i-1][j]
				if *top <= num {
					res *= 0
				}
			}

			if i != numOfRows-1 {
				bottom = &input[i+1][j]
				if *bottom <= num {
					res *= 0
				}
			}

			if j != 0 {
				left = &input[i][j-1]
				if *left <= num {
					res *= 0
				}
			}

			if j != numOfCols-1 {
				right = &input[i][j+1]
				if *right <= num {
					res *= 0
				}
			}

			if res == 1 {
				lowPoints = append(lowPoints, num)
			}
		}
	}

	sum := 0

	for _, point := range lowPoints {
		sum += point + 1
	}

	fmt.Println(sum)

	return nil
}

func (p Problem_2021_09) Solve_02() error {
	return nil
}
