package main

import (
	"example/generic/utils"
	"fmt"
)

func main() {
	ints := []int{1, 2, 3}
	floats := []float64{1.0, 2.3, 3.5}
	uint64s := []uint64{1, 2, 3, 4}

	fmt.Println("=== not generic ===")
	fmt.Println("sum ints:", utils.SumInts(ints))
	fmt.Println("sum floats:", utils.SumFloats(floats))

	fmt.Println("")

	fmt.Println("=== generic ===")
	fmt.Println("sum ints:", utils.SumIntsOrFloats[int](ints))
	fmt.Println("sum floats:", utils.SumIntsOrFloats[float64](floats))

	fmt.Println("")

	fmt.Println("non-type-params sum ints:", utils.SumIntsOrFloats(ints))
	fmt.Println("non-type-params sum floats:", utils.SumIntsOrFloats(floats))

	fmt.Println("")

	fmt.Println("scalable sum ints:", utils.Sum(ints))
	fmt.Println("scalable sum floats:", utils.Sum(floats))
	fmt.Println("scalable sum uint64s:", utils.Sum(uint64s))
}
