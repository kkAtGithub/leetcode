package utils

func Min[T int | int64](a, b T) T {
	if a > b {
		return b
	}
	return a
}
