package problems_2021

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2021_07 struct{}

func (p Problem_2021_07) Input() string {
	filename := "data/2021/07.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2021_07) ParseInput(input string) ([]int, error) {
	parts := strings.Split(input, ",")

	var nums []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}

		nums = append(nums, num)
	}

	return nums, nil
}

func (p Problem_2021_07) Solve_01() error {
	input, err := p.ParseInput(p.Input())
	if err != nil {
		return err
	}

	i := 0
	maxNum := 0
	minConsumption := math.MaxInt

	for _, num := range input {
		if num > maxNum {
			maxNum = num
		}
	}

	for {
		sum := 0

		if i > maxNum {
			break
		}

		for j := range input {
			// taking modulus
			sum += int(math.Sqrt(math.Pow((float64(input[j] - i)), 2)))
		}

		if sum < minConsumption {
			minConsumption = sum
		}

		i++
	}

	fmt.Println(minConsumption)

	return nil
}

func (p Problem_2021_07) Solve_02() error {
	input, err := p.ParseInput(p.Input())
	if err != nil {
		return err
	}

	i := 0
	maxNum := 0
	minConsumption := math.MaxInt

	for _, num := range input {
		if num > maxNum {
			maxNum = num
		}
	}

	for {
		sum := 0

		if i > maxNum {
			break
		}

		for j := range input {
			// taking modulus
			n := int(math.Sqrt(math.Pow(float64(input[j]-i), 2)))
			sum += n * (n + 1) / 2
		}

		if sum < minConsumption {
			minConsumption = sum
		}

		i++
	}

	fmt.Println(minConsumption)

	return nil
}
