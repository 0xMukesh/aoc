package problems_2025

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2025_05 struct{}

func (p Problem_2025_05) Input() string {
	filename := "data/2025/05.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	return string(bytes)
}

type Interval struct {
	L int64
	R int64
}

func (p Problem_2025_05) ParseInput(input string) ([]Interval, []int64, error) {
	splits := strings.SplitN(input, "\n\n", 2)

	var ranges []Interval
	var ids []int64

	for _, v := range strings.Split(splits[0], "\n") {
		split := strings.Split(v, "-")
		lowerBound, err := strconv.Atoi(split[0])
		if err != nil {
			return ranges, ids, err
		}

		upperBound, err := strconv.Atoi(split[1])
		if err != nil {
			return ranges, ids, err
		}

		ranges = append(ranges, Interval{
			L: int64(lowerBound),
			R: int64(upperBound),
		})
	}

	for _, v := range strings.Split(splits[1], "\n") {
		id, err := strconv.Atoi(v)
		if err != nil {
			return ranges, ids, err
		}

		ids = append(ids, int64(id))
	}

	return ranges, ids, nil
}

func (p Problem_2025_05) Solve_01() error {
	ranges, ids, err := p.ParseInput(p.Input())
	if err != nil {
		return err
	}

	count := 0

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].L == ranges[i].R {
			return ranges[i].R < ranges[j].R
		}
		return ranges[i].L < ranges[j].L
	})
	slices.Sort(ids)

	for _, id := range ids {
		for len(ranges) > 0 && id > ranges[0].R {
			ranges = ranges[1:]
		}

		if len(ranges) == 0 {
			break
		}

		if id >= ranges[0].L && id <= ranges[0].R {
			count++
		}
	}

	fmt.Println(count)
	return nil
}

func (p Problem_2025_05) Solve_02() error {
	ranges, _, err := p.ParseInput(p.Input())
	if err != nil {
		return err
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].L == ranges[j].L {
			return ranges[i].R < ranges[j].R
		}
		return ranges[i].L < ranges[j].L
	})

	merged := []Interval{}
	curr := ranges[0]

	for i := 1; i < len(ranges); i++ {
		if ranges[i].L <= curr.R {
			if ranges[i].R > curr.R {
				curr.R = ranges[i].R
			}
		} else {
			merged = append(merged, curr)
			curr = ranges[i]
		}
	}

	merged = append(merged, curr)

	count := 0

	for _, v := range merged {
		count += int(v.R) - int(v.L) + 1
	}

	fmt.Println(count)
	return nil
}
