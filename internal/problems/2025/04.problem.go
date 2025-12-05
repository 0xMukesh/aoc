package problems_2025

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2025_04 struct{}

func (p Problem_2025_04) Input() string {
	filename := "data/2025/04.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2025_04) ParseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))

	for i, line := range lines {
		row := make([]string, 0, len(line))

		for _, r := range line {
			row = append(row, string(r))
		}

		grid[i] = row
	}

	return grid
}

func (p Problem_2025_04) Solve_01() error {
	input := p.ParseInput(p.Input())
	counter := 0

	for i := range input {
		row := input[i]
		for j := range row {
			ch := row[j]

			if ch != "@" {
				continue
			}

			ajacents := utils.GetAdjacentPositionValuesInGrid(input, i, j)
			n := 0

			for _, v := range ajacents {
				if v == "@" {
					n = n + 1
				}
			}

			if n < 4 {
				counter = counter + 1
			}
		}
	}

	fmt.Println(counter)
	return nil
}

func (p Problem_2025_04) Solve_02() error {
	input := p.ParseInput(p.Input())
	k := math.MaxInt64
	counter := 0

	for k > 0 {
		k = 0

		for i := range input {
			row := input[i]

			for j := range row {
				ch := row[j]

				if ch != "@" {
					continue
				}

				ajacents := utils.GetAdjacentPositionValuesInGrid(input, i, j)
				n := 0

				for _, v := range ajacents {
					if v == "@" {
						n = n + 1
					}
				}

				if n < 4 {
					input[i][j] = "x"
					k = k + 1
					counter = counter + 1
				}
			}
		}
	}

	fmt.Println(counter)
	return nil
}
