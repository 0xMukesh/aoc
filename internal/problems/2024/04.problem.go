package problems_2024

import (
	"fmt"
	"os"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2024_04 struct{}

func (p Problem_2024_04) Input() string {
	filename := "data/2024/04.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2024_04) ParseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))

	for i := range grid {
		grid[i] = make([]string, len(lines[i]))
	}

	for i := range lines {
		for j := range lines[i] {
			grid[i][j] = string(lines[i][j])
		}
	}

	return grid
}

func (p Problem_2024_04) CheckDirection(grid [][]string, word string, startRow, startCol, rowStep, colStep int) bool {
	rows := len(grid)
	cols := len(grid[0])

	endRow := startRow + (len(word)-1)*rowStep
	endCol := startCol + (len(word)-1)*colStep

	if endRow < 0 || endRow >= rows || endCol < 0 || endCol >= cols {
		return false
	}

	for i := 0; i < len(word); i++ {
		r := startRow + i*rowStep
		c := startCol + i*colStep

		if grid[r][c] != string(word[i]) {
			return false
		}
	}

	return true
}

func (p Problem_2024_04) Solve_01() error {
	grid := p.ParseInput(p.Input())
	word := "XMAS"
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	rows := len(grid)
	cols := len(grid[0])

	sum := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if p.CheckDirection(grid, word, r, c, dir[0], dir[1]) {
					fmt.Println(r, c)
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
	return nil
}

func (p Problem_2024_04) Solve_02() error {
	return nil
}
