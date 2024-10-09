package problems_2021

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2021_04 struct{}

func (p Problem_2021_04) Input() string {
	filename := "data/2021/04.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2021_04) ParseInput(input string) ([]string, [][][]string) {
	lines := strings.Split(input, "\n\n")
	randomNums := strings.Split(lines[0], ",")

	var bingos [][][]string

	for i := range lines {
		if i == 0 {
			continue
		}

		bingoRows := strings.Split(lines[i], "\n")
		var bingo [][]string

		for _, row := range bingoRows {
			bingo = append(bingo, strings.Fields(row))
		}

		bingos = append(bingos, bingo)
	}

	return randomNums, bingos
}

func (p Problem_2021_04) CalculateScore(bingo [][]string, lastNumber string) int {
	result := 0

	for i := range bingo {
		for j := range bingo[i] {
			if bingo[i][j] != "x" {
				num, _ := strconv.Atoi(bingo[i][j])
				result += num
			}
		}
	}

	lastNum, _ := strconv.Atoi(lastNumber)

	return result * lastNum
}

func (p Problem_2021_04) Solve_01() error {
	randomNums, bingos := p.ParseInput(p.Input())

outerLoop:
	for _, randomNum := range randomNums {
		for _, bingo := range bingos {
			for i := range bingo {
				for j := range bingo[i] {
					if bingo[i][j] == randomNum {
						bingo[i][j] = "x"
					}
				}

				if strings.Join(bingo[i], "") == "xxxxx" {
					log.Print(p.CalculateScore(bingo, randomNum))
					break outerLoop
				}
			}

			for _, row := range utils.Transpose(bingo) {
				if strings.Join(row, "") == "xxxxx" {
					log.Print(p.CalculateScore(bingo, randomNum))
					break outerLoop
				}
			}
		}
	}

	return nil
}

func (p Problem_2021_04) Solve_02() error {
	return nil
}
