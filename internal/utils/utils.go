package utils

import (
	"fmt"
	"os"
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
