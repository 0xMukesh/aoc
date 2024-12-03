package problems_2024

import (
	"fmt"
	"os"
	"regexp"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2024_03 struct{}

func (p Problem_2024_03) Input() string {
	filename := "data/2024/03.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p *Problem_2024_03) Solve_01() error {
	input := p.Input()

	r, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	if err != nil {
		return err
	}

	sum := 0
	matches := r.FindAllString(input, -1)

	for _, text := range matches {
		var num1, num2 int
		fmt.Sscanf(text, "mul(%d,%d)", &num1, &num2)
		sum += num1 * num2
	}

	fmt.Println(sum)
	return nil
}

func (p Problem_2024_03) Solve_02() error {
	input := p.Input()

	r, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	if err != nil {
		return err
	}

	sum := 0
	isDisabled := false
	matches := r.FindAllString(input, -1)

	for _, text := range matches {
		if text == "don't()" {
			isDisabled = true
		}

		if text == "do()" {
			isDisabled = false
		}

		if !isDisabled {
			var num1, num2 int
			fmt.Sscanf(text, "mul(%d,%d)", &num1, &num2)
			sum += num1 * num2
		}
	}

	fmt.Println(sum)
	return nil
}
