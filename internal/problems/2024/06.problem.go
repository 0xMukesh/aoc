package problems_2024

import (
	"os"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2024_06 struct{}

func (p Problem_2024_06) Input() string {
	filename := "data/2024/06.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2024_06) Solve_01() error {
	// way too embarrassed to push the shitty, inefficient and crap code using which i have solve the question
	// TODO: improvize it and push it later
	return nil
}

func (p Problem_2024_06) Solve_02() error {
	// be me
	// > do some premature optimization with full confidence while solving part 2
	// > take a 30 mins break
	// > come back
	// > can't understand shit
	// > try to fix the buggy code
	// > debug for 2 hrs
	// > can't fix
	// > finally give up and chose not to push the code
	// TODO: improvize it and push it later
	return nil
}
