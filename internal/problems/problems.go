package problems

import problems_2021 "github.com/0xmukesh/aoc/internal/problems/2021"

type Problem interface {
	Input() string
	Solve_01() error
	Solve_02() error
}

type ProblemYearMap map[string]Problem

var ProblemsMap = map[string]ProblemYearMap{
	"2021": {
		"01": problems_2021.Problem_2021_01{},
		"02": problems_2021.Problem_2021_02{},
		"03": problems_2021.Problem_2021_03{},
		"04": problems_2021.Problem_2021_04{},
	},
}
