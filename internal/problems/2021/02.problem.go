package problems_2021

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2021_02 struct{}

func (p Problem_2021_02) Input() string {
	filename := "data/2021/02.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2021_02) ParseInput(input string) []string {
	lines := strings.Split(input, "\n")
	return lines
}

func (p Problem_2021_02) Solve_01() error {
	input := p.ParseInput(p.Input())
	hor := 0
	depth := 0

	for _, line := range input {
		if strings.HasPrefix(line, "forward") {
			dist, err := strconv.Atoi(strings.Split(line, " ")[1])
			if err != nil {
				return err
			}

			hor += dist
		} else if strings.HasPrefix(line, "down") {
			dist, err := strconv.Atoi(strings.Split(line, " ")[1])
			if err != nil {
				return err
			}

			depth += dist
		} else if strings.HasPrefix(line, "up") {
			dist, err := strconv.Atoi(strings.Split(line, " ")[1])
			if err != nil {
				return err
			}

			depth -= dist
		} else {
			return errors.New("invalid direction")
		}
	}

	fmt.Println(hor * depth)

	return nil
}

func (p Problem_2021_02) Solve_02() error {
	input := p.ParseInput(p.Input())
	hor := 0
	depth := 0
	aim := 0

	for _, line := range input {
		if strings.HasPrefix(line, "forward") {
			dist, err := strconv.Atoi(strings.Split(line, " ")[1])
			if err != nil {
				return err
			}

			hor += dist
			depth += (aim * dist)
		} else if strings.HasPrefix(line, "down") {
			dist, err := strconv.Atoi(strings.Split(line, " ")[1])
			if err != nil {
				return err
			}

			aim += dist
		} else if strings.HasPrefix(line, "up") {
			dist, err := strconv.Atoi(strings.Split(line, " ")[1])
			if err != nil {
				return err
			}

			aim -= dist
		} else {
			return errors.New("invalid direction")
		}
	}

	fmt.Println(hor * depth)

	return nil
}
