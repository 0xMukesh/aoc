package problems_2024

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/0xmukesh/aoc/internal/utils"
)

type Problem_2024_03 struct {
	input string
	idx   int
}

func (p *Problem_2024_03) Input() string {
	filename := "data/2024/03.txt"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		utils.EPrint(err.Error())
	}

	p.input = string(bytes)
	return string(bytes)
}

func (p *Problem_2024_03) Peek() byte {
	return p.input[p.idx]
}

func (p *Problem_2024_03) Advance() {
	p.idx++
}

func (p *Problem_2024_03) Curr() byte {
	return p.input[p.idx-1]
}

func (p *Problem_2024_03) IsAtEnd() bool {
	return p.idx >= len(p.input)
}

func (p *Problem_2024_03) ConsumeStr(str string) bool {
	s := p.input[p.idx : p.idx+len(str)]
	if s == str {
		p.idx += len(str) + 1
		return true
	}

	return false
}

func (p *Problem_2024_03) ReadNum() string {
	s := string(p.Curr())

	for unicode.IsDigit(rune(p.Peek())) {
		p.Advance()
		s += string(p.Curr())
	}

	return s
}

func (p *Problem_2024_03) TrimFirstChar(s string) string {
	return s[1:]
}

func (p *Problem_2024_03) Solve_01() error {
	p.Input()

	sum := 0

	for !p.IsAtEnd() {
		p.Advance()

		ch := p.Curr()

		if ch == 'm' {
			found := p.ConsumeStr(p.TrimFirstChar("mul("))
			if !found {
				continue
			}

			firstNumStr := p.ReadNum()
			found = p.ConsumeStr(",")
			if !found {
				continue
			}

			secondNumStr := p.ReadNum()
			found = p.ConsumeStr(")")
			if !found {
				continue
			}

			firstNum, err := strconv.Atoi(firstNumStr)
			if err != nil {
				return err
			}
			secondNum, err := strconv.Atoi(secondNumStr)
			if err != nil {
				return err
			}

			sum += firstNum * secondNum
		}
	}

	fmt.Println(sum)
	return nil
}

func (p Problem_2024_03) Solve_02() error {
	p.Input()

	sum := 0
	isDisabled := false

	for !p.IsAtEnd() {
		p.Advance()

		ch := p.Curr()

		if ch == 'm' && !isDisabled {
			found := p.ConsumeStr(p.TrimFirstChar("mul("))
			if !found {
				continue
			}

			firstNumStr := p.ReadNum()
			found = p.ConsumeStr(",")
			if !found {
				continue
			}

			secondNumStr := p.ReadNum()
			found = p.ConsumeStr(")")
			if !found {
				continue
			}

			firstNum, err := strconv.Atoi(firstNumStr)
			if err != nil {
				return err
			}
			secondNum, err := strconv.Atoi(secondNumStr)
			if err != nil {
				return err
			}

			fmt.Printf("%d-%d\n", firstNum, secondNum)

			sum += firstNum * secondNum
		} else if ch == 'd' {
			found := p.ConsumeStr(p.TrimFirstChar("don't()"))
			if found {
				isDisabled = true
			}

			found = p.ConsumeStr(p.TrimFirstChar("do()"))
			if found {
				isDisabled = false
			}
		}
	}

	fmt.Println(sum)
	return nil
}
