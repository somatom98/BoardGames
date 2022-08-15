package mmath

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Diff(x int, y int) int {
	return Abs(x - y)
}
