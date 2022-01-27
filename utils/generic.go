package utils

func SumIntsOrFloats[T int | float64](data []T) T {
	var total T
	for _, value := range data {
		total += value
	}

	return total
}

type Number interface {
	int | float64 | uint64
}

func Sum[T Number](data []T) T {
	var total T
	for _, value := range data {
		total += value
	}

	return total
}
