package handlers

import (
	"fmt"

	"github.com/0xmukesh/aoc/internal/problems"
	"github.com/0xmukesh/aoc/internal/utils"
)

func ProblemRunner(year string, problemNumber string, part string) {
	yearMap, ok := problems.ProblemsMap[year]
	if !ok {
		utils.EPrint(fmt.Sprintf("%s year doesn't exist", year))
	}

	problem, ok := yearMap[problemNumber]
	if !ok {
		utils.EPrint(fmt.Sprintf("%s problem doesn't exist in %s year", problemNumber, year))
	}

	switch part {
	case "1":
		if err := problem.Solve_01(); err != nil {
			utils.EPrint(fmt.Sprintf("an error occured while running part 1 of %s problem of %s year\nerror: %s", problemNumber, year, err.Error()))
		}
	case "2":
		if err := problem.Solve_02(); err != nil {
			utils.EPrint(fmt.Sprintf("an error occured while running part 2 of %s problem of %s year\nerror: %s", problemNumber, year, err.Error()))
		}
	default:
		utils.EPrint("invalid part number")
	}
}
