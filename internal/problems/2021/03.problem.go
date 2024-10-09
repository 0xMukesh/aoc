package problems_2021

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2021_03 struct{}

func (p Problem_2021_03) Input() string {
	filename := "data/2021/03.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

func (p Problem_2021_03) ParseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	array2d := make([][]string, len(lines))
	parsedArray2d := make([][]int, len(lines))

	for i := range lines {
		array2d[i] = strings.Split(lines[i], "")
	}

	for i := range parsedArray2d {
		parsedArray2d[i] = make([]int, len(lines[0]))
	}

	for i, row := range array2d {
		for j := range row {
			num, _ := strconv.Atoi(array2d[i][j])
			parsedArray2d[i][j] = num
		}
	}

	return parsedArray2d
}

// func (p Problem_2021_03) GenerateTranspose(input [][]int) [][]int {
// 	transpose := make([][]int, len(input[0]))

// 	for i := range transpose {
// 		transpose[i] = make([]int, len(input))
// 	}

// 	for i, row := range input {
// 		for j := range row {
// 			transpose[j][i] = input[i][j]
// 		}
// 	}

// 	return transpose
// }

func (p Problem_2021_03) Solve_01() error {
	input := utils.Transpose(p.ParseInput(p.Input()))

	gammaRateStr := ""
	epsilonRateStr := ""

	for _, bits := range input {
		_, mostCommonElem := utils.MostCommonElem(bits)
		_, leastCommonElem := utils.LeastCommonElem(bits)
		gammaRateStr += strconv.Itoa(mostCommonElem)
		epsilonRateStr += strconv.Itoa(leastCommonElem)
	}

	gammaRate, _ := strconv.ParseInt(gammaRateStr, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRateStr, 2, 64)

	fmt.Println(gammaRate * epsilonRate)

	return nil
}

func (p Problem_2021_03) Solve_02() error {
	array2d := p.ParseInput(p.Input())

	for i := 0; len(array2d) > 1; i++ {
		transpose := utils.Transpose(array2d)

		if len(array2d) == 1 {
			break
		}

		bits := transpose[i]
		hashmap, mostCommonElem := utils.MostCommonElem(bits)

		if hashmap[0] == hashmap[1] {
			mostCommonElem = 1
		}

		array2d = utils.Filter(array2d, func(v []int) bool {
			return v[i] == mostCommonElem
		})
	}

	oxygenGeneratorStr := ""
	for i := range array2d[0] {
		oxygenGeneratorStr += strconv.Itoa(array2d[0][i])
	}
	oxygenGeneratorRating, _ := strconv.ParseInt(oxygenGeneratorStr, 2, 64)

	array2d = p.ParseInput(p.Input())

	for i := 0; len(array2d) > 1; i++ {
		transpose := utils.Transpose(array2d)

		if len(array2d) == 1 {
			break
		}

		bits := transpose[i]
		hashmap, leastCommonElem := utils.LeastCommonElem(bits)

		if hashmap[0] == hashmap[1] {
			leastCommonElem = 0
		}

		array2d = utils.Filter(array2d, func(v []int) bool {
			return v[i] == leastCommonElem
		})
	}

	co2ScrubberRatingStr := ""
	for i := range array2d[0] {
		co2ScrubberRatingStr += strconv.Itoa(array2d[0][i])
	}
	co2ScrubberRating, _ := strconv.ParseInt(co2ScrubberRatingStr, 2, 64)

	log.Print(oxygenGeneratorRating * co2ScrubberRating)

	return nil
}
