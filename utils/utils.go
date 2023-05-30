package utils

import (
	"sort"
)

func OrderPackSizes(order map[int]int) []int {
	// CREATE ORDERED SLICE OF PACKSIZES
	packSizes := make([]int, 0, len(order))
	for size := range order {
		packSizes = append(packSizes, size)
	}
	sort.Ints(packSizes)
	return packSizes
}

func Reverse(input []int) []int {
	var output []int

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}

func GetSum(arr []int) int {
	var sum int
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}
