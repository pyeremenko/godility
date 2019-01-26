package solution

func Solution(X int, Y int, D int) int {
	if X == Y {
		return 0
	}
	return (Y-X-1)/D + 1
}
