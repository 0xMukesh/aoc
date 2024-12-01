package problems_2024

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2024_01 struct{}

func (p Problem_2024_01) Input() string {
	filename := "data/2024/01.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2024_01) ParseInput(input string) ([]int, []int, error) {
	lines := strings.Split(input, "\n")
	var leftList []int
	var rightList []int

	for _, line := range lines {
		nums := strings.Split(line, "   ")
		if len(nums) > 2 {
			return nil, nil, errors.New("invalid format")
		}

		leftListNum, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, nil, err
		}

		rightListNum, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, nil, err
		}

		leftList = append(leftList, leftListNum)
		rightList = append(rightList, rightListNum)
	}

	return leftList, rightList, nil
}

func (p Problem_2024_01) Solve_01() error {
	input := p.Input()
	leftList, rightList, err := p.ParseInput(input)
	if err != nil {
		return err
	}

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})

	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	sum := 0.0

	for i := range leftList {
		sum += math.Sqrt(math.Pow(float64(leftList[i]-rightList[i]), 2))
	}

	fmt.Println(strconv.FormatFloat(sum, 'f', -1, 64))

	return nil
}

func (p Problem_2024_01) Solve_02() error {
	return nil
}
