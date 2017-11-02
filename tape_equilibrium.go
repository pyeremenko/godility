package solution

import (
	"math"
)

func Solution(A []int) int {
	n := len(A)
	total := sum(A)
	rightSum := 0
	var minDifference float64 = 2000 + 1
	for p := 1; p < n; p++ {
		rightSum += A[p-1]
		difference := math.Abs(float64(2*rightSum - total))
		if difference < minDifference {
			minDifference = difference
		}
	}
	return int(minDifference)
}

func sum(A []int) int {
	total := 0
	for _, val := range A {
		total += val
	}
	return total
}
