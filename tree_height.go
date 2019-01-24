package solution

import (
	"math"
)

type Tree struct {
	X int
	L *Tree
	R *Tree
}

func Solution(T *Tree) int {
	if T == nil {
		return -1
	}
	if T.L == nil && T.R == nil {
		return 0
	}
	return 1 + int(math.Max(float64(Solution(T.L)), float64(Solution(T.R))))
}
