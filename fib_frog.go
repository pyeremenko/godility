package solution

const infinity = 200000

func Solution(A []int) int {
	n := len(A)
	possibleJumpLengths := getFibonacci(n)

	solutions := make([]int, n+2)
	solutions[0] = 0

	for i := 0; i < n; i++ {
		current := i + 1
		if A[i] == 0 {
			solutions[current] = infinity
		} else {
			solutions[current] = getCurrentSolution(current, possibleJumpLengths, solutions)
		}
	}

	solutions[n+1] = getCurrentSolution(n+1, possibleJumpLengths, solutions)

	result := solutions[n+1]
	if result == infinity {
		result = -1
	}
	return result
}

func getCurrentSolution(current int, possibleJumpLengths, solutions []int) int {
	minSolution := infinity
	for j := 1; possibleJumpLengths[j] <= current; j++ {
		if current-possibleJumpLengths[j] >= 0 {
			possibleSolution := solutions[current-possibleJumpLengths[j]]
			if possibleSolution < minSolution {
				minSolution = possibleSolution
			}
		}
	}
	result := minSolution + 1
	if result > infinity {
		result = infinity
	}
	return result
}

func getFibonacci(n int) []int {
	numbers := make([]int, 50)
	numbers[1] = 1
	current := 0
	for i := 2; current < (n+1)*2; i++ {
		current = numbers[i-1] + numbers[i-2]
		numbers[i] = current
	}
	return numbers
}
