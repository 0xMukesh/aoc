package problems_2025

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2025_06 struct{}

func (p Problem_2025_06) Input() string {
	filename := "data/2025/06.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2025_06) ParseInput(input string) ([][]int, []string, error) {
	lines := strings.Split(input, "\n")

	numericRows := [][]int{}
	symbols := strings.Fields(lines[len(lines)-1])

	for i := 0; i < len(lines)-1; i++ {
		row := []int{}
		fields := strings.Fields(lines[i])

		for _, v := range fields {
			num, err := strconv.Atoi(v)
			if err != nil {
				return numericRows, symbols, err
			}

			row = append(row, num)
		}

		numericRows = append(numericRows, row)
	}

	return numericRows, symbols, nil
}

func (p Problem_2025_06) Solve_01() error {
	numericRows, symbols, err := p.ParseInput(p.Input())
	if err != nil {
		return err
	}

	total := 0

	for i, v := range symbols {
		ans := 0
		if v == "*" {
			ans = 1
		}

		for j := range numericRows {
			switch v {
			case "+":
				ans += numericRows[j][i]
			case "*":
				ans *= numericRows[j][i]
			}
		}

		total += ans
	}

	fmt.Println(total)
	return nil
}

func (p Problem_2025_06) isSeparator(runeLines [][]rune, col int) bool {
	for r := range runeLines {
		if runeLines[r][col] != ' ' {
			return false
		}
	}
	return true
}

func (p Problem_2025_06) Solve_02() error {
	input := p.Input()
	lines := strings.Split(input, "\n")

	runeLines := make([][]rune, len(lines))
	for i, ln := range lines {
		r := []rune(ln)
		runeLines[i] = r
	}

	width := len(lines[0])
	opRow := len(runeLines) - 1
	total := 0

	for col := width - 1; col >= 0; {
		if p.isSeparator(runeLines, col) {
			col--
			continue
		}

		cols := []int{}
		for col >= 0 && !p.isSeparator(runeLines, col) {
			cols = append(cols, col)
			col--
		}

		nums := []int{}

		for _, c := range cols {
			digits := []rune{}
			for r := range opRow {
				ch := runeLines[r][c]
				if ch != ' ' {
					digits = append(digits, ch)
				}
			}

			if len(digits) == 0 {
				nums = append(nums, 0)
				continue
			}

			s := string(digits)
			n, err := strconv.Atoi(s)
			if err != nil {
				return err
			}

			nums = append(nums, n)
		}

		op := runeLines[opRow][cols[len(cols)-1]]

		switch op {
		case '+':
			sum := 0
			for _, n := range nums {
				sum += n
			}
			total += sum
		case '*':
			prod := 1
			for _, n := range nums {
				prod *= n
			}

			total += prod
		}
	}

	fmt.Println(total)
	return nil
}
