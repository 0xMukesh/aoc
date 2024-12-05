package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func EPrint(m any) {
	fmt.Println(m)
	os.Exit(1)
}

func Map[S ~[]E, E any](s S, f func(E) E) S {
	result := make(S, len(s))
	for i := range s {
		result[i] = f(s[i])
	}
	return result
}

func Filter[S ~[]E, E any](s S, f func(E) bool) S {
	result := S{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func MostCommonElem[S ~[]E, E comparable](s S) (map[E]int, E) {
	hashmap := make(map[E]int)
	maxCount := 0
	var maxElem E

	for _, elem := range s {
		hashmap[elem]++
	}

	for k, v := range hashmap {
		if v > maxCount {
			maxCount = v
			maxElem = k
		}
	}

	return hashmap, maxElem
}

func LeastCommonElem[S ~[]E, E comparable](s S) (map[E]int, E) {
	hashmap := make(map[E]int)
	leastCount := math.MaxInt
	var leastElem E

	for _, elem := range s {
		hashmap[elem]++
	}

	for k, v := range hashmap {
		if v < leastCount {
			leastCount = v
			leastElem = k
		}
	}

	return hashmap, leastElem
}

func Transpose[S ~[][]E, E any](s S) [][]E {
	transpose := make([][]E, len(s[0]))

	for i := range transpose {
		transpose[i] = make([]E, len(s))
	}

	for i := range s {
		for j := range s[i] {
			transpose[j][i] = s[i][j]
		}
	}

	return transpose
}

func Mod(num float64) float64 {
	return math.Sqrt(math.Pow(num, 2))
}

func FormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func GetIdxByVal[T ~[]E, E comparable](s T, val E) *int {
	for i, v := range s {
		if v == val {
			return &i
		}
	}

	return nil
}
