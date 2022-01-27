package utils

func SumInts(data []int) int {
	var total int
	for _, value := range data {
		total += value
	}

	return total
}

func SumFloats(data []float64) float64 {
	var total float64
	for _, value := range data {
		total += value
	}

	return total
}
