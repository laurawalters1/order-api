package calculatepacks

import (
	"sort"
)

func recursive(count int, packSizes []int, accumulatedSizes []int) []int {
	var sum int = getSum(accumulatedSizes)
	if sum >= count {
		return accumulatedSizes
	} else {

		for i, size := range packSizes {
			var prevNumsSum = getSum(packSizes[0 : i+1]) // returns sum of previous and current number in the packsizes array
			// BREAKING CONDITION
			if sum+size >= count {
				accumulatedSizes = append(accumulatedSizes, size)
				return accumulatedSizes
			} else if prevNumsSum+sum >= count { // if previous & curr number + sum >= count current number can be added as we know we won't need a larger number
				accumulatedSizes = append(accumulatedSizes, size)
				return recursive(count, packSizes, accumulatedSizes)
			} else if i == len(packSizes)-1 { // if current value is the largest packsize and count has not been met, we must add it to the array and loop back round
				accumulatedSizes = append(accumulatedSizes, size)
				return recursive(count, packSizes, accumulatedSizes)
			}
		}
		return accumulatedSizes
	}

}

func orderPackSizes(order map[int]int) []int {
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

func getSum(arr []int) int {
	var sum int
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}

func getOptimalSlice(arr [][]int) []int {
	var returnVal []int = arr[0]
	var optimalArrSum int

	for i, array := range arr {
		sum := getSum(array)
		if i == 0 {
			optimalArrSum = sum
			returnVal = array
		} else if sum < optimalArrSum {
			optimalArrSum = sum
			returnVal = array
		}
	}
	return returnVal
}

func CalculatePacks(count int) map[int]int {
	order := map[int]int{
		250:  0,
		500:  0,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	// PUT PACK SIZES IN ORDER
	packSizes := orderPackSizes(order)

	reversePackSizes := Reverse(packSizes)

	// This variable will store arrays which contain all the number combinations that fulfill the order
	var allPossibleCalculations = [][]int{}

	for _, size := range reversePackSizes {
		acc := []int{size}
		var allCalcs []int = recursive( count, packSizes, acc)
		allPossibleCalculations = append(allPossibleCalculations, allCalcs)

	}
	optimalSlice := getOptimalSlice(allPossibleCalculations)

	for _, size := range optimalSlice {
		order[size]++
	}
	return order

}
