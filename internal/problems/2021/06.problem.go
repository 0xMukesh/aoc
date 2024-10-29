package problems_2021

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2021_06 struct{}

func (p Problem_2021_06) Input() string {
	filename := "data/2021/06.mock.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2021_06) ParseInput(input string) []int {
	nums := strings.Split(input, ",")
	var parsed []int

	for i := range nums {
		num, _ := strconv.Atoi(nums[i])
		parsed = append(parsed, num)
	}

	return parsed
}

func (p Problem_2021_06) Solve_01() error {
	input := p.ParseInput(p.Input())
	i := 0

	for {
		if i >= 18 {
			break
		}

		for j := range input {
			if input[j] == 0 {
				input[j] = 6
				input = append(input, 8)
				continue
			}

			input[j]--
		}

		i++
	}

	fmt.Println(len(input))

	return nil
}

func (p Problem_2021_06) Solve_02() error {
	input := p.ParseInput(p.Input())
	i := 0

	for {
		if i >= 256 {
			break
		}

		for j := range input {
			if input[j] == 0 {
				input[j] = 6
				input = append(input, 8)
				continue
			}

			input[j]--
		}

		i++
	}

	fmt.Println(len(input))

	return nil
}
