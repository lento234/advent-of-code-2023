package utils

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
