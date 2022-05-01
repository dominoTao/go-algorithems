package dynamic_programming

const (
	IntSize = 32 << (^uint(0) >> 63)
	IntMin = -1 << (IntSize - 1)
)


func CutRod(n int, l map[int]int) int {
	if n == 0 {
		return 0
	}
	max := IntMin
	for i := 1; i <= n; i++ {
		max = getMax(max, l[i] + CutRod(n - i, l))
	}
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

