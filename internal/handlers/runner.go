package handlers

import (
	"fmt"

	"github.com/0xmukesh/aoc/internal/problems"
	problems_2021 "github.com/0xmukesh/aoc/internal/problems/2021"
	"github.com/0xmukesh/aoc/internal/utils"
)

type ProblemYearMap map[string]problems.Problem

var problemsMap = map[string]ProblemYearMap{
	"2021": {
		"01": problems_2021.Problem_2021_01{},
	},
}

func ProblemRunner(year string, problemNumber string, part string) {
	yearMap, ok := problemsMap[year]
	if !ok {
		utils.EPrint(fmt.Sprintf("%s year doesn't exist", year))
	}

	problem, ok := yearMap[problemNumber]
	if !ok {
		utils.EPrint(fmt.Sprintf("%s problem doesn't exist in %s year", problemNumber, year))
	}

	if part == "1" {
		if err := problem.Solve_01(); err != nil {
			utils.EPrint(fmt.Sprintf("an error occured while running part 1 of %s problem of %s year\nerror: %s", problemNumber, year, err.Error()))
		}
	} else if part == "2" {
		if err := problem.Solve_02(); err != nil {
			utils.EPrint(fmt.Sprintf("an error occured while running part 2 of %s problem of %s year\nerror: %s", problemNumber, year, err.Error()))
		}
	} else {
		utils.EPrint("invalid part number")
	}
}
