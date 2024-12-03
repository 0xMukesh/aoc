package problems_2024

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

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

func (p Problem_2024_03) Solve_01() error {
	r, _ := regexp.Compile(`mul\(\s*(\d+)\s*,\s*(\d+)\s*\)`)
	matches := r.FindAllSubmatch([]byte(p.Input()), -1)

	sum := 0

	for _, match := range matches {
		firstNum, err := strconv.Atoi(string(match[1]))
		if err != nil {
			return err
		}
		secondNum, err := strconv.Atoi(string(match[2]))
		if err != nil {
			return err
		}

		sum += firstNum * secondNum
	}

	fmt.Println(sum)
	return nil
}

func (p Problem_2024_03) Solve_02() error {
	return nil
}
