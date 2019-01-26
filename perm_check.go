package solution

func Solution(A []int) int {
	size := len(A)
	flags := make([]bool, size)

	pointer := 0
	for {
		if pointer >= size {
			break
		}
		if flags[pointer] {
			pointer += 1
			continue
		}
		if A[pointer] == pointer+1 {
			flags[pointer] = true
			pointer += 1
			continue
		}
		if A[pointer] > size || A[pointer] < 1 {
			return 0
		}

		current := A[pointer]

		if current == A[current-1] {
			return 0
		} else {
			A[pointer] = A[current-1]
			A[current-1] = current
			flags[current-1] = true
			if A[pointer] == pointer+1 {
				flags[pointer] = true
			}
		}
	}
	return 1
}
