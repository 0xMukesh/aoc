package problems_2025

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2025_02 struct{}

func (p Problem_2025_02) Input() string {
	filename := "data/2025/02.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2025_02) ParseInput(input string) [][]int {
	splitByComma := strings.Split(input, ",")
	parsed := make([][]int, len(splitByComma))

	for i, v := range splitByComma {
		splitByHyphen := strings.Split(v, "-")
		lowerBound, _ := strconv.Atoi(splitByHyphen[0])
		upperBound, _ := strconv.Atoi(splitByHyphen[1])

		parsed[i] = []int{lowerBound, upperBound}
	}

	return parsed
}

func (p Problem_2025_02) Solve_01() error {
	input := p.ParseInput(p.Input())
	var counter int64 = 0

	for _, v := range input {
		lowerBound := v[0]
		upperBound := v[1]

		var repeatedIds []int64
		maxDigits := len(strconv.FormatInt(int64(upperBound), 10))

		// an invalid id consists of repeating unit S
		// if length of that repeating unit S is k
		// then all the ids containing S with every possible k-digit integer is invalid
		// the range of k-digit integer is 10^(k - 1) to (10^k) - 1
	OuterLoop:
		for k := 1; 2*k <= maxDigits; k++ {
			lowS := math.Pow10(k - 1)
			highS := math.Pow10(k) - 1

			for s := lowS; s <= highS; s++ {
				sStr := strconv.FormatInt(int64(s), 10)
				nStr := sStr + sStr // seq of digits repeated twice
				n, err := strconv.Atoi(nStr)
				if err != nil {
					break OuterLoop
				}

				if n < lowerBound {
					continue
				}

				if n > upperBound {
					continue
				}

				repeatedIds = append(repeatedIds, int64(n))
			}
		}

		for _, id := range repeatedIds {
			counter = counter + id
		}
	}

	fmt.Println(counter)
	return nil
}

func (p Problem_2025_02) Solve_02() error {
	input := p.ParseInput(p.Input())
	var counter int64 = 0

	for _, v := range input {
		lowerBound := v[0]
		upperBound := v[1]

		// var repeatedIds []int64
		repeatedIds := make(map[int64]bool)
		maxDigits := len(strconv.FormatInt(int64(upperBound), 10))

		// an invalid id consists of repeating unit S
		// if length of that repeating unit S is k
		// then all the ids containing S with every possible k-digit integer is invalid
		// the range of k-digit integer is 10^(k - 1) to (10^k) - 1
	OuterLoop:
		for k := 1; 2*k <= maxDigits; k++ {
			lowS := int64(math.Pow10(k - 1))
			highS := int64(math.Pow10(k) - 1)

			maxR := maxDigits / k
			for r := 2; r <= maxR; r++ {
				for s := lowS; s <= highS; s++ {
					sStr := strconv.FormatInt(s, 10)
					nStr := strings.Repeat(sStr, r)
					n, err := strconv.ParseInt(nStr, 10, 64)
					if err != nil {
						break OuterLoop
					}

					if n < int64(lowerBound) {
						continue
					}

					if n > int64(upperBound) {
						break
					}

					repeatedIds[n] = true
				}
			}
		}

		for id := range repeatedIds {
			counter = counter + id
		}
	}

	fmt.Println(counter)

	return nil
}
