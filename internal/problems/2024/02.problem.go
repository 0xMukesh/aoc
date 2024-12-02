package problems_2024

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2024_02 struct{}

func (p Problem_2024_02) Input() string {
	filename := "data/2024/02.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}
func (p Problem_2024_02) ParseInput(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")
	parsed := make([][]int, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		row := make([]int, len(parts))

		for j, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			row[j] = num
		}
		parsed[i] = row
	}

	return parsed, nil
}

func (p Problem_2024_02) Solve_01() error {
	input := p.Input()
	levels, err := p.ParseInput(input)
	if err != nil {
		return err
	}

	sum := 0

	for _, row := range levels {
		isSafe := true

		for j := range row {
			if j == len(row)-1 {
				continue
			}

			diff := row[j+1] - row[j]
			mod := math.Sqrt(math.Pow(float64(diff), 2))

			if mod != 1 && mod != 2 && mod != 3 {
				isSafe = false
				break
			}

			if row[1]-row[0] > 0 && diff < 0 {
				isSafe = false
				break
			} else if row[1]-row[0] < 0 && diff > 0 {
				isSafe = false
				break
			} else if diff == 0 {
				isSafe = false
				break
			}
		}

		if isSafe {
			sum++
		}
	}

	fmt.Println(sum)

	return nil
}

func (p Problem_2024_02) Solve_02() error {
	return nil
}
