package problems_2025

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2025_03 struct{}

func (p Problem_2025_03) Input() string {
	filename := "data/2025/03.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2025_03) ParseInput(input string) []string {
	splits := strings.Split(p.Input(), "\n")
	return splits
}

func (p Problem_2025_03) Solve_01() error {
	input := p.ParseInput(p.Input())
	counter := 0

	for _, v := range input {
		highestJoltage := 0

		for i, ch := range v {
			for j := i + 1; j <= len(v)-1; j++ {
				joltage, err := strconv.Atoi(string(ch) + string(v[j]))
				if err != nil {
					return err
				}

				if joltage > highestJoltage {
					highestJoltage = joltage
				}
			}
		}

		counter = counter + highestJoltage
	}

	fmt.Println(counter)
	return nil
}

func (p Problem_2025_03) Solve_02() error {
	input := p.ParseInput(p.Input())
	n := 12
	counter := big.NewInt(0)

	for _, v := range input {
		stack := make([]byte, 0, n)
		m := len(v)

		for i := range m {
			for len(stack) > 0 && v[i] > stack[len(stack)-1] && len(stack)+m-i-1 >= n {
				stack = stack[:len(stack)-1]
			}

			if len(stack) < n {
				stack = append(stack, v[i])
			}
		}

		num, err := strconv.Atoi(string(stack))
		if err != nil {
			return err
		}

		counter = counter.Add(counter, big.NewInt(int64(num)))
	}

	fmt.Println(counter.String())
	return nil
}
