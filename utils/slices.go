package utils

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Sum[T Number](arr []T) T {
	result := T(0)
	for _, v := range arr {
		result += v
	}
	return result
}

func Prod[T Number](arr []T) T {
	result := T(1)
	for _, v := range arr {
		result *= v
	}
	return result
}
