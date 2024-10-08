package problems

type Problem interface {
	Input() string
	Solve_01() error
	Solve_02() error
}
