package problems

import (
	problems_2021 "github.com/0xmukesh/aoc/internal/problems/2021"
	problems_2024 "github.com/0xmukesh/aoc/internal/problems/2024"
)

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
		"05": problems_2021.Problem_2021_05{},
		"06": problems_2021.Problem_2021_06{},
		"07": problems_2021.Problem_2021_07{},
		"09": problems_2021.Problem_2021_09{},
	},
	"2024": {
		"01": problems_2024.Problem_2024_01{},
		"02": problems_2024.Problem_2024_02{},
		"03": &problems_2024.Problem_2024_03{},
	},
}
