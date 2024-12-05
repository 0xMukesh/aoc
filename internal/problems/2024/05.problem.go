package problems_2024

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2024_05 struct{}

func (p Problem_2024_05) Input() string {
	filename := "data/2024/05.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2024_05) ParseInput(input string) (map[int][]int, [][]int) {
	parts := strings.Split(input, "\n\n")
	rules := strings.Split(parts[0], "\n")
	updateLines := strings.Split(parts[1], "\n")

	rulesMapping := make(map[int][]int)

	for _, rule := range rules {
		var leftPageNum, rightPageNum int
		fmt.Sscanf(rule, "%d|%d", &leftPageNum, &rightPageNum)
		rulesMapping[leftPageNum] = append(rulesMapping[leftPageNum], rightPageNum)
	}

	var updates [][]int

	for _, update := range updateLines {
		strNums := strings.Split(update, ",")
		intNums := make([]int, len(strNums))

		for i, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			intNums[i] = num
		}

		updates = append(updates, intNums)
	}

	return rulesMapping, updates
}

func (p Problem_2024_05) Solve_01() error {
	rules, updates := p.ParseInput(p.Input())

	sum := 0

	for _, update := range updates {
		isValidUpdate := 1

		for i, pageNum := range update {
			rule, ok := rules[pageNum]

			if !ok {
				continue
			}

			for _, rulePageNum := range rule {
				idx := utils.GetIdxByVal(update, rulePageNum)

				if idx != nil {
					if *idx <= i {
						isValidUpdate *= 0
					}
				}
			}
		}

		if isValidUpdate != 0 {
			middleElemIdx := int(math.Floor(float64(len(update)) / 2))
			middleElem := update[middleElemIdx]
			sum += middleElem
		}
	}

	fmt.Println(sum)
	return nil
}

func (p Problem_2024_05) Solve_02() error {
	return nil
}
