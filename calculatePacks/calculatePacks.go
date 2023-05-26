package calculatepacks

import (
	"fmt"
	"math"
	"sort"
)

func recursive(allPossibleCalcs [][]int, count int, packSizes []int, accumulatedSizes []int) []int {
	var sum int
	for _, valueInt := range accumulatedSizes {
		sum += valueInt
	}
	// fmt.Println(accumulatedSizes)
	if sum >= count {
		// allPossibleCalcs = accumulatedSizes
		// fmt.Println(allPossibleCalcs)
		return accumulatedSizes
	} else {
		for _, size := range packSizes {
			accumulatedSizes = append(accumulatedSizes, size)
			return recursive(allPossibleCalcs, count, packSizes, accumulatedSizes)
		}
	}
	return accumulatedSizes
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

func getNumberRequiredForEachSize(packSizes []int, count int) map[int]float64 {
	var requiredNums = map[int]float64{}
	for _, size := range packSizes {

		requiredNums[size] = math.Ceil(float64(count) / float64(size))
		// fmt.Printf("%T", size)
	}
	return requiredNums
}

func allValuesEqualOne(numsMap map[int]float64) bool {
	returnVal := true
	for _, num := range numsMap {
		if num != 1 {
			returnVal = false
		}
		// fmt.Printf("%T", size)
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

	// GET HOW MANY OF EACH SIZE WOULD BE REQUIRED IF YOU COULD ONLY TAKE ONE SIZE
	var requiredNumsForEachSize = getNumberRequiredForEachSize(packSizes, count)

	// This variable will store arrays which contain all the number combinations that fulfill the order
	var allPossibleCalculations = [][]int{}
	fmt.Println(requiredNumsForEachSize)
	if allValuesEqualOne(requiredNumsForEachSize) {
		order[packSizes[0]]++
		return order
	} else {
		for _, size := range packSizes {
			acc := []int{size}
			var allCalcs []int = recursive(allPossibleCalculations, count, packSizes, acc)
			allPossibleCalculations = append(allPossibleCalculations, allCalcs)
			fmt.Println(allCalcs)

		}
		fmt.Println("allPossibleCalculations", allPossibleCalculations)
		return order
	}

}
