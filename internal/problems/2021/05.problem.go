package problems_2021

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2021_05 struct{}

func (p Problem_2021_05) Input() string {
	filename := "data/2021/05.txt"
	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2021_05) ParseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	var array_2d [][]int

	for i := range lines {
		parts := strings.Split(lines[i], " -> ")
		start_coords := strings.Split(parts[0], ",")
		end_coords := strings.Split(parts[1], ",")

		start_coord_x, _ := strconv.Atoi(start_coords[0])
		start_coord_y, _ := strconv.Atoi(start_coords[1])
		end_coord_x, _ := strconv.Atoi(end_coords[0])
		end_coord_y, _ := strconv.Atoi(end_coords[1])

		array_2d = append(array_2d, []int{start_coord_x, start_coord_y, end_coord_x, end_coord_y})
	}

	return array_2d
}

func (p Problem_2021_05) MaxNum(input [][]int) int {
	max := 0

	for i := range input {
		for j := range input[i] {
			if input[i][j] > max {
				max = input[i][j]
			}
		}
	}

	return max
}

func (p Problem_2021_05) AddPoint(diagram *[][]string, coordinates []int) error {
	x := coordinates[0]
	y := coordinates[1]

	if diagram == nil {
		return errors.New("can't add a point to nil diagram")
	}

	if (*diagram)[y][x] == "." {
		(*diagram)[y][x] = "1"
	} else {
		num, err := strconv.Atoi((*diagram)[y][x])
		if err != nil {
			return err
		}

		(*diagram)[y][x] = strconv.Itoa(num + 1)
	}

	return nil
}

func (p Problem_2021_05) Solve_01() error {
	input := p.ParseInput(p.Input())
	diagram := make([][]string, p.MaxNum(input)+1)

	for i := range diagram {
		diagram[i] = make([]string, p.MaxNum(input)+1)

		for j := range diagram[i] {
			diagram[i][j] = "."
		}
	}

	for i := range input {
		row := input[i]

		start_coord_x := row[0]
		start_coord_y := row[1]
		end_coord_x := row[2]
		end_coord_y := row[3]

		if !((start_coord_x == end_coord_x) || (start_coord_y == end_coord_y)) {
			continue
		}

		p.AddPoint(&diagram, []int{start_coord_x, start_coord_y})
		p.AddPoint(&diagram, []int{end_coord_x, end_coord_y})

		if start_coord_x == end_coord_x {
			var start int
			var end int

			if start_coord_y < end_coord_y {
				start = start_coord_y + 1
				end = end_coord_y
			} else {
				start = end_coord_y + 1
				end = start_coord_y
			}

			for j := start; j < end; j++ {
				p.AddPoint(&diagram, []int{start_coord_x, j})
			}
		}

		if start_coord_y == end_coord_y {
			var start int
			var end int

			if start_coord_x < end_coord_x {
				start = start_coord_x + 1
				end = end_coord_x
			} else {
				start = end_coord_x + 1
				end = start_coord_x
			}

			for j := start; j < end; j++ {
				p.AddPoint(&diagram, []int{j, start_coord_y})
			}
		}

	}

	count := 0

	for i := range diagram {
		for j := range diagram[i] {
			if diagram[i][j] == "." {
				continue
			}

			num, err := strconv.Atoi(diagram[i][j])
			if err != nil {
				return err
			}

			if num > 1 {
				count++
			}
		}
	}

	log.Print(count)

	return nil
}

func (p Problem_2021_05) Solve_02() error {
	return nil
}
