package main

import (
	"os"

	"github.com/0xmukesh/aoc/internal/handlers"
	"github.com/0xmukesh/aoc/internal/utils"
)

func main() {
	args := os.Args

	if len(args) != 4 {
		utils.EPrint("invalid usage\nusage: aoc <year> <problem>\n")
	}

	year := os.Args[1]
	problemNumber := os.Args[2]
	part := os.Args[3]

	handlers.ProblemRunner(year, problemNumber, part)
}
