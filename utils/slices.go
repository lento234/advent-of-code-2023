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

func LCMSlice(arr []int) int {
	result := LCM(arr[0], arr[1])
	for i := 2; i < len(arr); i++ {
		result = LCM(result, arr[i])
	}
	return result
}
